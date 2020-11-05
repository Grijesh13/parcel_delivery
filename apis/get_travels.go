package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"parcelDelivery/global"
	dto "parcelDelivery/request_dto"
	responseDto "parcelDelivery/dto/response"
	"github.com/olivere/elastic/v7"
	"strconv"
	"context"
)

// GetTravels ...
func GetTravels(w http.ResponseWriter, r *http.Request) {
	var newEvent *dto.LazyLoad
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

	// var response map[string]interface{}
	// var buf bytes.Buffer

	if newEvent.SrcDistance == 0 {
		// set to default
		newEvent.SrcDistance = 10
	}

	var filter []map[string]interface{}
	filter = append(filter, map[string]interface{}{
		"geo_distance": map[string]interface{}{
			"distance": strconv.Itoa(newEvent.SrcDistance) + "km",
			"mysrc": map[string]interface{}{
				"lat": newEvent.SrcLatitude,
				"lon": newEvent.SrcLongitude,
			},
		},
	})

	if newEvent.DestDistance == 0 {
		// set to default
		newEvent.DestDistance = 10
	}

	if newEvent.DestGiven {
		filter = append(filter, map[string]interface{}{
			"geo_distance": map[string]interface{}{
				"distance": strconv.Itoa(newEvent.DestDistance) + "km",
				"mydest": map[string]interface{}{
					"lat": newEvent.DestLatitude,
					"lon": newEvent.DestLongitude,
				},
			},
		})
	}

	sort := map[string]interface{}{
		"from": newEvent.From,
		"size": newEvent.Many,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": filter,
			},
		},
		"sort": map[string]interface{}{
			"_geo_distance": map[string]interface{}{
				"mysrc": map[string]interface{}{
					"lat": newEvent.SrcLatitude,
					"lon": newEvent.SrcLongitude,
				},
				"order": "asc",
				"unit":  "km",
			},
		},
	}

	query, _ := json.Marshal(sort)
	search, searchErr :=	global.ES2.Search().
															Index(global.ESTravelIndex).
															Source(string(query)).
															Do(context.Background())

	if searchErr != nil {
		fmt.Println("error preparing es search for travels query:", searchErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("problem")
		return
	}

	travels, err := decodeTravels(search)
	if err != nil {
		fmt.Println("error decoding es search for travels query:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("problem")
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(travels)
}

func decodeTravels(res *elastic.SearchResult) ([]*responseDto.Travel, error) {
	if res == nil || res.TotalHits() == 0 {
		return nil, nil
	}
	var travels []*responseDto.Travel
	for _, hit := range res.Hits.Hits {
		travel := new(responseDto.Travel)
		if err := json.Unmarshal(hit.Source, travel); err != nil {
			return nil, err
		}
		travel.ID = hit.Id
		if len(hit.Sort) > 0 {
			travel.Distance = hit.Sort[0].(float64)
		}
		travels = append(travels, travel)
	}
	return travels, nil
}
