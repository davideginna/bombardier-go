package middleware

import (
	"log"
	"net/http"
)

//LoggingMiddleware logga le informazioni della request http
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(" ------------ NEW REQUEST --------------- ")
		log.Println(" RequestUri   : ", r.RequestURI)
		log.Println(" RequestHeader: ", r.Header)
		log.Println(" ------------ ........... --------------- ")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
