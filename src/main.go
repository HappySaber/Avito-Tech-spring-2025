package main

import (
	"context"
	"fmt"
	"os"

	openapiclient "PVZ/internal/generated-client"
)

func main() {
	// Создание конфигурации клиента
	cfg := openapiclient.NewConfiguration()

	// Создание клиента
	client := openapiclient.NewAPIClient(cfg)

	// Создание запроса
	dummyLoginPostRequest := *openapiclient.NewDummyLoginPostRequest("Role_example")
	resp, r, err := client.DefaultAPI.DummyLoginPost(context.Background()).DummyLoginPostRequest(dummyLoginPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DummyLoginPost`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	// response from `DummyLoginPost`: string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DummyLoginPost`: %v\n", resp)
}
