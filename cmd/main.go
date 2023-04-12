package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	handler "github.com/lucas-simao/cloud-function-example"
)

func main() {
	fmt.Println("Starting Cloud Function Debug...")
	printPayloadExample()

	ctx := context.Background()
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/", handler.Handler); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext: %v\n", err)
	}
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}

func printPayloadExample() {
	body := []byte(`{
		"title": "This is a Title",
		"body": "and this is a body",
		"value": 100.25
	}`)

	fmt.Println("Here you can find a payload example in bytes:")

	var payload string

	for i, v := range body {
		if i != len(body)-1 {
			payload += fmt.Sprintf("%v, ", v)
		} else {
			payload += fmt.Sprintf("%v", v)
		}
	}

	fmt.Printf("[%s]\n", payload)
}
