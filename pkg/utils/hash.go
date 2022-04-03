package utils

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

//CalculateHash calcola l'hash
func CalculateHash(v interface{}) (string, error) {
	sha512 := sha512.New()
	content, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	hash := sha512.Sum(content)
	return base64.URLEncoding.EncodeToString(hash), nil
}

//CalculateHash calcola l'hash di file

func CalculateHashFromByteArray(b []byte) string {

	hasher := sha512.New()

	hasher.Write(b)

	return hex.EncodeToString(hasher.Sum(nil))

}
