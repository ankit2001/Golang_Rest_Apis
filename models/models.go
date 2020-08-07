package models

import (
	"Golang_Rest_Apis/security"
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Email     string    `gorm:"size:50;not null;unique" json:"email"`
	Password  string    `gorm:"size:60;not null" json:"password"`
	CreatedAt time.Time `gorm:"" json:"created_at"`
	UpdatedAt time.Time `gorm:"" json:"updated_at"`
}