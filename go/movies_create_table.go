package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	// create an aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://127.0.0.1:8000"),
		//EndPoint: aws.String("https://dynamodb.us-east-1.amazonaws.com"),
	}))

	// create a dynamodb instance
	db := dynamodb.New(sess)

	// create the api params
	params := &dynamodb.CreateTableInput{
		TableName: aws.String("Movies"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("year"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("title"), KeyType: aws.String("RANGE")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("year"), AttributeType: aws.String("N")},
			{AttributeName: aws.String("title"), AttributeType: aws.String("S")},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(100),
		},
	}

	// create the table
	resp, err := db.CreateTable(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// print the response data
	fmt.Println(resp)
}
