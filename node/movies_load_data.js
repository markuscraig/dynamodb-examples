var AWS = require('aws-sdk');
var fs = require('fs');

DATA_FILE = '../data/moviedata.json';

AWS.config.update({
  region: "us-east-1",
  endpoint: "http://127.0.0.1:8000"
  //endpoint: "https://dynamodb.us-east-1.amazonaws.com"
});

var docClient = new AWS.DynamoDB.DocumentClient();

console.log("Importing movies into DynamoDB...");

var allMovies = JSON.parse(fs.readFileSync(DATA_FILE, 'utf8'));

allMovies.forEach(function(movie) {
  var params = {
    TableName: "Movies",
    Item: {
      year:  movie.year,
      title: movie.title,
      info:  movie.info
    }
  };

  docClient.put(params, function(err, data) {
    if (err) {
      console.error("Unable to add movie", movie.title,
        ". Error JSON:", JSON.stringify(err, null, 2));
    } else {
      console.log("Put item successful:", movie.title);
    }
  })
})