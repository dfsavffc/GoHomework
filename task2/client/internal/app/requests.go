package app

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/dfsavffc/GoHomework/task2/client/pkg/models"
	"io"
	"log"
	"net/http"
	"time"
)

const clientTimeLimit = 15 * time.Second

func (c *Client) GetVersion() (string, error) {
	request, err := http.NewRequest(http.MethodGet, c.url+"/version", nil)
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

func (c *Client) GetHardOp() (bool, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), clientTimeLimit)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url+"/hard-op", nil)
	if err != nil {
		log.Printf("failed to create request: %v\n", err)
		return false, http.StatusBadRequest, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return false, http.StatusRequestTimeout, nil
		}
		log.Printf("failed to execute request: %v\n", err)
		return false, http.StatusBadRequest, err
	}

	err = response.Body.Close()
	if err != nil {
		log.Printf("failed to close response body: %v\n", err)
		return false, http.StatusBadRequest, err
	}
	return true, response.StatusCode, nil
}

func (c *Client) PostDecode(requestBody models.Request) (string, error) {
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		log.Printf("failed to encode request body: %v\n", err)
		return "", err
	}

	request, err := http.NewRequest(http.MethodPost, c.url+"/decode", bytes.NewBuffer(requestBodyJSON))
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

	return responseBody.Output, nil
}
