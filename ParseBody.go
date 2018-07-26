package httpparser

import (
	"errors"
	"net/http/httptest"
	"strings"

	"github.com/albrow/forms"
	"github.com/aws/aws-lambda-go/events"
)

//ParseBody parses APIGatewayProxyRequest object to forms.Data
func ParseBody(request events.APIGatewayProxyRequest) (*forms.Data, error) {
	if len(request.Path) <= 0 || len(request.HTTPMethod) <= 0 {
		return forms.CreateFromMap(make(map[string]string)), errors.New("invalid body request")
	}
	httpRequest := httptest.NewRequest(request.HTTPMethod, request.Path, strings.NewReader(request.Body))
	if content, ok := request.Headers["Content-Type"]; ok {
		httpRequest.Header.Set("Content-Type", content)
	} else if content, ok := request.Headers["content-type"]; ok {
		httpRequest.Header.Set("Content-Type", content)
	}

	err := httpRequest.ParseForm()
	if err != nil {
		return nil, err
	}
	userData, err := forms.Parse(httpRequest)
	return userData, err
}
