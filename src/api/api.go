package api

import (
	"net/http"
	"log"
)

// path: /api
func ApiServer(w http.ResponseWriter, req *http.Request)  {
	log.Println("request to api")
}