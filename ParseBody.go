package httpparser

import (
	"net/http/httptest"
	"strings"

	"github.com/albrow/forms"
	"github.com/aws/aws-lambda-go/events"
)

//ParseBody parses APIGatewayProxyRequest object to forms.Data
func ParseBody(request events.APIGatewayProxyRequest) (*forms.Data, error, string) {
	httpRequest := httptest.NewRequest(request.HTTPMethod, request.Path, strings.NewReader(request.Body))
	c := "abc"
	if content, ok := request.Headers["Content-Type"]; ok {
		httpRequest.Header.Set("Content-Type", content)
		c = content
	} else if content, ok := request.Headers["content-type"]; ok {
		httpRequest.Header.Set("Content-Type", content)
		c = content
	}

	err := httpRequest.ParseForm()
	if err != nil {
		return nil, err, ""
	}
	userData, err := forms.Parse(httpRequest)
	return userData, err, c
}
