package shared

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
)

type AWSUtil struct{}

var sess *session.Session

func init() {
	var err error
	// Create a new AWS session
	sess, err = session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}
}

// fargateProfiles returns AWS Fargate Profiles in the EKS Cluster
func (util AWSUtil) fargateProfiles() {
	// Create a new EKS client
	svc := eks.New(sess)

	// Call the DescribeFargateProfile function
	resp, err := svc.DescribeFargateProfile(&eks.DescribeFargateProfileInput{
		FargateProfileName: aws.String("my-fargate-profile"),
		ClusterName:        aws.String("my-cluster"),
	})
	if err != nil {
		fmt.Println("Error describing Fargate profile:", err)
		return
	}

	// Print the response
	fmt.Println(resp)

	panic("unimplemented")
}

// IsFargatePod evaluates if pod runs on Fargate
func (util AWSUtil) IsFargatePod(namespace string, labels map[string]string) bool {
	panic("unimplemented")
}
