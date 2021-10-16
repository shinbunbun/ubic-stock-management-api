package database

import (
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/guregu/dynamo"
)

type DynamoDBHandler struct {
	conn *dynamo.DB
}

type DataBaseError string

func (d DataBaseError) Error() string {
	return string(d)
}

const (
	NotFoundError DataBaseError = DataBaseError("Failed to Get Item Because there are no correspond item")
)

func NewDynamoDBHandler() *DynamoDBHandler {
	// DynamoDBHandlerを生成して返します
	sess := GetDynamoDBSession()
	conn := dynamo.New(sess, &aws.Config{
		Region:   aws.String(config.AWSRegion()),
		Endpoint: aws.String(config.DynamoDBEndpoint()),
	})
	return &DynamoDBHandler{
		conn: conn,
	}
}

func (h *DynamoDBHandler) NewUbicFoodHandler() *UbicFoodHandler {
	table := h.conn.Table(config.DataTable())
	return &UbicFoodHandler{
		table: &table,
	}
}
