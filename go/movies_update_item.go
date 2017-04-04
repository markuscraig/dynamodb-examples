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

	// update values
	rating := 4.5
	plot := "Everything happens ONCE AGAIN all at once."
	actors := []string{
		"Larry", "Moe", "Curly",
	}

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
		UpdateExpression: aws.String("set info.rating=:r, info.plot=:p, info.actors=:a"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {N: aws.String(strconv.FormatFloat(rating, 'f', -1, 64))},
			":p": {S: aws.String(plot)},
			":a": {SS: aws.StringSlice(actors)},
		},
		ReturnValues: aws.String(dynamodb.ReturnValueUpdatedNew),
	}

	// update the item
	resp, err := db.UpdateItem(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return
	}

	// unmarshal the dynamodb attribute values into a custom struct
	var movie types.Movie
	err = dynamodbattribute.UnmarshalMap(resp.Attributes, &movie)

	// print the response data
	fmt.Printf("Updated Movie = %+v\n", movie)
}
