package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

var (
	PORT = 0
	SECRETKEY []byte
	DBURL = ""
	DBDRIVER =""
	STOREURL []byte
)

func Load() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error is : ", err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT")) 

	if err != nil {
		PORT = 8080
	}
	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"),os.Getenv("DB_NAME"))

	SECRETKEY = []byte(os.Getenv("API_SECRET"))

	STOREURL = []byte(os.Getenv("STORE_URL"))

	
}