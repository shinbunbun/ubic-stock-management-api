package database

import "github.com/aws/aws-sdk-go/aws/session"

var dynamoDBSession *session.Session

func GetDynamoDBSession() *session.Session {
	if dynamoDBSession != nil {
		return dynamoDBSession
	}
	dynamoDBSession = session.Must(session.NewSession())
	return dynamoDBSession
}
