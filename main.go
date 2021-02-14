package main

import (
	"log"
	"net/http"
	"os"

	"github.com/CezarGarrido/nasa-sonda/handler"
	"github.com/CezarGarrido/nasa-sonda/sonda"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := map[bool]string{true: os.Getenv("PORT"), false: "8089"}[os.Getenv("PORT") != ""]

	headersOk := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Cache-Control", "X-File-Name", "Origin", "X-Session-ID"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE"})

	log.Println("Run on port:"+ port)

	router := mux.NewRouter()
	sonda := sonda.NewProbe()

	sondaHandler := handler.NewHandler(sonda)
	sondaHandler.MakeRoutes(router)

	err := http.ListenAndServe(":"+port, handlers.CORS(headersOk, methodsOk, originsOk)(router))
	if err != nil {
		log.Panic(err)
	}
}
