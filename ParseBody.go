package httpparser

import (
	"fmt"
	"net/http/httptest"
	"strings"

	"github.com/albrow/forms"
	"github.com/aws/aws-lambda-go/events"
)

//ParseBody parses APIGatewayProxyRequest object to forms.Data
func ParseBody(request events.APIGatewayProxyRequest) (*forms.Data, error) {
	httpRequest := httptest.NewRequest(request.HTTPMethod, request.Path, strings.NewReader(request.Body))
	if content, ok := request.Headers["Content-Type"]; ok {
		fmt.Println("Content-Type")
		fmt.Println(content)
		httpRequest.Header.Add("Content-Type", content)
	} else if content, ok := request.Headers["content-type"]; ok {
		fmt.Println("content-type")
		fmt.Println(content)
		httpRequest.Header.Add("Content-Type", content)
	} else {
		fmt.Println("no content type")
		fmt.Println(request.Headers)
	}

	err := httpRequest.ParseForm()
	if err != nil {
		return nil, err
	}
	userData, err := forms.Parse(httpRequest)
	return userData, err
}
