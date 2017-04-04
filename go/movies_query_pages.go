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
	year := 1985

	// create the api params
	params := &dynamodb.QueryInput{
		TableName:              aws.String("Movies"),
		KeyConditionExpression: aws.String("#yr = :yyyy"), // 'year' is reserved keyword
		ExpressionAttributeNames: map[string]*string{
			"#yr": aws.String("year"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":yyyy": {
				N: aws.String(strconv.Itoa(year)),
			},
		},
	}

	// scan and filter for the items
	err := db.QueryPages(params, queryHandler)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return
	}
}

func queryHandler(page *dynamodb.QueryOutput, lastPage bool) bool {
	// dump the response data
	//fmt.Println(page)

	// Unmarshal the slice of dynamodb attribute values
	// into a slice of custom structs
	var movies []types.Movie
	err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &movies)
	if err != nil {
		// print the error and continue receiving pages
		fmt.Printf("\nCould not unmarshal AWS data: err = %v\n", err)
		return true
	}

	// print the response data
	for _, m := range movies {
		fmt.Printf("Movie: '%s' (%d)\n", m.Title, m.Year)
	}

	// if not done receiving all of the pages
	if lastPage == false {
		fmt.Printf("\n*** NOT DONE RECEIVING PAGES ***\n\n")
	} else {
		fmt.Printf("\n*** RECEIVED LAST PAGE ***\n\n")
	}

	// continue receiving pages (can be used to limit the number of pages)
	return true
}
