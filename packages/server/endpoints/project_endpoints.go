package endpoints

import (
	"github.com/datatug/datatug/packages/api"
	"net/http"
)

var _ ProjectEndpoints = (*ProjectAgentEndpoints)(nil)

type ProjectAgentEndpoints struct {
}

func (_ ProjectAgentEndpoints) CreateProject(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Creation of a new DataTug project is not implemented at agent yet. For now use DataTug CLI to create a new project."))
}

func (_ ProjectAgentEndpoints) DeleteProject(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Deletion of a DataTug project is not implemented at agent yet."))
}

// GetProjectSummary a handler to return project summary
func GetProjectSummary(w http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get(urlQueryParamID)
	projectSummary, err := api.GetProjectSummary(id)
	returnJSON(w, request, http.StatusOK, err, projectSummary)
}