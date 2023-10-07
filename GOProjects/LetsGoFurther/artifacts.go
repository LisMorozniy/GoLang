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
        Year int32 `json:"year"`
        Type string `json:"type"`
        }
        err := json.NewDecoder(r.Body).Decode(&input)
        if err != nil {
        app.errorResponse(w, r, http.StatusBadRequest, err.Error())
        return
        }
        // Dump the contents of the input struct in a HTTP response.
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
    
    