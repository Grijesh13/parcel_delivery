package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"parcelDelivery/dto"

	"github.com/go-sql-driver/mysql"
)

// IParcels ...
type IParcels interface {
	AddParcel(parcel *dto.Parcel) error
	GetParcels(username string) []*dto.Parcel
}

// ParcelsImpl ...
type ParcelsImpl struct {
	// db client
	DB *sql.DB
}

// GetParcels ...
func (impl *ParcelsImpl) GetParcels(username string) []*dto.Parcel {
	sqlQuery := "SELECT id, username, note, items, src_address, dest_address, src_lat, src_long, dest_lat, dest_long, pick_up_start, pick_up_end, created_at, status, price, is_negotiable, completed_at FROM parcel_delivery.parcels WHERE username = ? ORDER BY created_at DESC"
	stmt, err := impl.DB.Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		fmt.Printf("error preparing the get parcels user sql statement")
		return nil
	}
	res, e := stmt.Query(username)
	if e != nil {
		fmt.Printf("error getting user parcels with error = %s", e.Error())
		return nil
	}
	var parcels []*dto.Parcel
	for res.Next() {
		var parcel dto.Parcel
		// need to map each attribute retrieved else parcel is empty
		_ = res.Scan(&parcel.ID, &parcel.UserName, &parcel.Note, &parcel.SQLItems,
			&parcel.SourceAddress, &parcel.DestinationAddress,
			&parcel.SourceLatitude, &parcel.SourceLongitude,
			&parcel.DestinationLatitude, &parcel.DestinationLongitude,
			&parcel.PickUpStart, &parcel.PickUpEnd, &parcel.CreatedAt, &parcel.Status,
			&parcel.Price, &parcel.IsNegotiable, &parcel.CompletedAt)
		var items *[]dto.Item
		er := json.Unmarshal([]byte(parcel.SQLItems), &items)
		if er != nil {
			fmt.Printf("error unmarshalling sql_items for parcels with error = %s", er.Error())
			return nil
		}
		parcels = append(parcels, &parcel)
	}
	return parcels
}

// AddParcel ...
func (impl *ParcelsImpl) AddParcel(parcel *dto.Parcel) error {
	sqlQuery := "INSERT INTO parcel_delivery.parcels VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )"
	stmt, err := impl.DB.Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		fmt.Println("error preparing the add parcel sql statement")
		return errors.New("error preparing the add parcel sql statement")
	}
	serItems, _ := json.Marshal(parcel.Items)
	parcel.SQLItems = string(serItems)
	_, err = stmt.Exec(parcel.ID, parcel.UserName, parcel.Note, parcel.SQLItems,
		parcel.SourceAddress, parcel.DestinationAddress,
		parcel.SourceLatitude, parcel.SourceLongitude,
		parcel.DestinationLatitude, parcel.DestinationLongitude,
		parcel.PickUpStart, parcel.PickUpEnd, parcel.CreatedAt, parcel.Status,
		parcel.Price, parcel.IsNegotiable, nil)
	sqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		if sqlErr.Number == 1062 {
			fmt.Println("error adding parcel : duplicate key")
			return errors.New("duplicate key")
		}
		// foreign key constraint
		if sqlErr.Number == 1452 {
			fmt.Println("error adding parcel : username not present")
			return errors.New("username not present")
		}
	}
	if err != nil {
		fmt.Println("error adding parcel with error:", err.Error())
		return errors.New("error executing the add parcel sql statement")
	}
	return nil
}
