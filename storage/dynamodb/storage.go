package dynamodb_storage

import (
	"log"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v1"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	storage "github.com/nickschuch/karma/storage"
)

var (
	cliAWSRegion = kingpin.Flag("aws-region", "The region DynamoDB runs.").Default("us-west-2").OverrideDefaultFromEnvar("KARMA_AWS_REGION").String()
	cliAWSTable  = kingpin.Flag("karma", "The table to store the data.").Default("karma").OverrideDefaultFromEnvar("KARMA_AWS_TABLE").String()
	keyName      = "Username"
	valueName    = "Karma"
)

type DynamoDBStorage struct{}

func init() {
	storage.Register("dynamodb", &DynamoDBStorage{})
}

func (p *DynamoDBStorage) Get(n string) int {
	c := client()

	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			keyName: {
				S: aws.String(n),
			},
		},
		TableName: aws.String(*cliAWSTable),
	}
	resp, err := c.GetItem(params)
	check(err)

	// If we cannot find the user we just want to return a 0.
	if len(resp.Item) <= 0 {
		return 0
	}

	// Look for the value inside the list of items.
	for k, s := range resp.Item {
		if k == valueName {
			strconv.Atoi(*s.N)
		}
	}

	// Looks like we didn't find the field we were looking for.
	log.Println("Could not find field for", n)
	return 0
}

func (p *DynamoDBStorage) Set(n string, v int) {
	c := client()

	amount := strconv.Itoa(v)
	params := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{ // Required
			keyName: {
				S: aws.String(n),
			},
			valueName: {
				N: aws.String(amount),
			},
		},
		TableName: aws.String(*cliAWSTable),
	}
	_, err := c.PutItem(params)
	check(err)
}

func (p *DynamoDBStorage) Increase(n string, v int) {
	value := p.Get(n)
	value = value + v
	p.Set(n, value)
}

func (p *DynamoDBStorage) Decrease(n string, v int) {
	value := p.Get(n)
	value = value - v
	p.Set(n, value)
}

func client() *dynamodb.DynamoDB {
	// Connect to the AWS API.
	c := dynamodb.New(&aws.Config{Region: *cliAWSRegion})

	// Ensure the table exists.
	params := &dynamodb.DescribeTableInput{
		TableName: aws.String(*cliAWSTable),
	}
	_, err := c.DescribeTable(params)
	if awsErr, ok := err.(awserr.Error); ok {
		// If we cannot find the table we create one.
		if awsErr.Code() == "ResourceNotFoundException" {
			log.Println("Table does not exist so we are creating one.")
			params := &dynamodb.CreateTableInput{
				AttributeDefinitions: []*dynamodb.AttributeDefinition{
					{
						AttributeName: aws.String(keyName),
						AttributeType: aws.String("S"),
					},
				},
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String(keyName),
						KeyType:       aws.String("HASH"),
					},
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Long(1),
					WriteCapacityUnits: aws.Long(1),
				},
				TableName: aws.String(*cliAWSTable),
			}
			_, err := c.CreateTable(params)
			check(err)
		}
	}

	return c
}

func check(err error) {
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS Error with Code, Message, and original error (if any)
			log.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				log.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, The SDK should alwsy return an
			// error which satisfies the awserr.Error interface.
			log.Println(err.Error())
		}
	}
}
