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
	return NewDynamoDBHandlerWithEndpoint("")
}

func NewDynamoDBHandlerWithEndpoint(endpoint string) *DynamoDBHandler {
	sess := GetDynamoDBSession()

	cfg := &aws.Config{}
	if region := config.AWSRegion(); region != "" {
		cfg.Region = aws.String(region)
	}
	if endpoint != "" {
		cfg.Endpoint = aws.String(endpoint)
	}

	conn := dynamo.New(sess, cfg)
	return &DynamoDBHandler{
		conn: conn,
	}
}

func (h *DynamoDBHandler) NewUbicFoodHandler() *UbicFoodHandler {
	return h.NewUbicFoodHandlerWithTableName(config.DataTable())
}

func (h *DynamoDBHandler) NewUbicFoodHandlerWithTableName(tableName string) *UbicFoodHandler {
	table := h.conn.Table(tableName)
	return &UbicFoodHandler{
		table: &table,
	}
}
