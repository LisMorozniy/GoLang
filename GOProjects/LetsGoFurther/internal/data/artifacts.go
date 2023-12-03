package data

import (
	"LGF/internal/validator"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Artifact struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Origin    string    `json:"origin"`
	Year      Year      `json:"year,omitempty"`
	Type      string    `json:"type"`
	Version   int32     `json:"version"`
}

func ValidateArtifact(v *validator.Validator, artifact *Artifact) {
	v.Check(artifact.Name != "", "name", "must be provided")
	v.Check(len(artifact.Name) <= 500, "name", "must not be more than 500 bytes long")
	v.Check(artifact.Year != 0, "year", "must be provided")
	v.Check(artifact.Type != "", "type", "must be provided")
	v.Check(len(artifact.Type) <= 500, "type", "must not be more than 500 bytes long")
	v.Check(artifact.Type != "", "type", "must be provided")
	v.Check(len(artifact.Type) <= 500, "type", "must not be more than 500 bytes long")
}

type ArtifactModel struct {
	DB *sql.DB
}

func (m ArtifactModel) Insert(artifact *Artifact) error {

	query := `
	INSERT INTO artifacts (name, origin, year, type)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, version`

	args := []interface{}{artifact.Name, artifact.Origin, artifact.Year, artifact.Type}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return m.DB.QueryRowContext(ctx, query, args...).Scan(&artifact.ID, &artifact.CreatedAt, &artifact.Version)
}

func (m ArtifactModel) Get(id int64) (*Artifact, error) {

	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
    SELECT id, created_at, name, origin, year, type, version
    FROM artifacts
    WHERE id = $1`

	var artifact Artifact

	err := m.DB.QueryRow(query, id).Scan(
		&artifact.ID,
		&artifact.CreatedAt,
		&artifact.Name,
		&artifact.Origin,
		&artifact.Year,
		&artifact.Type,
		&artifact.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &artifact, nil
}

func (m ArtifactModel) GetAll(name string, origin string, artifactType string, filters Filters) ([]*Artifact, Metadata, error) {
	// Update the SQL query to include the filter conditions.
	query := fmt.Sprintf(`
    SELECT count(*) OVER(), id, created_at, name, origin, year, type, version
    FROM artifacts
    WHERE (to_tsvector('simple', name) @@ plainto_tsquery('simple', $1) OR $1 = '')
    AND (LOWER(origin) = LOWER($2) OR $2 = '')
    AND (LOWER(type) = LOWER($3) OR $3 = '')
    ORDER BY %s %s, id ASC
	LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection()) // Assuming SortColumn and SortDirection are methods of Filters

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	args := []interface{}{name, origin, artifactType, filters.limit(), filters.offset()}
	// Pass the name, origin, and type as the placeholder parameter values.
	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()

	totalRecords := 0
	artifacts := []*Artifact{}
	for rows.Next() {
		var artifact Artifact
		err := rows.Scan(
			&totalRecords,
			&artifact.ID,
			&artifact.CreatedAt,
			&artifact.Name,
			&artifact.Origin,
			&artifact.Year,
			&artifact.Type,
			&artifact.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		artifacts = append(artifacts, &artifact)
	}
	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)
	return artifacts, metadata, nil
}

func (m ArtifactModel) Update(artifact *Artifact) error {

	query := `
    UPDATE artifacts
    SET name = $1, origin = $2, year = $3, type = $4, version = version + 1
    WHERE id = $5 AND version = $6
    RETURNING version`

	args := []interface{}{
		artifact.Name,
		artifact.Origin,
		artifact.Year,
		artifact.Type,
		artifact.ID,
		artifact.Version,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&artifact.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (m ArtifactModel) Delete(id int64) error {

	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
    DELETE FROM artifacts
    WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
