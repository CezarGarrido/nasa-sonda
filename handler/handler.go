package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	sonda "github.com/CezarGarrido/nasa-sonda/sonda"
	"github.com/gorilla/mux"
)

type SondaHandler struct {
	sonda *sonda.Probe
}

func NewHandler(sonda *sonda.Probe) *SondaHandler {
	return &SondaHandler{
		sonda: sonda,
	}
}

type Command struct {
	Movimentos []string
}

func (handler *SondaHandler) MakeRoutes(router *mux.Router) {
	router.HandleFunc("/api/sonda", handler.FindSonda).Methods("GET")
	router.HandleFunc("/api/sonda/commands", handler.Commands).Methods("POST")
	router.HandleFunc("/api/sonda/restart", handler.RestartSondaPosition).Methods("PUT")
}

func (handler *SondaHandler) FindSonda(w http.ResponseWriter, r *http.Request) {
	JSON(w, handler.sonda, http.StatusOK)
}

func (handler *SondaHandler) RestartSondaPosition(w http.ResponseWriter, r *http.Request) {
	JSON(w, handler.sonda.Restart(), http.StatusOK)
}

func (handler *SondaHandler) Commands(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		Error(w, err, http.StatusInternalServerError)
		return
	}

	var command Command

	err = json.Unmarshal(b, &command)
	if err != nil {
		log.Println(err.Error())
		Error(w, err, http.StatusInternalServerError)
		return
	}

	err = handler.sonda.Run(command.Movimentos)
	if err != nil {
		Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	JSON(w, handler.sonda, http.StatusOK)
}

func JSON(w http.ResponseWriter, body interface{}, status int) {

	payload, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func Error(w http.ResponseWriter, err interface{}, status int) {

	payload, err := json.Marshal(map[string]interface{}{
		"erro": err,
	})

	if err != nil {
		panic(err)
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
