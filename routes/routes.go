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
	router.HandleFunc("/from-auth/user/get", controller.LoginUser).Methods("GET")

	return router
}