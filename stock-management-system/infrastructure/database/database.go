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

type Widget struct {
	ID       string `dynamo:"ID,hash"`
	DataType string `dynamo:"DataType" index:"Data-DataType-index,range"`
	Data     string `dynamo:"Data" index:"Data-DataType-index,hash"`
	DataKind string `dynamo:"DataKind" index:"DataKind-index,hash"`
	IntData  int    `dynamo:"IntData"`
}

func (h *DynamoDBHandler) AddItem(w Widget) (string, error) {
	// ID欄をユニークな値に変えてデータを追加します
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
	// ID欄を同じユニークな値に変えてデータを追加します
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
	// IDの値から一致するデータのリストを返します
	table := h.conn.Table(config.DataTable())

	var res []Widget
	err := table.Get("ID", id).All(&res)
	if err != nil {
		return []Widget{}, err
	}
	return res, nil
}

func (h *DynamoDBHandler) DeleteByID(id string) error {
	// 同じIDを持つデータを消去します
	table := h.conn.Table(config.DataTable())

	return table.Delete("ID", id).Run()
}

func (h *DynamoDBHandler) GetByDataAndDataType(data string, datatype string) (UbicFoodWidget, error) {
	// DataとDataTypeの値から探して単一の要素を返します
	table := h.conn.Table(config.DataTable())

	var res UbicFoodWidget
	err := table.Get("Data", data).
		Range("DataType", dynamo.Equal, datatype).
		Index("Data-DataType-index").
		One(&res)
	switch err {
	case dynamo.ErrNotFound:
		return UbicFoodWidget{}, NotFoundError
	case nil:
	default:
		return UbicFoodWidget{}, err
	}
	return res, nil
}

func (h *DynamoDBHandler) GetByData(data string) ([]Widget, error) {
	// Dataの値が一致するデータのリストを返します
	table := h.conn.Table(config.DataTable())

	var res []Widget
	err := table.Get("Data", data).
		Index("Data-DataType-index").
		All(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *DynamoDBHandler) GetByDataKind(dataKind string) ([]Widget, error) {
	// DataKindが一致するデータを返します
	table := h.conn.Table(config.DataTable())

	var res []Widget
	err := table.Get("DataKind", dataKind).
		Index("DataKind-index").
		All(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *DynamoDBHandler) NewUbicFoodHandler() *UbicFoodHandler {
	table := h.conn.Table(config.DataTable())
	return &UbicFoodHandler{
		table: &table,
	}
}