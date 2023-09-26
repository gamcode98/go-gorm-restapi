package main

import (
	"net/http"

	"github.com/gamcode98/go-gorm-restapi/db"
	"github.com/gamcode98/go-gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
