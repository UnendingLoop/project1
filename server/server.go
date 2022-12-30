package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"project/config"
	"project/result"

	"github.com/gorilla/mux"
)

func listenAndServeHTTP() {
	router := mux.NewRouter()
	router.HandleFunc("/api", handleAPI).Methods("GET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf("GET: Full info %v", config.GlobalConfig.Addr)

	log.Fatal(http.ListenAndServe(config.GlobalConfig.Addr, router))
}

func handleAPI(w http.ResponseWriter, r *http.Request) {
	resultT := result.ResultT{Status: false, Error: "Error on collect data sim"}

	resultSetT := result.GetResultData()
	if result.CheckResult(resultSetT) {
		resultT.Status = true
		resultT.Data = resultSetT
		resultT.Error = ""
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	response, err := json.Marshal(resultT)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte{})
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}

func handleMMS(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(config.GlobalConfig.MMSFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte{})
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func handleSupport(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(config.GlobalConfig.SupportFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte{})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func handleIncident(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(config.GlobalConfig.IncidentFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte{})
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(data)
}

func StartServer() {
	listenAndServeHTTP()
}
