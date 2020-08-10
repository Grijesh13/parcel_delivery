package apis

import (
	"encoding/json"
	"fmt"
	"time"
	"io/ioutil"
	"net/http"
	"parcelDelivery/global"
	dto "parcelDelivery/request_dto"
	responseDto "parcelDelivery/dto/response"
	"github.com/olivere/elastic/v7"
	"strconv"
	"context"
)

// GetParcels ...
func GetParcels(w http.ResponseWriter, r *http.Request) {
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

	sDate, _ := time.Parse("2006-01-02 15:04:05", "2020-08-06 15:04:05")
	uDate := sDate.Format("2006-01-02")

	eDate, _ := time.Parse("2006-01-02 15:04:05", "2020-08-07 15:04:05")
	uuDate := eDate.Format("2006-01-02")

	fmt.Printf("%v", sDate)
	fmt.Printf("%v", uDate)

	var must []map[string]interface{}
	must = append(must, map[string]interface{}{
		"range": map[string]interface{}{
			"pick_up_start": map[string]interface{}{
				"gte": uDate,
				"lte": uuDate,
				"format": "yyyy-MM-dd",
			},
		},
	})

	filter = append(filter, map[string]interface{}{
		"bool": map[string]interface{}{
			"must": must,
    },
	})

	sort := map[string]interface{}{
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
															Index(global.ESParcelIndex).
															Source(string(query)).
															From(newEvent.From).
															Size(newEvent.Many).
															Do(context.Background())

	if searchErr != nil {
		fmt.Println("error preparing es search for parcels query:", searchErr.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("problem")
		return
	}

	parcels, err := decodeParcels(search)
	if err != nil {
		fmt.Println("error decoding es search for parcels query:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("problem")
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(parcels)
}

func decodeParcels(res *elastic.SearchResult) ([]*responseDto.Parcel, error) {
	if res == nil || res.TotalHits() == 0 {
		return nil, nil
	}
	var parcels []*responseDto.Parcel
	for _, hit := range res.Hits.Hits {
		parcel := new(responseDto.Parcel)
		if err := json.Unmarshal(hit.Source, parcel); err != nil {
			return nil, err
		}
		parcel.ID = hit.Id
		if len(hit.Sort) > 0 {
			parcel.Distance = hit.Sort[0].(float64)
		}
		parcels = append(parcels, parcel)
	}
	return parcels, nil
}
