package models

import (
	"database/sql"
	"fmt"
	"log"
	"onestep/configs"
)

type UserInfo struct {
	ID string `json:"id"`
	PW string `json:"pw"`
}

func CreateUser(userinfo UserInfo) string {
	db := configs.CreateConnection()
	defer db.Close()

	_, err := db.Exec(`INSERT INTO userinfo VALUES (?, ?)`, userinfo.ID, userinfo.PW)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", userinfo.ID)
	return userinfo.ID
}

func GetAllUsers() ([]UserInfo, error) {
	db := configs.CreateConnection()
	defer db.Close()

	var userInfos []UserInfo

	rows, err := db.Query("SELECT * FROM userinfo")

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rows.Close()

	for rows.Next() {
		var userInfo UserInfo

		// unmarshal the row object to stock
		err = rows.Scan(&userInfo.ID, &userInfo.PW)
		fmt.Print(userInfo)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the stock in the stocks slice
		userInfos = append(userInfos, userInfo)
	}

	return userInfos, err
}

func GetUser(id string) (UserInfo, error) {
	db := configs.CreateConnection()
	defer db.Close()

	var userInfo UserInfo

	row := db.QueryRow("SELECT * FROM userinfo WHERE id = ?", id)

	if err := row.Scan(&userInfo.ID, &userInfo.PW); err != nil {
		if err == sql.ErrNoRows {
			return userInfo, fmt.Errorf("albumsById %s: no such user", id)
		}
		return userInfo, fmt.Errorf("albumsById %s: %v", id, err)
	}

	return userInfo, nil
}
