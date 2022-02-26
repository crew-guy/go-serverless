package user

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"errors"
	"encoding/json"
	"github.com/aws/aws-sdk-go/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-lambda-go/events"
)

var(
	ErrorFailedToFetchRecord = "failed to fetch record",
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName" `
}

func FetchUser(email, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {
	input := &dynamodb.GetItemInput{
		Key:map[string]*AttributeValue{
			"email":{
				S:aws.String(email)
			}
		},
		TableName := aws.String(tableName)
	}
	result, err := dynaClient.GetItem(input)
	if err!=nil{
		return nil,errors.New(ErrorFailedToFetchRecord)
	}

	item := new(User)
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err!=nil{
		return nil, errors.New(ErrorFailedToFetchRecord)
	}
	return item, nil
}

func FetchUsers(tableName string, dynaClient dynamodbiface.DynamoDBAPI)(*[]User, error) {
	input := &dynaClient.ScanInput{
		TableName:aws.String(tableName)
	}
	result, err = dynaClient.Scan(input)

	if err!=nil{
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	data := new([]User)
	err = dynamodbattribute.UnmarshalMap(result.Items,data )
	return data, nil
}
func CreateUser() {}
func UpdateUser() {}
func DeleteUser() {}
