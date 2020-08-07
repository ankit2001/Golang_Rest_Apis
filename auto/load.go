package auto

import (
	"Golang_Rest_Apis/database"
	"Golang_Rest_Apis/models"
	"log"
)
func Load()  {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("this is an error :", err)
	}
	defer db.Close()
	err = db.Debug().DropTableIfExists(&models.User{}, &models.Post{}, &models.ExpertiseService{}, &models.PujaService{}, &models.OtherserviceService{}, models.PujaServiceDuration{}, models.PujaServicePrice{}, models.Availability{}, models.DateTime{}, models.Verification{}).Error
	if err != nil {
		log.Fatal("this is an error :", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.ExpertiseService{}, &models.PujaService{}, &models.OtherserviceService{}, models.PujaServiceDuration{}, models.PujaServicePrice{}, models.Availability{}, models.DateTime{}, models.Verification{}, models.BookPuja{}, models.FrontPageLoader{}, models.PujaServiceVideo{}).Error
	if err != nil {
		 log.Fatal("error occured : ", err )
	 }

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id","users(id)","cascade","cascade").Error
	if err != nil {
		log.Fatal("error occured : ", err )
	}

}