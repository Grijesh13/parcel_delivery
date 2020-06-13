package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"parcelDelivery/dto"
)

type IParcels interface {
	AddParcel(parcel *dto.Parcel) error
}

type ParcelsImpl struct {
	// db client
	DB *sql.DB
}

func (impl *ParcelsImpl) AddParcel(parcel *dto.Parcel) error {
	sqlQuery := "INSERT INTO parcel_delivery.parcels VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )"
	stmt, err := impl.DB.Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		fmt.Println("error preparing the add parcel sql statement")
		return errors.New("error preparing the add parcel sql statement")
	}
	_, err = stmt.Exec(parcel.ID, parcel.UserName, parcel.Note, parcel.Length, parcel.Breadth, parcel.Height,
		parcel.Weight, parcel.Category, parcel.SourceAddress, parcel.DestinationAddress,
		parcel.SourceLatitude, parcel.SourceLongitude, parcel.DestinationLatitude,
		parcel.DestinationLongitude, parcel.CreatedAt)
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
