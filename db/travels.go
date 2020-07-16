package db

import (
	"database/sql"
	"errors"
	"fmt"
	"parcelDelivery/dto"

	"github.com/go-sql-driver/mysql"
)

// ITravels ...
type ITravels interface {
	AddTravel(travel *dto.Travel) error
	GetTravels(username string) []*dto.Travel
}

// TravelsImpl ...
type TravelsImpl struct {
	// db client
	DB *sql.DB
}

// GetTravels ...
func (impl *TravelsImpl) GetTravels(username string) []*dto.Travel {
	sqlQuery := "SELECT id, username, note, mode, src_address, dest_address, src_lat, src_long, dest_lat, dest_long, created_at, status, start_date, end_date, completed_at FROM parcel_delivery.travels WHERE username = ? ORDER BY created_at DESC"
	stmt, err := impl.DB.Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		fmt.Printf("error preparing the get travels user sql statement")
		return nil
	}
	res, e := stmt.Query(username)
	if e != nil {
		fmt.Printf("error getting user travels with error = %s", e.Error())
		return nil
	}
	var travels []*dto.Travel
	for res.Next() {
		var travel dto.Travel
		// need to map each attribute retrieved else people is empty
		_ = res.Scan(&travel.ID, &travel.UserName, &travel.Note, &travel.Mode, &travel.SourceAddress,
			&travel.DestinationAddress, &travel.SourceLatitude, &travel.SourceLongitude,
			&travel.DestinationLatitude, &travel.DestinationLongitude, &travel.CreatedAt,
			&travel.Status, &travel.StartDate, &travel.EndDate, &travel.CompletedAt)
		travels = append(travels, &travel)
	}
	return travels
}

// AddTravel ...
func (impl *TravelsImpl) AddTravel(travel *dto.Travel) error {
	sqlQuery := "INSERT INTO parcel_delivery.travels VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )"
	stmt, err := impl.DB.Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		fmt.Println("error preparing the add travel sql statement")
		return errors.New("error preparing the add travel sql statement")
	}
	_, err = stmt.Exec(travel.ID, travel.UserName, travel.Note, travel.Mode,
		travel.SourceAddress, travel.DestinationAddress, travel.SourceLatitude,
		travel.SourceLongitude, travel.DestinationLatitude, travel.DestinationLongitude,
		travel.CreatedAt, travel.Status, travel.StartDate, travel.EndDate, nil)
	sqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		if sqlErr.Number == 1062 {
			fmt.Println("error adding travel : duplicate key")
			return errors.New("duplicate key")
		}
		// foreign key constraint
		if sqlErr.Number == 1452 {
			fmt.Println("error adding travel : username not present")
			return errors.New("username not present")
		}
	}
	if err != nil {
		fmt.Println("error adding travel with error:", err.Error())
		return errors.New("error executing the add travel sql statement")
	}
	return nil
}
