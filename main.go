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
	router.HandleFunc("/signup", apis.SignUp).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main()  {
	startHTTPServer()
}
