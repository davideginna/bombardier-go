package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// VersioneSoftware recupera la versione implementata dei servizi REST
func VersioneSoftware() interface{} {
	//log.Println("VersionSoftware:: servizio VersionSoftware")
	var data interface{}
	dir, _ := os.Getwd()
	raw, err := ioutil.ReadFile(dir + "/docs/version.json")
	if err != nil {
		return nil
	}
	json.Unmarshal(raw, &data)
	return data
}
