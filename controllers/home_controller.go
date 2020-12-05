package controllers

import (
	"encoding/json"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode("Home Page")
}
