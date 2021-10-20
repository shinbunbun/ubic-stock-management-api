package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/guregu/dynamo"
)

type UbicFoodHandler struct {
	table *dynamo.Table
}

type UbicFoodWidget struct {
	ID       string `dynamo:"ID,hash"`
	DataType string `dynamo:"DataType" index:"Data-DataType-index,range"`
	Data     string `dynamo:"Data" index:"Data-DataType-index,hash"`
	DataKind string `dynamo:"DataKind" index:"DataKind-index,hash"`
	IntData  int    `dynamo:"IntData"`
}

func (h *UbicFoodHandler) AddItem(w UbicFoodWidget) (string, error) {
	// ID欄をユニークな値に変えてデータを追加します
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	id := uuidObj.String()
	w.ID = id

	table := h.table

	err = table.Put(w).Run()

	if err != nil {
		return "", err
	}
	return id, nil
}

func (h *UbicFoodHandler) AddMultipleItems(widgets []UbicFoodWidget) (string, error) {
	// ID欄を同じユニークな値に変えてデータを追加します
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	id := uuidObj.String()

	table := h.table
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
		return "", errors.New("Failed to write All item")
	}

	return id, nil
}

func (h *UbicFoodHandler) ReplaceItems(id string, widgets []UbicFoodWidget) error {
	// IDに対応するデータをwidgetsで置き換えます

	table := h.table
	var items []interface{}
	var deleteKeys []dynamo.Keyed

	{
		deletes, err := h.GetByID(id)
		if err != nil {
			return err
		}
		for _, w := range deletes {
			deleteKeys = append(deleteKeys, dynamo.Keys{w.ID, w.DataType})
		}
	}

	for i := range widgets {
		widgets[i].ID = id
		items = append(items, widgets[i])
	}

	n, err := table.Batch("ID", "DataType").
		Write().
		Delete(deleteKeys...).
		Run()

	if err != nil {
		return err
	}

	if n != len(deleteKeys) {
		return errors.New("Failed to write")
	}

	n, err = table.Batch("ID", "DataType").
		Write().
		Put(items...).
		Run()

	if err != nil {
		return err
	}

	if n != len(widgets) {
		return errors.New("Failed to write")
	}

	return nil
}

func (h *UbicFoodHandler) GetByID(id string) ([]UbicFoodWidget, error) {
	// IDの値から一致するデータのリストを返します
	table := h.table

	var res []UbicFoodWidget
	err := table.Get("ID", id).All(&res)
	if err != nil {
		return []UbicFoodWidget{}, err
	}
	return res, nil
}

func (h *UbicFoodHandler) DeleteByID(id string) error {
	// 同じIDを持つデータを消去します
	table := h.table

	ws, err := h.GetByID(id)
	if err != nil {
		return err
	}
	var keys []dynamo.Keyed
	for _, w := range ws {
		keys = append(keys, dynamo.Keys{w.ID, w.DataType})
	}
	n, err := table.Batch("ID", "DataType").
		Write().
		Delete(keys...).Run()
	if n != len(ws) {
		return errors.New("Failed all items")
	}
	if err != nil {
		return err
	}
	return nil
}

func (h *UbicFoodHandler) GetByDataAndDataType(data string, datatype string) (UbicFoodWidget, error) {
	// DataとDataTypeの値から探して単一の要素を返します
	table := h.table

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

func (h *UbicFoodHandler) GetByData(data string) ([]UbicFoodWidget, error) {
	// Dataの値が一致するデータのリストを返します
	table := h.table

	var res []UbicFoodWidget
	err := table.Get("Data", data).
		Index("Data-DataType-index").
		All(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *UbicFoodHandler) GetByDataKind(dataKind string) ([]UbicFoodWidget, error) {
	// DataKindが一致するデータを返します
	table := h.table

	var res []UbicFoodWidget
	err := table.Get("DataKind", dataKind).
		Index("DataKind-index").
		All(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
