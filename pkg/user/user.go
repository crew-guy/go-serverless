package user

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorFailedToFetchRecord = "failed to fetch record"
	ErrorInvalidUserData = "invalid user data"
	ErrorInvalidEmail = "invalid email"
	ErrorCouldNotMarshalItem = "could not marshal item"
	ErrorCouldNotDeleteItem = "could not delete item"
	ErrorCouldNotDynamoPutItem = "could not dynamo put item"
	ErrorUserAlreadyExists = "user already exists"
	ErrorUserDoesNotExist = "user does not exist"
);

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName" `
}

func FetchUser(email, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*User, error) {
	input := &dynamodb.GetItemInput{
		Key:map[string]*dynamodb.AttributeValue{
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
func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI)(
	*User,
	error
) {
	var u User;
	if err := json.Unmarshal([]byte(req.body), &u);
	err != nil{
		return nil, errors.New(ErrorInvalidUserData)
	}
	if !validators.IsEmailValid(u.Email){
		return nil, errors.New(ErrorInvalidEmail)
	}

	// check if user exists
	currentUser := FetchUser(u.Email, tableName, dynaClient)
	if currentUser !=nil && len(currentUser.Email) != 0{
		return nil, errors.New(ErrorUserAlreadyExists)
	}
	av, err := dynamodbattribute.marshalMap(u)

	if err!=nil{
		return errors.New(ErrorCouldNotMarshalItem)
	}

	input:= &dynamodb.PutItemInput{
		Item:av,
		TableName : aws.String(tableName)
	}

	dynaClient.PutItem(input)
}
func UpdateUser() {}
func DeleteUser() {}
