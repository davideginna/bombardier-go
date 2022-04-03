package utils

import "encoding/json"

//Serialize ...
func Serialize(v interface{}) string {
	jsonContent, _ := json.MarshalIndent(v, " ", " ")
	return string(jsonContent)
}
