package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"parcelDelivery/apis"
	"parcelDelivery/global"
)

func startHTTPServer() {
	fmt.Println("starting http server")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/signup", apis.SignUp).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main()  {
	f()
	//startHTTPServer()

	//ParcelImpl := db.ParcelsImpl{
	//	DB : global.DB,
	//}
	//currentTime := time.Now()
	//currentTS := currentTime.Format("2006-01-02 15:04:05")
	//err := ParcelImpl.AddParcel(&dto.Parcel{
	//	ID: "4343",
	//	UserName: "444",
	//	Note: "note",
	//	Length: 12,
	//	Breadth: 12,
	//	Height: 12,
	//	Weight:12,
	//	Category: "sex toy",
	//	SourceAddress: "wqfwegrwae",
	//	DestinationAddress: "qwfqewfqeg",
	//	SourceLatitude: 12.4343,
	//	SourceLongitude: 23.13314,
	//	DestinationLatitude: 32.3523,
	//	DestinationLongitude: 32.2332,
	//	CreatedAt: currentTS,
	//})
	//fmt.Println(err)
}

// adding entries to es
func f() {
	o1 := esObj{
		MyLoc: loc{
			Lat: 12.12,
			Long: 13.13,
		},
		info: "123123123",
	}
	o2 := esObj{
		MyLoc: loc{
			Lat: 12.144,
			Long: 13.1312,
		},
		info: "456456",
	}

	payload, _ := json.Marshal(o1)
	b := bytes.NewBuffer(payload)
	_, e := global.ES.Index("test", b)
	fmt.Println(e)

	payload, _ = json.Marshal(o2)
	b = bytes.NewBuffer(payload)
	_, e = global.ES.Index("test", b)
	fmt.Println(e)
}

type esObj struct {
	MyLoc loc   `json:"myloc"`
	info string `json:"info"`
}

type loc struct {
	Lat float64 `json:"lat"`
	Long float64 `json:"lon"`
}
