package main

import (
	"fmt"

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

	// movie data
	movie := types.Movie{
		Year:  2015,
		Title: "The Big New Movie",
		Info: types.MovieInfo{
			Plot:   "Nothing happens at all.",
			Rating: 1.1,
		},
	}

	// marshal the movie struct into an aws attribute value
	movieAVMap, err := dynamodbattribute.MarshalMap(movie)
	if err != nil {
		panic("Cannot marshal movie into AttributeValue map")
	}

	// create the api params
	params := &dynamodb.PutItemInput{
		TableName: aws.String("Movies"),
		Item:      movieAVMap,
	}

	// put the item
	resp, err := db.PutItem(params)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return
	}

	// print the response data
	fmt.Println("Success")
	fmt.Println(resp)
}
