package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func startInstance(instanceID string) (string, error) {
	svc := ec2.New(session.New(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
	}
	result, err := svc.StartInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return "info", nil
	}
	fmt.Println(result)
	return "info", nil
}

func stopInstance(instanceID string) (string, error) {
	session, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
	svc := ec2.New(session)
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
	}

	result, err := svc.StopInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return "error", err
	}

	fmt.Println(result)
	return "success", nil
}

func main() {
	// lambda.Start(greet)
	// instanceID := "i-099f282532153947e"
	instanceID := "i-099f282532153947e"
	startInstance(instanceID)
	// stopInstance(instanceID)
}
