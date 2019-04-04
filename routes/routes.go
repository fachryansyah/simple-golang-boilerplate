package routes

import (
	"github.com/gorilla/mux"
	"simple-golang-boilerplate/app/http/controller"
)

func SetRoutes() *mux.Router {

	router := mux.NewRouter()

	/*
		Masukkan Route kamu disini.
	*/
	router.HandleFunc("/user/get", controller.GetUser).Methods("GET")

	return router
}