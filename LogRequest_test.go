package httpparser

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

type StructTest1 struct {
	Name  string  `log:"false"`
	Body  *string `log:"false"`
	Test2 StructTest2
	Test3 []StructTest3
	Test4 StructTest2 `log:"false"`
}

type StructTest2 struct {
	Hello string
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
		Name: "ken",
		Body: aws.String("hello"),
		Test2: StructTest2{
			Hello: "There",
		},
		Test4: StructTest2{
			Hello: "yew",
		},
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
	if s1.Name != "ken" {
		t.Fail()
	}
	if *s1.Body != "hello" {
		t.Fail()
	}
}
