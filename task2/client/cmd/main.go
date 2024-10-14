package main

import (
	"fmt"
	"github.com/dfsavffc/GoHomework/task2/client/internal/app"
	"github.com/dfsavffc/GoHomework/task2/client/pkg/models"
	"log"
)

func main() {
	client := app.NewClient("http://localhost:8000")

	version, err := client.GetVersion()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(version)

	encodedData := "0JrQsNC60LDRjy3RgtC+INC30LDQutC+0LTQuNGA0L7QstCw0L3QvdCw0Y8g0YHRgtGA0L7QutCw"
	request := models.Request{Input: encodedData}
	decodedData, err := client.PostDecode(request)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(decodedData)

	successful, status, err := client.GetHardOp()
	if err != nil {
		log.Fatal(err)
	}
	if successful {
		fmt.Println(successful, status)
	}

}
