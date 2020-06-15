package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"parcelDelivery/global"
	dto "parcelDelivery/request_dto"
)

func GetParcels(w http.ResponseWriter, r *http.Request) {
	var newEvent *dto.LazyParcels
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("error in reading data")
		return
	}
	err = json.Unmarshal(reqBody, &newEvent)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode("error in unmarshalling data")
		return
	}

	if newEvent.Many == 0 {
		// set to default
		newEvent.Many = global.DefaultMany
	}

	var response map[string]interface{}
	var buf bytes.Buffer

	sort := map[string]interface{}{
		"sort": map[string]interface{}{
			"_geo_distance": map[string]interface{}{
				"myloc": map[string]interface{}{
					"lat": newEvent.Latitude,
					"lon": newEvent.Longitude,
				},
				"order": "asc",
				"unit":  "km",
			},
		},
	}

	// encode from map string-interface into json format
	if err := json.NewEncoder(&buf).Encode(sort); err != nil {
		fmt.Println("error in encoding query:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("problem")
		return
	}

	search, searchErr := global.ES.Search(
		global.ES.Search.WithSize(newEvent.Many),
		global.ES.Search.WithIndex(global.ESIndex), // the index defined in Elasticsearch
		global.ES.Search.WithBody(&buf),
		global.ES.Search.WithPretty(),
		global.ES.Search.WithFrom(newEvent.From),
	)

	if searchErr != nil {
		fmt.Println("error preparing es search for parcels query:", searchErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("problem")
		return
	}
	defer search.Body.Close()

	if err := json.NewDecoder(search.Body).Decode(&response); err != nil {
		fmt.Println("error parsing the response body for es search:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("problem")
		return
	}

	result := response["hits"].(map[string]interface{})["hits"]
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}
