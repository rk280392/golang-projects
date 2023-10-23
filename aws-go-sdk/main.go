package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	createec2instance "github.com/rk280392/aws-go-sdk/createEC2Instance"
	"github.com/rk280392/aws-go-sdk/deleteEC2Instance"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <action>")
		return
	}
	var (
		keyName string
		region  string
	)

	ctx := context.Background()
	keyName = "aws-go-key"
	region = "ap-south-1"
	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile("personal"))

	if err != nil {
		fmt.Printf("loadconfig error: %s", err)
	}
	action := os.Args[1]

	switch action {
	case "create":
		err = createec2instance.CreateKeyPairs(ctx, region, keyName, cfg)
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}

		instanceId, err := createec2instance.CreateEC2Instance(ctx, keyName, cfg)
		if err != nil {
			fmt.Printf("createEC2Instance error: %s", err)
		}
		fmt.Printf("instanceId: %s\n", instanceId)
	case "delete":
		err := deleteEC2Instance.DeleteKeyPairs(ctx, keyName, cfg)
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
		err = deleteEC2Instance.DeleteEC2Instance(ctx, keyName, cfg)
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid action. Use 'create' or 'delete'.")
	}

}
