package routing

import (
	"log"
	"net/http"
)

func Handle() {
	routes()
	log.Println("127.0.0.1:8008 Started")
	http.ListenAndServe(":8008", r)
}
