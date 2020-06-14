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
	router.HandleFunc("/addparcel", apis.AddParcel).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main()  {
	//f()
	startHTTPServer()

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
	//o1 := dto.ESParcel{
	//	MyLoc: dto.Loc{
	//		Lat: 12.12,
	//		Long: 13.13,
	//	},
	//	Info: "info_1",
	//}
	//o2 := dto.ESParcel{
	//	MyLoc: dto.Loc{
	//		Lat: 1.1,
	//		Long: 3.3,
	//	},
	//	Info: "info_2",
	//}
	//o3 := dto.ESParcel{
	//	MyLoc: dto.Loc{
	//		Lat: 6.6,
	//		Long: 7.8,
	//	},
	//	Info: "info_3",
	//}

	o1 := map[string]interface{}{
		"asd":  "info_1",
		"myloc": map[string]interface{}{
			"lat": 6.6,
			"lon": 7.7,
		},
	}

	o2 := map[string]interface{}{
		"asd":  "info_2",
		"myloc": map[string]interface{}{
			"lat": 6.6,
			"lon": 7.7,
		},
	}

	o3 := map[string]interface{}{
		"doc": map[string]interface{}{
			"asd":  "info_3",
			"myloc": map[string]interface{}{
				"lat": -21.6,
				"lon": 180,
			},
		},
	}

	payload, _ := json.Marshal(o1)
	b := bytes.NewBuffer(payload)
	_, e := global.ES.Index("test", b, global.ES.Index.WithDocumentID("123"))
	fmt.Println(e)

	payload, _ = json.Marshal(o2)
	b = bytes.NewBuffer(payload)
	_, e = global.ES.Index("test", b, global.ES.Index.WithDocumentID("234"))
	fmt.Println(e)

	payload, _ = json.Marshal(o3)
	b = bytes.NewBuffer(payload)
	re, e := global.ES.Update("test", "234", b)

	fmt.Println(re)
	fmt.Println(e)

	//_, e = global.ES.Delete("test", "234")
}
