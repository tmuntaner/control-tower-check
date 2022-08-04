package aws_api

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/organizations"
)

type OrganizationsApi struct {
	client *organizations.Organizations
}

func NewOrganizationsApi(sess *session.Session) *OrganizationsApi {
	config := aws.Config{Region: aws.String(endpoints.UsEast1RegionID)}
	client := organizations.New(sess, &config)

	organizationsApi := OrganizationsApi{
		client: client,
	}

	return &organizationsApi
}

func (p *OrganizationsApi) ListAccountsForOu(organizationalUnitId string) ([]*organizations.Account, error) {
	input := organizations.ListAccountsForParentInput{ParentId: aws.String(organizationalUnitId)}
	var accounts []*organizations.Account

	err := p.client.ListAccountsForParentPages(&input, func(page *organizations.ListAccountsForParentOutput, lastPage bool) bool {
		accounts = append(accounts, page.Accounts...)
		return true
	})
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
