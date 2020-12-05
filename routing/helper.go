package routing

import "github.com/gorilla/mux"

func createSubRouter(urlText string) *mux.Router {
	return r.PathPrefix("/" + urlText).Subrouter()
}
