package database

import (
	"Golang_Rest_Apis/config"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)
	if err != nil {
		return nil, err
	}
	return db, nil

}