package createKeyPairs

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go/aws"
)

func CreateKeyPairs(ctx context.Context, region, keyName string) error {

	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile("personal"))

	if err != nil {
		return fmt.Errorf("loadconfig error: %s", err)
	}

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
	} else {
		return fmt.Errorf("existingKeyError: %s", err)
	}
	return nil
}
