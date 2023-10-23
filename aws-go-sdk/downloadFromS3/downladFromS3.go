package downloadFromS3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func DownloadFromS3(ctx context.Context, s3Client *s3.Client, bucketName, fileName string) ([]byte, error) {
	downloader := manager.NewDownloader(s3Client)
	buffer := manager.NewWriteAtBuffer([]byte{})
	numBytes, err := downloader.Download(ctx, buffer, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    &fileName,
	})

	if numBytes != int64(len(buffer.Bytes())) {
		return nil, fmt.Errorf("number of bytes doesn't match. Received %d, Actual %d", len(buffer.Bytes()), numBytes)
	}
	if err != nil {
		return nil, fmt.Errorf("error download file from s3: %s", err)
	}

	return buffer.Bytes(), nil

}
