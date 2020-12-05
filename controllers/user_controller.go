package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/alameddinc/terminal-wars/schemas"
	"net/http"
)

func PostRegister(w http.ResponseWriter, r *http.Request) {
	var reqSchema schemas.UserRegisterPostSchema
	json.NewDecoder(r.Body).Decode(&reqSchema)
	fmt.Println(reqSchema)
	// TODO Save Method
	json.NewEncoder(w).Encode("Selam Sahip " + reqSchema.Username)
}

func GetLogs(w http.ResponseWriter, r *http.Request) {

}

func PostLogin(w http.ResponseWriter, r *http.Request) {

}

func PostChangePassword(w http.ResponseWriter, r *http.Request) {

}

func PostChangeHash(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}
