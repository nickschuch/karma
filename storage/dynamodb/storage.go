package dynamodb_storage

import (
	"gopkg.in/alecthomas/kingpin.v1"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	storage "github.com/nickschuch/karma/storage"
)

var (
	cliAWSRegion = kingpin.Flag("aws-region", "The region DynamoDB runs.").Default("us-west-2").OverrideDefaultFromEnvar("KARMA_AWS_REGION").String()
	cliAWSKey    = kingpin.Flag("aws-key", "The key which is used to authenticate.").Default("").OverrideDefaultFromEnvar("KARMA_AWS_KEY").String()
	cliAWSSecret = kingpin.Flag("aws-secret", "The secret which is used to authenticate.").Default("").OverrideDefaultFromEnvar("KARMA_AWS_SECRET").String()
	cliAWSTable  = kingpin.Flag("karma", "The table to store the data.").Default("karma").OverrideDefaultFromEnvar("KARMA_AWS_TABLE").String()
)

type DynamoDBStorage struct{}

func init() {
	storage.Register("dynamodb", &DynamoDBStorage{})
}

func (p *DynamoDBStorage) Get(n string) int {
	c := client()

	return 0
}

func (p *DynamoDBStorage) Set(n string, v int) {
	// Nothing for now.
}

func (p *DynamoDBStorage) Increase(n string, v int) {
	// Nothing for now.
}

func (p *DynamoDBStorage) Decrease(n string, v int) {
	// Nothing for now.
}

func client() string {
	// Connect to the AWS API.
	c := dynamodb.New(nil)

	// Ensure the table exists.
	params := &dynamodb.ListTablesInput{
		ExclusiveStartTableName: aws.String(*cliAWSTable),
		Limit: aws.Long(1),
	}
	resp, _ := c.ListTables(params)

	return c
}
