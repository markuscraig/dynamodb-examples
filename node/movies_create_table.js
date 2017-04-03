var AWS = require('aws-sdk');

AWS.config.update({
  region: "us-east-1",
  endpoint: "http://127.0.0.1:8000"
  //endpoint: "https://dynamodb.us-east-1.amazonaws.com"
});

var dynamodb = new AWS.DynamoDB();

var params = {
  TableName: "Movies",
  KeySchema: [
    { AttributeName: "year",  KeyType: "HASH" },  // partition key
    { AttributeName: "title", KeyType: "RANGE" }, // sort key
  ],
  AttributeDefinitions: [
    { AttributeName: "year",  AttributeType: "N" }, // number
    { AttributeName: "title", AttributeType: "S" }, // string
  ],
  ProvisionedThroughput: {
    ReadCapacityUnits: 10,
    WriteCapacityUnits: 100
  }
};

dynamodb.createTable(params, function(err, data) {
  if (err) {
    console.error("Could not create table: err =", JSON.stringify(err, null, 2));
  } else {
    console.log("Created table. Table description =", JSON.stringify(data, null, 2));
  }
});
