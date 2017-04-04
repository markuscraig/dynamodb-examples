package main

import (
	"fmt"
	"strconv"
	"strings"

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

	// condition parameters
	minActors := 3

	// create the api params
	params := &dynamodb.UpdateItemInput{
		TableName: aws.String("Movies"),
		Key: map[string]*dynamodb.AttributeValue{
			"year": {
				N: aws.String(strconv.Itoa(year)),
			},
			"title": {
				S: aws.String(title),
			},
		},

		// NOTE: To remove an item by index, the attribute needs to be
		// an AWS 'L' (list) type.  The attribute would need to be created
		// as a slice of AttributeValues using 'dynamodbattribute.MarshalList(actors)'.
		UpdateExpression:    aws.String("remove info.actors[0]"),
		ConditionExpression: aws.String("size(info.actors) >= :num"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":num": {
				N: aws.String(strconv.Itoa(minActors)),
			},
		},
		ReturnValues: aws.String(dynamodb.ReturnValueAllNew),
	}

	// update the item conditionally
	resp, err := db.UpdateItem(params)
	if err != nil {
		// if the conditional check failed
		if strings.Contains(err.Error(), "ConditionalCheckFailedException") {
			fmt.Printf("Conditional check failed; did not update\nerr = %v\n", err)
		} else {
			fmt.Printf("ERROR: %v\n", err)
		}
		return
	}

	// unmarshal the dynamodb attribute values into a custom struct
	var movie types.Movie
	err = dynamodbattribute.UnmarshalMap(resp.Attributes, &movie)

	// print the response data
	fmt.Printf("Updated Movie = %+v\n", movie)
}
