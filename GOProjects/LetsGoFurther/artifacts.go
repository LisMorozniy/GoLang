package main
import (
"encoding/json"
"fmt"
"net/http"
"time"
"LetsGoFurther/internal/data"
"LetsGoFurther/internal/validator"
)
func (app *application) createArtifactHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Name string `json:"name"`
        Origin string `json:"origin"`
        Year data.Year `json:"year"`
        Type string `json:"type"`
        }
err := app.readJSON(w, r, &input)
if err != nil {
app.badRequestResponse(w, r, err)
return
}
artifact := &data.Artifact{
    Name: input.Name,
    Origin: input.Origin,
    Year: input.Year,
    Type: input.Type,
    }
    // Initialize a new Validator.
    v := validator.New()
    if data.ValidateArtifact(v, artifact); !v.Valid() {
    app.failedValidationResponse(w, r, v.Errors)
    return
    }
fmt.Fprintf(w, "%+v\n", input)    
}
func (app *application) showArtifactHandler(w http.ResponseWriter, r *http.Request) {
    id, err := app.readIDParam(r)
    if err != nil {
    app.notFoundResponse(w, r)
    return
    }
    artifact := data.Artifact{
    ID: id,
    CreatedAt: time.Now(),
    Name: "Crown",
    Origin: "Kazakh",
    Year: 102,
    Type: "Jewlry",
    Version: 1,
    }
    err = app.writeJSON(w, http.StatusOK, envelope{"artifact": artifact}, nil)
    if err != nil {
        app.serverErrorResponse(w, r, err)
        }
    }
    
    