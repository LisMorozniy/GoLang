package main
import (
"fmt"
"net/http"
"strconv"
"github.com/julienschmidt/httprouter"
)
func (app *application) createArtifactHandler(w http.ResponseWriter, r *http.Request) {
fmt.Fprintln(w, "create a new artifact")
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
    
    