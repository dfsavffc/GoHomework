package app

import (
	"net/http"
	"strconv"
)

const versionApi = "v1.0.0"

func NewServer(host int) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/version", getVersion)
	mux.HandleFunc("/hard-op", getHardOp)
	mux.HandleFunc("/decode", postDecode)

	srv := &http.Server{
		Handler: mux,
		Addr:    "localhost:" + strconv.Itoa(host),
	}
	return srv
}
