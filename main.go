package main

import (
	"log"
	"net/http"

	"github.com/CezarGarrido/nasa-sonda/handler"
	"github.com/CezarGarrido/nasa-sonda/sonda"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	headersOk := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Cache-Control", "X-File-Name", "Origin", "X-Session-ID"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE"})

	log.Println("Run on port:8089")

	router := mux.NewRouter()
	sonda := sonda.NewProbe()

	sondaHandler := handler.NewHandler(sonda)
	sondaHandler.MakeRoutes(router)

	err := http.ListenAndServe(":8089", handlers.CORS(headersOk, methodsOk, originsOk)(router))
	if err != nil {
		log.Panic(err)
	}
}
