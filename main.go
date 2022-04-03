package main

import (
	"log"

	"github.com/gorilla/mux"
)

var appName = "bombardier-go"

// Init function, runs before main()
func init() {
	log.SetPrefix("[bombardier-go] ")
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func main() {

	var commonController = controller.CommonControllerFactory()

	r := mux.NewRouter()
	r.Use(utility.LoggingMiddleware)
	//Scommentare solo se si vuole utilizzare i controlli jwt
	r.Use(authentication.JwtAuthMiddleware)
	r.HandleFunc(config.GetInstance().JWTContextPath+"/swagger", commonController.SwaggerConfig).Methods("GET")

}
