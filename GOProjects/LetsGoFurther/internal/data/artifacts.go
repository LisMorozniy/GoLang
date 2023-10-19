package data

import (
	"LGF/internal/validator"
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
