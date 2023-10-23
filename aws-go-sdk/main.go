package main

import (
	"context"
	"fmt"
	"os"

	createKeyPairs "github.com/rk280392/aws-go-sdk/createKeyPair"
)

func main() {

	var (
		keyName string
		region  string
	)
	keyName = "aws-go-key"
	region = "ap-south-1"
	err := createKeyPairs.CreateKeyPairs(context.Background(), region, keyName)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	//	instanceId, err := createEC2Instance.createEC2Instance

}
