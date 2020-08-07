package server

import (
	"Golang_Rest_Apis/auto"
	"Golang_Rest_Apis/config"
	"Golang_Rest_Apis/router"
	"fmt"
	"log"
	"net/http"
)

func Run () {
	config.Load()
	auto.Load()
	fmt.Printf("\n\tListening.......[::]:%d \n", config.PORT)

	Listen(config.PORT)
}

func Listen(port int) {
	r := router.New()
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Fatal("error is : ", err)
	}
}