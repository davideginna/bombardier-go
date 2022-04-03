package utils

import (
	"log"
	"net/url"
)

//ExtractParam ...
func ExtractParam(param url.Values, name string) string {
	if param[name] != nil {
		log.Printf(">>> %s : %s \n", name, param[name])
		return param[name][0]
	}
	return ""
}
