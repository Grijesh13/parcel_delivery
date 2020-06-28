package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"parcelDelivery/apis"
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
	router.HandleFunc("/getTravelsForUser", apis.GetParcelsForUser).Methods("POST")
	router.HandleFunc("/getTravels", apis.GetTravels).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main()  {
	startHTTPServer()
}
