package aws_api

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type CloudFormationApi struct {
	client *cloudformation.CloudFormation
}

func NewCloudformationApi(sess *session.Session, region string) *CloudFormationApi {
	config := aws.Config{Region: aws.String(region)}
	client := cloudformation.New(sess, &config)

	organizationsApi := CloudFormationApi{
		client: client,
	}

	return &organizationsApi
}

func (p *CloudFormationApi) ListStacks() ([]*cloudformation.StackSummary, error) {
	var stacks []*cloudformation.StackSummary

	input := cloudformation.ListStacksInput{}
	err := p.client.ListStacksPages(&input, func(page *cloudformation.ListStacksOutput, lastPage bool) bool {
		stacks = append(stacks, page.StackSummaries...)
		return true
	})
	if err != nil {
		return nil, err
	}

	return stacks, nil
}
