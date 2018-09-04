package httpparser

import (
	"encoding/json"
	"fmt"
	"os"
)

//LogRequest logs request to stderr
func LogRequest(request interface{}) {
	b, err := json.Marshal(request)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Fprintf(os.Stderr, "REQUEST: %s\n", string(b))
}
