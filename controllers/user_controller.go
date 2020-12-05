package controllers

import (
	"encoding/json"
	"github.com/alameddinc/terminal-wars/schemas"
	"net/http"
)

func PostRegister(w http.ResponseWriter, r *http.Request){
	json.NewDecoder(r.Body).Decode(&requestSchema)
	json.NewEncoder(w).Encode("Register Sayfası")
}

