package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/fitquick/dynamodb-examples/go/types"
)

func main() {
	// read the json data file
	f, err := ioutil.ReadFile("../data/moviedata.json")
	if err != nil {
		panic("Could not read movie json data file")
	}

	// parse the json movie data
	var movies []types.Movie
	if err := json.Unmarshal(f, &movies); err != nil {
		panic("Could not parse json movie data")
	}

	// create an aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://127.0.0.1:8000"),
		//EndPoint: aws.String("https://dynamodb.us-east-1.amazonaws.com"),
	}))

	// create a dynamodb instance
	db := dynamodb.New(sess)

	// iterate through each movie
	for _, m := range movies {
		// marshal the movie struct into an aws attribute value map
		movieAVMap, err := dynamodbattribute.MarshalMap(m)
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
			fmt.Printf("Unable to add movie: %v\n", err.Error())
		} else {
			// print the response data
			fmt.Printf("Put item successful: '%s' (resp = '%+v')\n", m.Title, resp)
		}
	}
}
