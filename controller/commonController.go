package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type CommonController struct {
}

const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json; charset=UTF-8"
)

func CommonControllerFactory() CommonController {
	log.Println("CommonControllerFactory -- START --")
	var commonController = CommonController{}
	return commonController
}

// SwaggerConfig genera la configurazione swagger dei servizi REST
func (c CommonController) Hello(w http.ResponseWriter, r *http.Request) {
	log.Println("CommonController:: servizio hello")
	var data interface{}
	respondWithJSON(w, http.StatusOK, data)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.MarshalIndent(payload, "", " ")
	//response, _ := json.Marshal(payload)
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(code)
	w.Write(response)
}
