package main

//aws iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json

//aws lambda create-function =-function-name go-lambda-function --zip-file fileb://function.zip --handler go-lambda-function --runtime go1.x --role arn:aws:iam::-:role/lambda-ex

//aws lambda invoke --function-name go-lambda-function-3 --cli-binary-format raw-in-base64-out --payload {"What is your name?": "Jim", "How old are you?": 33}' output.txt

// be aware of the architecture of the lambda function

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"What is your name?"`
	Age  int    `json:"How old are you?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cwd)
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
