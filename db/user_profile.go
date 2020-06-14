package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"parcelDelivery/dto"
)

type IUserProfile interface {
	AddUser(profile *dto.User) error
	GetUserData(username string) *dto.User
	UpdatePassword(profile *dto.User) error
	UpdatePhone(profile *dto.User) error
}

type UserProfileImpl struct {
	// db client
	DB *sql.DB
}

func (up *UserProfileImpl) UpdatePassword(profile *dto.User) error {
	sqlQuery := "UPDATE parcel_delivery.people SET password = ? WHERE username = ?"
	stmt, err := up.DB.Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		fmt.Printf("error preparing the update password user sql statement")
		return errors.New("error preparing the update password user sql statement")
	}
	res, e := stmt.Exec(profile.Password, profile.UserName)
	if e != nil {
		fmt.Printf("error updating user password with error = %s", e.Error())
		return errors.New("error executing the updating password user sql statement")
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("no user with the given username:", profile.UserName)
	}
	return nil
}

func (up *UserProfileImpl) UpdatePhone(profile *dto.User) error {
	sqlQuery := "UPDATE parcel_delivery.people SET country_code = ?, phone_number = ? WHERE username = ?"
	stmt, err := up.DB.Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		fmt.Printf("error preparing the update phone user sql statement")
		return errors.New("error preparing the update phone user sql statement")
	}
	res, e := stmt.Exec(profile.CountryCode, profile.Phone, profile.UserName)
	if e != nil {
		fmt.Printf("error updating user phone with error = %s", e.Error())
		return errors.New("error executing the updating phone user sql statement")
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		fmt.Println("no user with the given username:", profile.UserName)
	}
	return nil
}

func (up *UserProfileImpl) GetUserData(username string) *dto.User {
	sqlQuery := "SELECT username, password, country_code, phone_number, created_at FROM parcel_delivery.people where username = ?"
	stmt, err := up.DB.Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		fmt.Printf("error preparing the add user sql statement")
		return nil
	}
	res, e := stmt.Query(username)
	if e != nil {
		fmt.Printf("error getting user profile with error = %s", e.Error())
		return nil
	}
	var people dto.User
	if res.Next() {
		// need to map each attribute retrieved else people is empty
		_ = res.Scan(&people.UserName, &people.Password, &people.CountryCode, &people.Phone, &people.CreatedAt)
		return &people
	}
	return nil
}

func (up *UserProfileImpl) AddUser(profile *dto.User) error {
	sqlQuery := "INSERT INTO parcel_delivery.people VALUES ( ?, ?, ?, ?, ? )"
	stmt, err := up.DB.Prepare(sqlQuery)
	defer closeStmt(stmt)
	if err != nil {
		fmt.Printf("error preparing the add user sql statement")
		return errors.New("error preparing the add user sql statement")
	}
	_, err = stmt.Exec(profile.UserName, profile.Password,
		profile.CountryCode, profile.Phone, profile.CreatedAt)
	sqlErr, ok := err.(*mysql.MySQLError)
	if ok && sqlErr.Number == 1062 {
		fmt.Printf("error adding user profile : duplicate key")
		return errors.New("duplicate key")
	}
	if err != nil {
		fmt.Printf("error adding user profile with error = %s", err.Error())
		return errors.New("error executing the add user sql statement")
	}
	return nil
}

func closeStmt(stmt *sql.Stmt) {
	if stmt != nil {
		_ = stmt.Close()
	}
}
