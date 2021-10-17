package database

import "github.com/aws/aws-sdk-go/aws/session"

var dynamoDBSession *session.Session = nil

func GetDynamoDBSession() *session.Session {
	// セッションを生成して返します。もし，過去に生成しているならそれを使いまわします
	if dynamoDBSession != nil {
		return dynamoDBSession
	}
	dynamoDBSession = session.Must(session.NewSession())
	return dynamoDBSession
}
