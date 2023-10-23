package createec2instance

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func CreateKeyPairs(ctx context.Context, region, keyName string, cfg aws.Config) error {

	ec2Client := ec2.NewFromConfig(cfg)
	existingKeyPair, err := ec2Client.DescribeKeyPairs(ctx, &ec2.DescribeKeyPairsInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("key-name"),
				Values: []string{keyName},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("describeKeyPair err: %s", err)
	}
	if existingKeyPair == nil || len(existingKeyPair.KeyPairs) == 0 {
		keyPairs, err := ec2Client.CreateKeyPair(ctx, &ec2.CreateKeyPairInput{
			KeyName: aws.String(keyName),
		})
		if err != nil {
			return fmt.Errorf("createKeyPair error: %s", err)
		}
		err = os.WriteFile(keyName+".pem", []byte(*keyPairs.KeyMaterial), 0600)
		if err != nil {
			return fmt.Errorf("keypair write file error: %s", err)
		}
	}
	return nil
}

func CreateEC2Instance(ctx context.Context, keyName string, cfg aws.Config) (string, error) {
	ec2Client := ec2.NewFromConfig(cfg)
	describeImages, err := ec2Client.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("name"),
				Values: []string{"ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-*"},
			},
			{
				Name:   aws.String("virtualization-type"),
				Values: []string{"hvm"},
			},
		},
		Owners: []string{"099720109477"},
	})
	if err != nil {
		return "", fmt.Errorf("describeImageError: %s", err)
	}
	if len(describeImages.Images) == 0 {
		return "", fmt.Errorf("describeImages has empty length (%d)", len(describeImages.Images))
	}
	ec2Instance, err := ec2Client.RunInstances(ctx, &ec2.RunInstancesInput{
		MaxCount:     aws.Int32(1),
		MinCount:     aws.Int32(1),
		ImageId:      aws.String(*describeImages.Images[0].ImageId),
		InstanceType: types.InstanceTypeT3Micro,
		KeyName:      &keyName,
	})
	if err != nil {
		return "", fmt.Errorf("runInstance err %s", err)
	}

	return *ec2Instance.Instances[0].InstanceId, nil
}
