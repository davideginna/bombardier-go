//go:generate genny -in=template-main.go -out=main.go gen "contextPath=${tIPO} Template=${TIPO} template=${tIPO} base-path=${MS_PATH}"
// Package classification Template ms.
// @title Template API
// @version 1.0.0
// @description Microservizio per la gestione della risorsa Template
// @termsOfService

// @license.name ' '
// @contact.name API Support
// @contact.url http://www.regione.toscana
// @contact.email supporto.cochise@tdnet.it

// @host localhost:9091
// @BasePath /cochise/base-path

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
)

// Init function, runs before main()
func init() {
	log.SetPrefix(constants.PrefixLog)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
}

func main() {
	var config = config.GetInstance()
	log.Println("Starting ", constants.AppName)

	r := mux.NewRouter()
	//Attivazione middleware logging
	r.Use(middleware.LoggingMiddleware)

	//Scommentare se si vuole abilitare l'audit per l'attività utente.
	//E' necessario attivare anche authentication.JwtAuthMiddleware
	r.Use(audit.AuditMiddleware)

	gestioneEndpointPubblici(r)
	gestioneEndpointRisorsa(r)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "DELETE", "PATCH", "PUT", "OPTIONS"})

	handler := handlers.CORS(headersOk, originsOk, methodsOk)(r)

	zipkinStruct := utils.GetZipkinInstance()
	defer zipkinStruct.HTTPReporter.Close()
	serverMiddleware := zipkinhttp.NewServerMiddleware(zipkinStruct.Tracer, zipkinhttp.TagResponseSize(true))
	log.Println("Avvio server in ascolto su porta" + config.ServerPort + "!")
	err := http.ListenAndServe(config.ServerPort, serverMiddleware(handler))
	if err != nil {
		log.Fatal(err)
	}
}

func gestioneEndpointPubblici(r *mux.Router) {
	var config = config.GetInstance()
	var swaggerAPI = server.SwaggerAPIFactory()

	log.Println("----------------Endpoint Pubblici -----------------------------")
	r.HandleFunc(config.Context.PublicContextPath+"/swagger", swaggerAPI.SwaggerConfig).Methods("GET")
	log.Println("Mapped GET\t\t", config.Context.PublicContextPath+"/swagger")
	log.Println("----------------------------------------------------------------")

}

func gestioneEndpointRisorsa(r *mux.Router) {
	var templateAPI = server.TemplateAPIFactory()

	//Verifica ruolo per autorizzazione su metodo
	//middleware.CheckRole(<metodo>, lista ruoli)
	//r.HandleFunc(config.AppContextPath+"/", authentication.CheckRole(templateAPI.CreaTemplate, "OPEARR", "RESPSE")).Methods("POST")

	log.Println("----------------Endpoint Risorsa Contesto App-----------------------------")
	r.HandleFunc("hello", templateAPI.RecuperaTemplateByID).Methods("GET")
	log.Println("Mapped GET\t\t", "hello")
	log.Println("---------------------------------------------")
}
