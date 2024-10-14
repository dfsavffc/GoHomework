package app

import (
	"encoding/json"
	"github.com/dfsavffc/GoHomework/task2/server/internal/pkg/decoding"
	"github.com/dfsavffc/GoHomework/task2/server/pkg/models"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func getVersion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		log.Printf("getVersion method not allowed")
		return
	}
	_, err := w.Write([]byte(versionApi))
	if err != nil {
		log.Printf("getVersion write error: %v", err)
		return
	}
}
func getHardOp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		log.Printf("getHardOp method not allowed")
		return
	}
	time.Sleep(time.Duration(rand.Intn(11)+10) * time.Second)
	if statusId := rand.Intn(2); statusId != 1 {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		log.Printf("getHardOp write error: %v", statusId)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}

}
func postDecode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		log.Printf("postDecode method not allowed")
		return
	}
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("postDecode read body error: %v", err)
		return
	}
	request := models.Request{}
	err = json.Unmarshal(requestBody, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("postDecode unmarshal error: %v", err)
		return
	}
	outputString, err := decoding.DecodeBase64(request.Input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("postDecode decode string error: %v", err)
	}
	response := models.Response{Output: outputString}
	responseBody, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("postDecode marshal error: %v", err)
	}
	_, err = w.Write(responseBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("postDecode write error: %v", err)
		return
	}
}
