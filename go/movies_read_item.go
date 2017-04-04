package main

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/fitquick/dynamodb-examples/go/types"
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

	// query parameters
	year := 2015
	title := "The Big New Movie"

	// create the api params
	params := &dynamodb.GetItemInput{
		TableName: aws.String("Movies"),
		Key: map[string]*dynamodb.AttributeValue{
			"year": {
				N: aws.String(strconv.Itoa(year)),
			},
			"title": {
				S: aws.String(title),
			},
		},
	}

	// read the item
	resp, err := db.GetItem(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return
	}

	// unmarshal the dynamodb attribute values into a custom struct
	var movie types.Movie
	err = dynamodbattribute.UnmarshalMap(resp.Item, &movie)

	// print the response data
	fmt.Printf("Unmarshaled Movie = %+v\n", movie)
}
