package routing

import (
	. "github.com/alameddinc/terminal-wars/controllers"
	"github.com/gorilla/mux"
)

var r = mux.NewRouter()

func routes() {
	r.HandleFunc("/", GetHome).Methods("GET")
	// user/change_hash/...
	userRoutes()
	// api/farming/...
}

func userRoutes() {
	ur := createSubRouter("user")
	ur.HandleFunc("/register", PostRegister).Methods("POST")
}
