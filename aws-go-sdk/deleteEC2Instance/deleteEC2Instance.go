package deleteEC2Instance

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func DeleteKeyPairs(ctx context.Context, keyName string, cfg aws.Config) error {
	ec2Client := ec2.NewFromConfig(cfg)
	_, err := ec2Client.DeleteKeyPair(ctx, &ec2.DeleteKeyPairInput{
		KeyName: &keyName,
	})
	if err != nil {
		return fmt.Errorf("DeleteKeyPair error: %s", err)
	}
	return nil
}

func DeleteEC2Instance(ctx context.Context, keyName string, cfg aws.Config) error {
	ec2Client := *ec2.NewFromConfig(cfg)
	DescInstanceOut, err := ec2Client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("key-name"),
				Values: []string{keyName},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("DescribeInstance error : %s", err)
	}

	instanceIds := []string{}
	// Print the instance IDs
	for _, reservation := range DescInstanceOut.Reservations {
		for _, instance := range reservation.Instances {
			instanceIds = append(instanceIds, *instance.InstanceId)
		}
	}
	if len(instanceIds) == 0 {
		return fmt.Errorf("instances found error: %d", len(instanceIds))
	}

	ec2Client.TerminateInstances(ctx, &ec2.TerminateInstancesInput{
		InstanceIds: instanceIds,
	})

	err = os.Remove(keyName)
	if err != nil {
		return fmt.Errorf("file %s doesn't exists: ", keyName)
	}
	return nil

}
