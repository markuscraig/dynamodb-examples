package types

/*

NOTE: To unmarshal Go []string to an AWS 'L' type (list of AttributeValues),
you must use the []interface{} type.  The []string type will get unmarshaled
into an AWS 'SS' type (string-set of AttributeValues). The 'L' type is needed
for DynamoDB JSON expressions, like...

   "remove info.actors[0]"

Here is the AWS SDK source-code link that describes the Unmarshal types...

   https://github.com/aws/aws-sdk-go/blob/master/service/dynamodb/dynamodbattribute/decode.go#L36-L76

*/
type MovieInfo struct {
	Directors       []interface{} `json:"directors" dynamodbav:"directors"`
	ReleaseDate     string        `json:"release_date" dynamodbav:"release_date"`
	Rating          float64       `json:"rating" dynamodbav:"rating"`
	Genres          []interface{} `json:"genres" dynamodbav:"genres"`
	ImageURL        string        `json:"image_url" dynamodbav:"image_url"`
	Plot            string        `json:"plot" dynamodbav:"plot"`
	Rank            int           `json:"rank" dynamodbav:"rank"`
	RunningTimeSecs int           `json:"running_time_secs" dynamodbav:"running_time_secs"`
	Actors          []interface{} `json:"actors" dynamodbav:"actors"`
}

type Movie struct {
	Year  int       `json:"year" dynamodbav:"year"`
	Title string    `json:"title" dynamodbav:"title"`
	Info  MovieInfo `json:"info" dynamodbav:"info"`
}
