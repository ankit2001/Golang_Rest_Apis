package router

import (
	"Golang_Rest_Apis/router/routes"

	"github.com/gorilla/mux"
)

//New is function
func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetUpRoutesWithMiddlewares(r)
}