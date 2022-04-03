package main

import (
	"log"

	"github.com/davideginna/bombardier-go/controller"
	"github.com/davideginna/bombardier-go/utility"
	"github.com/gorilla/mux"
)

var appName = "bombardier-go"

// Init function, runs before main()
func init() {
	log.SetPrefix("[bombardier-go] ")
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func main() {
	log.Println("Starting ", appName)

	var commonController = controller.CommonControllerFactory()

	r := mux.NewRouter()
	r.Use(utility.LoggingMiddleware)
	//Scommentare solo se si vuole utilizzare i controlli jwt
	r.HandleFunc("/bombardier", commonController.Hello).Methods("GET")

}
