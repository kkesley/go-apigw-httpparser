package httpparser

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestParseBody(test *testing.T) {
	request := events.APIGatewayProxyRequest{
		HTTPMethod: "POST",
		Resource:   "hello",
		Path:       "/hello",
		Headers:    map[string]string{"content-type": "application/json"},
		Body:       "",
	}
	data, err := ParseBody(request)
	if err != nil {
		test.Fail()
	}
	fmt.Println(data)
}
