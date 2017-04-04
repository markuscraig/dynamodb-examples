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
	startYr := 1950
	endYr := 1959

	// create the api params
	params := &dynamodb.ScanInput{
		TableName:            aws.String("Movies"),
		ProjectionExpression: aws.String("#yr, title, info.rating"),
		FilterExpression:     aws.String("#yr between :start_yr and :end_yr"),
		ExpressionAttributeNames: map[string]*string{
			"#yr": aws.String("year"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":start_yr": {
				N: aws.String(strconv.Itoa(startYr)),
			},
			":end_yr": {
				N: aws.String(strconv.Itoa(endYr)),
			},
		},
	}

	// scan and filter for the items
	err := db.ScanPages(params, scanHandler)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		return
	}
}

func scanHandler(page *dynamodb.ScanOutput, lastPage bool) bool {
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
