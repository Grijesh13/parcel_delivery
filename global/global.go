package global

import (
	"context"
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"parcelDelivery/dto"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/aws/credentials"
	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/olivere/elastic/v7"
	"github.com/sha1sum/aws_signing_client"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser     = "admin"
	dbPassword = "db$123456"
	dbHost     = "database-1.cisb0uvrjg3i.ap-southeast-1.rds.amazonaws.com"
	dbPort     = "3306"
)

const (
	esUser = "elastic"
	esPsw  = "changeme"
)

// DB is the SQL client
var DB = initSQL()

// ES is the Elastic Search client
var ES2 = initES()

func initSQL() *sql.DB {
	// open up a database connection
	db, err := sql.Open("mysql",
		dbUser+":"+dbPassword+"@("+dbHost+":"+dbPort+")/")
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

func initES() *elastic.Client {
	sess := session.Must(session.NewSession())
	creds := credentials.NewCredentials(&ec2rolecreds.EC2RoleProvider{
		Client: ec2metadata.New(sess),
	})
	signer := v4.NewSigner(creds)
	awsClient, err := aws_signing_client.New(signer, nil, "es", "ap-southeast-1")
	if err != nil {
		panic("Elastic failed to initialize, AWS Error: " + err.Error())
	}
	client, err := elastic.NewClient(
		elastic.SetURL("https://search-dev-parcel-delivery-6egqf5z3c3z7borcxhoxrv7le4.ap-southeast-1.es.amazonaws.com"),
		elastic.SetScheme("https"),
		elastic.SetHttpClient(awsClient),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetBasicAuth(esUser, esPsw),
	)
	if err != nil {
		panic("Elastic failed to initialize , Err: " + err.Error())
	}

	// Ping the Elasticsearch server to get e.g. the version number
	// info, code, err := client.Ping("https://search-dev-parcel-delivery-6egqf5z3c3z7borcxhoxrv7le4.ap-southeast-1.es.amazonaws.com").Do(context.Background())
	// if err != nil {
	// 	// Handle error
	// 	panic("Ping Error" + err.Error())
	// }

	travelSetting := dto.ESTravelSetting{
		Settings: dto.Setting{
			Shards:   1,
			Replicas: 0,
		},
		Mappings: dto.TravelMapping{
			Properties: dto.TravelProperty{
				MySrcLoc: dto.Geo{
					Type: "geo_point",
				},
				MyDestLoc: dto.Geo{
					Type: "geo_point",
				},
				StartDate: dto.ESDate{
					Type:   "date",
					Format: "yyyy-MM-dd",
				},
				EndDate: dto.ESDate{
					Type:   "date",
					Format: "yyyy-MM-dd",
				},
			},
		},
	}

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(ESTravelIndex).Do(context.Background())
	if err != nil {
		// Handle error
		panic("IndexExists panic" + err.Error())
	}
	if !exists {
		payload, _ := json.Marshal(travelSetting)

		createIndex, err := client.CreateIndex(ESTravelIndex).Body(string(payload)).Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			println("Not Ack for Travel index")
		}
		println("travel index created", createIndex.Index)
	}

	parcelSetting := dto.ESParcelSetting{
		Settings: dto.Setting{
			Shards:   1,
			Replicas: 0,
		},
		Mappings: dto.ParcelMapping{
			Properties: dto.ParcelProperty{
				MySrcLoc: dto.Geo{
					Type: "geo_point",
				},
				MyDestLoc: dto.Geo{
					Type: "geo_point",
				},
				PickUpStart: dto.ESDate{
	        Type:   "date",
        	Format: "yyyy-MM-dd",
      	},
				PickUpEnd: dto.ESDate{
	        Type:   "date",
        	Format: "yyyy-MM-dd",
      	},
			},
		},
	}

	// Use the IndexExists service to check if a specified index exists.
	exists, err = client.IndexExists(ESParcelIndex).Do(context.Background())
	if err != nil {
		// Handle error
		panic("IndexExists panic" + err.Error())
	}
	if !exists {
		payload, _ := json.Marshal(parcelSetting)

		createIndex, err := client.CreateIndex(ESParcelIndex).Body(string(payload)).Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			println("Not Ack for Parcel index")
		}
		println("parcel index created", createIndex.Index)
	}

	return client
}
