package main

import (
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func startInstance(instanceID string) (bool, error) {
	svc := ec2.New(session.New(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
	}
	_, err := svc.StartInstances(input)
	if err != nil {
		// TODO: error handling
		// aerr, ok := err.(awserr.Error); ok {
		return false, err
	}
	return true, err
}

func stopInstance(instanceID string) (bool, error) {
	svc := ec2.New(session.New(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	}))
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(instanceID),
		},
	}
	_, err := svc.StopInstances(input)
	if err != nil {
		// TODO: error handling
		// aerr, ok := err.(awserr.Error); ok {
		return false, err
	}
	return true, err
}

func isActive(t time.Time) bool {
	if isForceRunning() {
		return true
	}
	if isHoliday(t) {
		switch {
		case t.Hour() < 7:
			return false
		case t.Hour() < 23:
			return true
		default:
			return false
		}
	}
	switch {
	case t.Hour() < 15:
		return false
	case t.Hour() < 22:
		return true
	default:
		return false
	}
}

func isHoliday(t time.Time) bool {
	// TODO: consider national holiday
	return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
}

func isForceRunning() bool {
	forceRunning, err := strconv.ParseBool(os.Getenv("FORCE_RUNNING"))
	if err != nil {
		println(err)
	}
	return forceRunning
}

//
// Handling...
//
func Handling() (string, error) {
	instanceID := os.Getenv("INSTANCE_ID")
	t := time.Now()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	current := t.In(jst)
	if isActive(current) {
		startInstance(instanceID)
		return "succeded: start instance.", nil
	}
	stopInstance(instanceID)
	return "succeeded: stop instance.", nil
}

func main() {
	lambda.Start(Handling)
}
