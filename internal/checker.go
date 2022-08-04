package internal

import (
	"control-tower-check/internal/aws_api"
	"fmt"
	"os"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/organizations"
)

type Checker struct {
	sess *session.Session
}

func NewChecker() *Checker {
	sess := session.Must(session.NewSession(&aws.Config{}))

	mapper := Checker{
		sess: sess,
	}

	return &mapper
}

func (a *Checker) Check(organizationalUnit string) error {
	api := aws_api.NewOrganizationsApi(a.sess)

	accounts, err := api.ListAccountsForOu(organizationalUnit)
	if err != nil {
		return err
	}

	for _, account := range accounts {
		regions := []string{endpoints.EuCentral1RegionID, endpoints.UsEast1RegionID, endpoints.UsEast2RegionID, endpoints.UsWest2RegionID, endpoints.ApSoutheast1RegionID, endpoints.ApSoutheast2RegionID, endpoints.EuWest1RegionID, endpoints.EuWest2RegionID, endpoints.EuNorth1RegionID}

		for _, region := range regions {
			sess, err := aws_api.AssumeRole(account, a.sess, region)
			if err != nil {
				return err
			}

			cfm := aws_api.NewCloudformationApi(sess, region)
			stacks, err := cfm.ListStacks()
			if err != nil {
				return err
			}

			stackName, configStatus, ok := cloudwatchStack(stacks)
			printCloudwatchStatus(account, region, stackName, configStatus, ok)

			stackName, configStatus, ok = configStack(stacks)
			printCloudwatchStatus(account, region, stackName, configStatus, ok)
		}
	}

	return nil
}

func printCloudwatchStatus(account *organizations.Account, region string, stackName string, status string, ok bool) {
	if !ok {
		_, _ = fmt.Fprintf(os.Stderr, "Config stack not found for account %s in region %s\n", aws.StringValue(account.Name), region)
	} else {
		fmt.Printf("%s\t%s\t%s\t%s\n", aws.StringValue(account.Name), region, stackName, status)
	}
}

func cloudwatchStack(stacks []*cloudformation.StackSummary) (string, string, bool) {
	r := regexp.MustCompile("StackSet-AWSControlTowerBP-BASELINE-CLOUDWATCH")

	for _, stack := range stacks {
		stackName := aws.StringValue(stack.StackName)
		if r.MatchString(stackName) {
			return stackName, aws.StringValue(stack.StackStatus), true
		}
	}

	return "", "", false
}

func configStack(stacks []*cloudformation.StackSummary) (string, string, bool) {
	r := regexp.MustCompile("StackSet-AWSControlTowerBP-BASELINE-CONFIG")

	for _, stack := range stacks {
		stackName := aws.StringValue(stack.StackName)
		if r.MatchString(stackName) {
			return stackName, aws.StringValue(stack.StackStatus), true
		}
	}

	return "", "", false
}
