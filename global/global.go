package global

import (
	"bytes"
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	_ "github.com/go-sql-driver/mysql"
	"parcelDelivery/dto"
	"time"
)

const (
	dbUser     = "root"
	dbPassword = "JustDoIt1308!"
	dbHost     = "localhost"
	dbPort     = "3306"
)

// DB is the SQL client
var DB = initSQL()

// ES is the Elastic Search client
var ES = initES()

func initES() *elasticsearch.Client {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		fmt.Println("ES Connection Failed!!")
		panic(err)
	}
	setting := dto.ESSetting{
		Settings: dto.Setting{
			Shards: 1,
			Replicas: 0,
		},
		Mappings: dto.Mapping{
			Properties: dto.Property{
				MyLoc: dto.Geo{
					Type: "geo_point",
				},
			},
		},
	}
	payload, _ := json.Marshal(setting)
	b := bytes.NewBuffer(payload)
	_, err = es.Indices.Create(ESIndex, es.Indices.Create.WithBody(b))
	if err != nil {
		fmt.Println("error creating shards:", err)
		panic(err)
	}
	return es
}

func initSQL() *sql.DB {
	// open up a database connection
	db, err:= sql.Open("mysql",
		dbUser + ":" + dbPassword + "@(" + dbHost + ":" + dbPort + ")/")
	if err != nil {
		fmt.Println("Connection Failed!!")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Ping Failed!!")
		panic(err)
	}
	// set configs
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Second * 10)
	return db
}
