package main

import (
	"fmt"
	//"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Info struct {
	Plot    string `dynamodbav:"plot"`
	Ratings int    `dynamodbav:"ratings"`
}

type Item struct {
	Year  int    `dynamodbav:"year"`
	Title string `dynamodbav:"title"`
	Info  Info   `dynamodbav:"info"`
}

func main() {
	// create a session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://127.0.0.1:8000"),
		//EndPoint: aws.String("https://dynamodb.us-east-1.amazonaws.com"),
	}))

	// create a dynamodb instance
	db := dynamodb.New(sess)

	// item data
	item := Item{
		Year:  2015,
		Title: "The Big New Movie",
		Info: Info{
			Plot:    "Nothing happens at all.",
			Ratings: 0,
		},
	}

	// marshal the info struct into an aws attribute value
	itemAVMap, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		panic("Cannot marshal item into AttributeValue map")
	}

	// create the api params
	params := &dynamodb.PutItemInput{
		TableName: aws.String("Movies"),
		Item:      itemAVMap,
	}

	// create the table
	resp, err := db.PutItem(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return
	}

	// print the response data
	fmt.Println("Success")
	fmt.Println(resp)
}
