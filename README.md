# kong-swagger
A Swagger definition for various elements of Kong

Example POST for key auth:

```go
var e struct{}
res, _, err := client.DefaultApi.ConsumersConsumerIdKeyAuthPost(consumerName, map[string]interface{}{ "empty": e })
```
