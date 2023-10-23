package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	createec2instance "github.com/rk280392/aws-go-sdk/createEC2Instance"
	"github.com/rk280392/aws-go-sdk/deleteEC2Instance"
	"github.com/rk280392/aws-go-sdk/downloadFromS3"
	"github.com/rk280392/aws-go-sdk/uploadToS3"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <service> <action>")
		return
	}

	ctx := context.Background()
	keyName := "aws-go-key"
	region := "ap-south-1"
	bucketName := "rajesh-aws-sdk-s3-test-shhwgdf12"
	fileName := "test.txt"
	service := os.Args[1]
	action := os.Args[2]

	if service == "ec2" {
		ec2Client, err := initEC2Client(ctx, region)
		if err != nil {
			fmt.Printf("initS3Client error: %s", err)
			os.Exit(1)
		}
		switch action {
		case "create":
			err := createec2instance.CreateKeyPairs(ctx, keyName, ec2Client)
			if err != nil {
				fmt.Printf("Error CreateInstance: %s", err)
				os.Exit(1)
			}

			instanceId, err := createec2instance.CreateEC2Instance(ctx, keyName, ec2Client)
			if err != nil {
				fmt.Printf("createEC2Instance error: %s", err)
			}
			fmt.Printf("instanceId: %s\n", instanceId)
		case "delete":
			err := deleteEC2Instance.DeleteKeyPairs(ctx, keyName, ec2Client)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}
			err = deleteEC2Instance.DeleteEC2Instance(ctx, keyName, ec2Client)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}
		default:
			fmt.Println("Invalid action. Use 'create' or 'delete'.")
		}
	} else if service == "s3" {
		s3Client, err := initS3Client(ctx, region)
		if err != nil {
			fmt.Printf("initS3Client error: %s\n", err)
			os.Exit(1)
		}
		switch action {
		case "upload":
			err := uploadToS3.CreateS3Bucket(ctx, s3Client, bucketName, region)
			if err != nil {
				fmt.Printf("Error Createbucket: %s\n", err)
				os.Exit(1)
			}
			err = uploadToS3.UploadToS3(ctx, s3Client, bucketName, fileName)
			if err != nil {
				fmt.Printf("Error Uploadfile: %s\n", err)
				os.Exit(1)
			}
			fmt.Println("Upload Success!!")
		case "download":
			contents, err := downloadFromS3.DownloadFromS3(ctx, s3Client, bucketName, fileName)
			if err != nil {
				fmt.Printf("Download from S3 error: %s\n", err)
			}
			fmt.Printf("Downloaded file with contents: %s\n", contents)
		}
	}

}

func initEC2Client(ctx context.Context, region string) (*ec2.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile("personal"))
	if err != nil {
		return nil, fmt.Errorf("loadconfig error: %s", err)
	}
	return ec2.NewFromConfig(cfg), nil
}

func initS3Client(ctx context.Context, region string) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile("personal"))
	if err != nil {
		return nil, fmt.Errorf("loadconfig error: %s", err)
	}
	return s3.NewFromConfig(cfg), nil
}
