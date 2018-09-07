package httpparser

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

type StructTest1 struct {
	Name  string  `log:"false"`
	Body  *string `log:"false"`
	Test2 *StructTest2
	Test3 []StructTest3
	Scope []*string
	Hello map[string]string
	Wow   map[string]*string
	Test4 *StructTest2 `log:"false"`
	Hehe  []byte
	Asd   []string `log:"false"`
}

type StructTest2 struct {
	Hello string
	Gg    map[string]string `log:"false"`
}

type StructTest3 struct {
	Bye string `log:"false"`
}

func (s StructTest1) Log() {
	tempT3 := make([]StructTest3, 0)
	for _, val := range s.Test3 {
		tempT3 = append(tempT3, val)
	}
	s.Test3 = tempT3
	LogRequest(&s)
}

func TestLogRequest(t *testing.T) {
	s1 := StructTest1{
		Name:  "ken",
		Body:  aws.String("hello"),
		Scope: []*string{aws.String("hello"), aws.String("tere")},
		Test2: &StructTest2{
			Hello: "There",
			Gg:    map[string]string{"hello": "there"},
		},
		Test4: &StructTest2{
			Hello: "yew",
			Gg:    map[string]string{"hello": "there"},
		},
		Hello: map[string]string{"hello": "there"},
		Wow:   map[string]*string{"hello": aws.String("there"), "hmmm": nil},
		Hehe:  []byte{byte('A'), byte('B')},
		Asd:   []string{"asd", "bef"},
	}
	s3T1 := StructTest3{
		Bye: "a",
	}
	s3T2 := StructTest3{
		Bye: "n",
	}
	s3 := []StructTest3{s3T1, s3T2}
	s1.Test3 = s3
	s1.Log()
	fmt.Println(s1)
	if s1.Name != "ken" {
		t.Fail()
	}
	if *s1.Body != "hello" {
		t.Fail()
	}
}
