package main

import (
	"LGF/internal/data"
	"LGF/internal/validator"
	"errors"
	"fmt"
	"net/http"
)

func (app *application) createArtifactHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name   string    `json:"name"`
		Origin string    `json:"origin"`
		Year   data.Year `json:"year"`
		Type   string    `json:"type"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	artifact := &data.Artifact{
		Name:   input.Name,
		Origin: input.Origin,
		Year:   input.Year,
		Type:   input.Type,
	}

	v := validator.New()
	if data.ValidateArtifact(v, artifact); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Artifacts.Insert(artifact)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/artifacts/%d", artifact.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"artifact": artifact}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
func (app *application) showArtifactHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	artifact, err := app.models.Artifacts.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"artifact": artifact}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateArtifactHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	artifact, err := app.models.Artifacts.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Name   *string    `json:"name"`
		Origin *string    `json:"origin"`
		Year   *data.Year `json:"year"`
		Type   *string    `json:"type"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Name != nil {
		artifact.Name = *input.Name
	}
	if input.Origin != nil {
		artifact.Origin = *input.Origin
	}
	if input.Year != nil {
		artifact.Year = *input.Year
	}
	if input.Type != nil {
		artifact.Type = *input.Type
	}

	v := validator.New()
	if data.ValidateArtifact(v, artifact); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Artifacts.Update(artifact)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"artifact": artifact}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) deleteArtifactHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Artifacts.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "artifact successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listArtifactsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name   string
		Origin string
		Type   string
		data.Filters
	}
	v := validator.New()
	qs := r.URL.Query()

	input.Name = app.readString(qs, "name", "")
	input.Origin = app.readString(qs, "origin", "")
	input.Type = app.readString(qs, "type", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "name", "origin", "year", "type", "-id", "-name", "-origin", "-year", "-type"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	artifacts, metadata, err := app.models.Artifacts.GetAll(input.Name, input.Origin, input.Type, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"artifacts": artifacts, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
