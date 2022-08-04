package aws_api

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/organizations"
)

func AssumeRole(account *organizations.Account, sess *session.Session, region string) (*session.Session, error) {
	roleArn := fmt.Sprintf("arn:aws:iam::%s:role/%s", aws.StringValue(account.Id), "AWSControlTowerExecution")
	credentials := stscreds.NewCredentials(sess, roleArn)

	config := aws.Config{
		Region:      aws.String(region),
		Credentials: credentials,
	}

	sess, err := session.NewSession(&config)
	if err != nil {
		return nil, err
	}

	return sess, nil
}
