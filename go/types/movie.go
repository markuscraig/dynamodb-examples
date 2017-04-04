package types

type MovieInfo struct {
	Directors       []string `json:"directors" dynamodbav:"directors"`
	ReleaseDate     string   `json:"release_date" dynamodbav:"release_date"`
	Rating          float64  `json:"rating" dynamodbav:"rating"`
	Genres          []string `json:"genres" dynamodbav:"genres"`
	ImageURL        string   `json:"image_url" dynamodbav:"image_url"`
	Plot            string   `json:"plot" dynamodbav:"plot"`
	Rank            int      `json:"rank" dynamodbav:"rank"`
	RunningTimeSecs int      `json:"running_time_secs" dynamodbav:"running_time_secs"`
	Actors          []string `json:"actors" dynamodbav:"actors"`
}

type Movie struct {
	Year  int       `json:"year" dynamodbav:"year"`
	Title string    `json:"title" dynamodbav:"title"`
	Info  MovieInfo `json:"info" dynamodbav:"info"`
}
