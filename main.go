package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"parcelDelivery/apis"

	"github.com/gorilla/mux"
)

func startHTTPServer() {
	fmt.Println("starting http server")
	router := mux.NewRouter().StrictSlash(true)
	// USER REGISTRATION APIs
	router.HandleFunc("/signUp", apis.SignUp).Methods("POST")
	//  PARCEL APIs
	router.HandleFunc("/addParcel", apis.AddParcel).Methods("POST")
	router.HandleFunc("/getParcelsForUser", apis.GetParcelsForUser).Methods("POST")
	router.HandleFunc("/getParcels", apis.GetParcels).Methods("POST")
	// TRAVEL APIs
	router.HandleFunc("/addTravel", apis.AddTravel).Methods("POST")
	router.HandleFunc("/getTravelsForUser", apis.GetTravelsForUser).Methods("POST")
	router.HandleFunc("/getTravels", apis.GetTravels).Methods("POST")
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode("health check")
	}).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	startHTTPServer()
}
