### Examples of accessing DynamoDB using the AWS SDK

Examples have been written in the following languagues...

1. Node.js (javascript)
   * http://docs.aws.amazon.com/amazondynamodb/latest/gettingstartedguide/GettingStarted.NodeJs.html
2. Golang (coming soon)
   * http://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/

### Local vs Cloud DynamoDB

These examples write to a locally running DynamoDB instance (downloadable from AWS)...

* http://docs.aws.amazon.com/amazondynamodb/latest/gettingstartedguide/intro-dynamodb-local.html

Change the 'endpoint' in the AWS.config in each script to write the AWS cloud DynamoDB...

```javascript
AWS.config.update({
   endpoint: "https://dynamodb.us-east-1.amazonaws.com"}
);
```
