package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/restapi/routes"
)

func main() {
	fmt.Println("Successfully Connected..!")
	router := mux.NewRouter()
	routes.BasicCrudRoutes(router)

	log.Fatal(http.ListenAndServe(":8000", router))
}
