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
	"time"
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
		"info":  "info_1",
		"myloc": map[string]interface{}{
			"lat": 1.6,
			"lon": 7.7,
		},
	}

	o2 := map[string]interface{}{
		"info":  "info_2",
		"myloc": map[string]interface{}{
			"lat": 6.8,
			"lon": 7.9,
		},
	}

	o3 := map[string]interface{}{
		"info":  "info_3",
		"myloc": map[string]interface{}{
			"lat": -21.6,
			"lon": 21,
		},
	}

	o4 := map[string]interface{}{
		"info":  "info_4",
		"myloc": map[string]interface{}{
			"lat": 6.8,
			"lon": 10.3,
		},
	}

	o5 := map[string]interface{}{
		"info":  "info_5",
		"myloc": map[string]interface{}{
			"lat": 16.6,
			"lon": 17.7,
		},
	}

	o6 := map[string]interface{}{
		"info":  "info_6",
		"myloc": map[string]interface{}{
			"lat": 11.6,
			"lon": 12.7,
		},
	}

	o7 := map[string]interface{}{
		"info":  "info_7",
		"myloc": map[string]interface{}{
			"lat": 16.69,
			"lon": 7.7,
		},
	}

	o8 := map[string]interface{}{
		"info":  "info_8",
		"myloc": map[string]interface{}{
			"lat": 5.6,
			"lon": 9.7,
		},
	}

	o9 := map[string]interface{}{
		"info":  "info_9",
		"myloc": map[string]interface{}{
			"lat": 1.6,
			"lon": 2.7,
		},
	}

	o10 := map[string]interface{}{
		"info":  "info_10",
		"myloc": map[string]interface{}{
			"lat": 10.6,
			"lon": 19.7,
		},
	}

	o11 := map[string]interface{}{
		"info":  "info_11",
		"myloc": map[string]interface{}{
			"lat": -0.6,
			"lon": -2.7,
		},
	}

	o12 := map[string]interface{}{
		"info":  "info_12",
		"myloc": map[string]interface{}{
			"lat": -16.6,
			"lon": -17.7,
		},
	}

	o13 := map[string]interface{}{
		"info":  "info_13",
		"myloc": map[string]interface{}{
			"lat": -6.6,
			"lon": -1.7,
		},
	}

	o14 := map[string]interface{}{
		"info":  "info_14",
		"myloc": map[string]interface{}{
			"lat": -1.6,
			"lon": -17.7,
		},
	}

	payload, _ := json.Marshal(o1)
	b := bytes.NewBuffer(payload)
	_, e := global.ES.Index("test", b, global.ES.Index.WithDocumentID("1"))
	//fmt.Println(e)

	payload, _ = json.Marshal(o2)
	b = bytes.NewBuffer(payload)
	_, e = global.ES.Index("test", b, global.ES.Index.WithDocumentID("2"))
	//fmt.Println(e)

	payload, _ = json.Marshal(o3)
	b = bytes.NewBuffer(payload)
	_, e = global.ES.Index("test", b, global.ES.Index.WithDocumentID("3"))
	//fmt.Println(e)

	payload, _ = json.Marshal(o4)
	b = bytes.NewBuffer(payload)
	_, e = global.ES.Index("test", b, global.ES.Index.WithDocumentID("4"))
	//fmt.Println(e)

	payload, _ = json.Marshal(o5)
	b = bytes.NewBuffer(payload)
	_, e = global.ES.Index("test", b, global.ES.Index.WithDocumentID("5"))

	payload, _ = json.Marshal(o6)
	b = bytes.NewBuffer(payload)
	_, e = global.ES.Index("test", b, global.ES.Index.WithDocumentID("6"))
	payload, _ = json.Marshal(o7)
	b = bytes.NewBuffer(payload)
	_, e = global.ES.Index("test", b, global.ES.Index.WithDocumentID("7"))
	payload, _ = json.Marshal(o8)
	b = bytes.NewBuffer(payload)
	_, e = global.ES.Index("test", b, global.ES.Index.WithDocumentID("8"))
	payload, _ = json.Marshal(o9)
	b = bytes.NewBuffer(payload)
	_, e = global.ES.Index("test", b, global.ES.Index.WithDocumentID("9"))
	payload, _ = json.Marshal(o10)
	b = bytes.NewBuffer(payload)
	_, e = global.ES.Index("test", b, global.ES.Index.WithDocumentID("10"))
	payload, _ = json.Marshal(o11)
	b = bytes.NewBuffer(payload)
	_, e = global.ES.Index("test", b, global.ES.Index.WithDocumentID("11"))
	payload, _ = json.Marshal(o12)
	b = bytes.NewBuffer(payload)
	_, e = global.ES.Index("test", b, global.ES.Index.WithDocumentID("12"))
	payload, _ = json.Marshal(o13)
	b = bytes.NewBuffer(payload)
	_, e = global.ES.Index("test", b, global.ES.Index.WithDocumentID("13"))
	payload, _ = json.Marshal(o14)
	b = bytes.NewBuffer(payload)
	_, e = global.ES.Index("test", b, global.ES.Index.WithDocumentID("14"))

	time.Sleep(2 * time.Second)

	//fmt.Println(e)
	//payload, _ = json.Marshal(o3)
	//b = bytes.NewBuffer(payload)
	//re, e := global.ES.Update("test", "234", b)
	//
	//fmt.Println(re)
	//fmt.Println(e)

	var response map[string]interface{}
	var buf bytes.Buffer

	sort := map[string]interface{}{
		"sort": map[string]interface{}{
			"_geo_distance": map[string]interface{}{
				"myloc": map[string]interface{}{
					"lat": 10,
					"lon": 10,
				},
				"order": "asc",
				"unit":  "km",
			},
		},
	}

	// We encode from map string-interface into json format.
	if err := json.NewEncoder(&buf).Encode(sort); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	search, e := global.ES.Search(
		global.ES.Search.WithSize(5),
		global.ES.Search.WithIndex("test"), // the index you defined in Elasticsearch
		global.ES.Search.WithBody(&buf),
		global.ES.Search.WithPretty(),
		global.ES.Search.WithFrom(5),
	)

	defer search.Body.Close()

	if e != nil {
		fmt.Println("error getting es searcg: ", e)
	}

	if err := json.NewDecoder(search.Body).Decode(&response); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	//fmt.Println(response)
	result := response["hits"].(map[string]interface{})["hits"]
	fmt.Println(result)

	//_, e = global.ES.Delete("test", "234")
}
