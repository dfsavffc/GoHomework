package decoding

import (
	"encoding/base64"
	"log"
)

func DecodeBase64(encodedString string) (string, error) {
	decodedString, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		log.Printf("error decoding base64 string: %v\n", err)
		return "", err
	}
	log.Printf("decode successful: %s\n", string(decodedString))
	return string(decodedString), nil
}
