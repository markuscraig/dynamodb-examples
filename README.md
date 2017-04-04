## Examples of accessing DynamoDB using the AWS SDK

Examples have been written in the following languagues...

1. Node.js (javascript)
   * http://docs.aws.amazon.com/amazondynamodb/latest/gettingstartedguide/GettingStarted.NodeJs.html
2. Golang (coming soon)
   * http://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/

## Local vs Cloud DynamoDB

These examples write to a locally running DynamoDB instance (downloadable from AWS)...

* http://docs.aws.amazon.com/amazondynamodb/latest/gettingstartedguide/intro-dynamodb-local.html

Change the 'endpoint' in the AWS.config in each script to write the AWS cloud DynamoDB...

```javascript
AWS.config.update({
   endpoint: "https://dynamodb.us-east-1.amazonaws.com"}
);
```

## Running Examples

### Go

```bash
$ go run go/movies_create_item.go
$ go run go/movies_create_table.go
$ go run go/movies_delete_item.go
$ go run go/movies_delete_table.go
$ go run go/movies_describe_table.go
$ go run go/movies_load_data.go
$ go run go/movies_query_year.go
$ go run go/movies_query_year_title.go
$ go run go/movies_read_item.go
$ go run go/movies_scan.go
$ go run go/movies_update_atomic_counter.go
$ go run go/movies_update_conditionally.go
$ go run go/movies_update_item.go
```

### Node.js

```bash
$ node node/movies_create_item.js
$ node node/movies_create_table.js
$ node node/movies_delete_item.js
$ node node/movies_delete_table.js
$ node node/movies_load_data.js
$ node node/movies_query_year.js
$ node node/movies_query_year_title.js
$ node node/movies_query_pages.js
$ node node/movies_read_item.js
$ node node/movies_scan.js
$ node node/movies_update_atomic_counter.js
$ node node/movies_update_conditionally.js
$ node node/movies_update_item.js
```

Made with :green_heart: in Campbell, CA
