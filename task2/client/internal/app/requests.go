package app

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/dfsavffc/GoHomework/server/pkg/models"
	"io"
	"log"
	"net/http"
	"time"
)

func (c *Client) GetVersion(url string) (string, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return "", err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("failed to execute request: %v\n", err)
		return "", err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("failed to reading response body: %v", err)
		return "", err
	}
	err = response.Body.Close()
	if err != nil {
		log.Printf("failed to close response body")
		return "", err
	}
	return string(body), nil
}

func (c *Client) GetHardOp(url string) (bool, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Printf("failed to create request: %v\n", err)
		return false, 0, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("failed to execute request: %v\n", err)
		return false, 0, err
	}
	err = response.Body.Close()
	if err != nil {
		log.Printf("failed to close response body: %v\n", err)
		return false, 0, err
	}
	return true, response.StatusCode, nil
}

func (c *Client) PostDecode(url string, data []byte) (string, error) {
	encodedString := base64.StdEncoding.EncodeToString(data)

	requestBody := models.Request{InputString: encodedString}
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		log.Printf("failed to encode request body: %v\n", err)
		return "", err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		log.Printf("failed to create request: %v\n", err)
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Printf("failed to execute request: %v\n", err)
		return "", err
	}

	responseBodyJSON, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("failed to read response body: %v\n", err)
		return "", err
	}

	responseBody := models.Response{}
	err = json.Unmarshal(responseBodyJSON, &responseBody)
	if err != nil {
		log.Printf("failed to decode response body: %v\n", err)
		return "", err
	}

	return responseBody.OutputString, nil
}
