package models

import (
	"crypto/aes"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"onestep/configs"
	"os"
	"strings"

	"github.com/gorilla/securecookie"
	"github.com/joho/godotenv"
)

type UserInfo struct {
	ID string `json:"id"`
	PW string `json:"pw"`
}

func CreateUser(userInfo UserInfo) string {
	db := configs.CreateConnection()
	defer db.Close()

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	key := os.Getenv("BLOCK_KEY")

	block, _ := aes.NewCipher([]byte(key))
	encryptPW := make([]byte, 16)

	userInfo.PW = userInfo.PW + strings.Repeat(os.Getenv("CHAR_TO_FILL_BLANK"), 16-len(userInfo.PW))
	block.Encrypt(encryptPW, []byte(userInfo.PW))

	_, err := db.Exec(`INSERT INTO userinfo VALUES (?, ?)`, userInfo.ID, hex.EncodeToString(encryptPW))
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", userInfo.ID)
	return userInfo.ID
}

func Login(userInfo UserInfo) (*http.Cookie, error) {
	db := configs.CreateConnection()
	defer db.Close()

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	var user UserInfo
	row := db.QueryRow("SELECT * FROM userinfo WHERE id = ?", userInfo.ID)

	if err := row.Scan(&user.ID, &user.PW); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s: no such user", userInfo.ID)
		}
		return nil, fmt.Errorf("%s: %v", userInfo.ID, err)
	}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	key := os.Getenv("BLOCK_KEY")

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	encryptPW := make([]byte, 16)
	userInfo.PW = userInfo.PW + strings.Repeat(os.Getenv("CHAR_TO_FILL_BLANK"), 16-len(userInfo.PW))
	block.Encrypt(encryptPW, []byte(userInfo.PW))

	if hex.EncodeToString(encryptPW) != user.PW {
		return nil, fmt.Errorf("%s", "invalid password")
	}

	value := map[string]string{
		"name": user.ID,
	}

	var cookieHandler = securecookie.New(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32))

	encoded, err := cookieHandler.Encode("session", value)

	if err != nil {
		return nil, fmt.Errorf("%s", "undetected error")
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: encoded,
		Path:  "/",
	}

	return cookie, nil
}
