package main

import (
	"log"
	"fmt"
	"net/http"
	"simple-golang-boilerplate/routes"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.Handle("/", routes.SetRoutes())
	fmt.Println("Backend started on port : 9000")
	log.Fatal(http.ListenAndServe(":9000", routes.SetRoutes()))
}