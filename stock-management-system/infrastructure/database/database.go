package database

import (
	"errors"

	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
)

type DynamoDBHandler struct {
	conn *dynamo.DB
}

type Widget struct {
	ID       string `dynamo:"ID,hash"`
	DataType string `dynamo:"DataType" index:"Data-DataType-index,range"`
	Data     string `dynamo:"Data" index:"Data-DataType-index,hash"`
	DataKind string `dynamo:"DataKind" index:"DataKind-index,hash"`
	IntData  int    `dynamo:"IntData"`
}

func NewDynamoDBHandler() *DynamoDBHandler {
	sess := GetDynamoDBSession()
	conn := dynamo.New(sess, &aws.Config{Region: aws.String(config.AWSRegion())})
	return &DynamoDBHandler{
		conn: conn,
	}
}

func (h *DynamoDBHandler) AddItem(w Widget) (string, error) {
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	id := uuidObj.String()
	w.ID = id

	table := h.conn.Table(config.DataTable())
	err = table.Put(w).Run()
	if err != nil {
		return "", err
	}
	return id, nil
}

func (h *DynamoDBHandler) AddItems(widgets []Widget) (string, error) {
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	id := uuidObj.String()

	table := h.conn.Table(config.DataTable())
	var items []interface{}

	for i := range widgets {
		widgets[i].ID = id
		items = append(items, widgets[i])
	}

	n, err := table.Batch().
		Write().
		Put(items...).
		Run()

	if err != nil {
		return "", err
	}

	if n != len(widgets) {
		return "", errors.New("Failed to write")
	}

	return id, nil
}

func (h *DynamoDBHandler) GetByID(id string) ([]Widget, error) {
	table := h.conn.Table(config.DataTable())

	var res []Widget
	err := table.Get("ID", id).All(&res)
	if err != nil {
		return []Widget{}, err
	}
	return res, nil
}

func (h *DynamoDBHandler) DeleteByID(id string) error {
	table := h.conn.Table(config.DataTable())

	return table.Delete("ID", id).Run()
}

func (h *DynamoDBHandler) GetByDataAndDataType(data string, datatype string) (Widget, error) {
	table := h.conn.Table(config.DataTable())

	var res Widget
	err := table.Get("Data", data).
		Range("DataType", dynamo.Equal, datatype).
		One(&res)
	if err != nil {
		return Widget{}, err
	}
	return res, nil
}

func (h *DynamoDBHandler) GetByData(data string) ([]Widget, error) {
	table := h.conn.Table(config.DataTable())

	var res []Widget
	err := table.Get("Data", data).
		All(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *DynamoDBHandler) GetByDataKind(dataKind string) ([]Widget, error) {
	table := h.conn.Table(config.DataTable())

	var res []Widget
	err := table.Get("DataKind", dataKind).
		All(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
