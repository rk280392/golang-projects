package uploadToS3

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func CreateS3Bucket(ctx context.Context, s3Client *s3.Client, bucketName, region string) error {
	bucketsList, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return fmt.Errorf("listBucker error: %s", err)
	}

	found := false
	for _, bucket := range bucketsList.Buckets {
		if *bucket.Name == bucketName {
			found = true
		}
	}

	if !found {
		_, err = s3Client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
			CreateBucketConfiguration: &types.CreateBucketConfiguration{
				LocationConstraint: types.BucketLocationConstraint(region),
			},
		})
		if err != nil {
			return fmt.Errorf("error CreateS3Bucket : %s", err)
		}
	} else {
		return fmt.Errorf("bucket already exists: %s", bucketName)
	}

	return nil
}

func UploadToS3(ctx context.Context, s3Client *s3.Client, bucketName, fileName string) error {
	uploader := manager.NewUploader(s3Client)
	_, err := uploader.Upload(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("test.txt"),
		Body:   strings.NewReader("This is testing"),
	})
	if err != nil {
		return fmt.Errorf("uploader error: %s", err)
	}

	return nil
}
