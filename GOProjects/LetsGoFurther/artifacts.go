package main
import (
"encoding/json"
"fmt"
"net/http"
"strconv"
"github.com/julienschmidt/httprouter"
)
func (app *application) createArtifactHandler(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Name string `json:"name"`
        Year data.Year `json:"year"`
        Type string `json:"type"`
        }
        err := app.readJSON(w, r, &input)
if err != nil {
// Use the new badRequestResponse() helper.
app.badRequestResponse(w, r, err)
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
    
    