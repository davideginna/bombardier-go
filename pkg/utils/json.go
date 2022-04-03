package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
)

//Costanti
const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json; charset=UTF-8"
)

//RespondWithError : funzione per la restituzione di un errore JSON
func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{
		"status":  strconv.Itoa(code),
		"message": msg,
	})
}

// RespondWithJSON : funzione per la restituzione di una risposta JSON
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.MarshalIndent(payload, "", " ")
	//response, _ := json.Marshal(payload)
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(code)
	w.Write(response)
}

//ParsePost : funzione per il parsing di una richiesta POST
func ParsePost(request *http.Request, body interface{}) error {
	decoder := json.NewDecoder(request.Body)
	return decoder.Decode(&body)
}
