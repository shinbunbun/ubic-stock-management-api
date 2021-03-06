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
	return h.AddItemWithID(id, w)
}

func (h *UbicFoodHandler) AddItemWithID(id string, w UbicFoodWidget) (string, error) {
	// データを追加します
	table := h.table
	w.ID = id

	err := table.Put(w).Run()

	if err != nil {
		return "", err
	}
	return w.ID, nil
}

func (h *UbicFoodHandler) AddMultipleItems(widgets []UbicFoodWidget) (string, error) {
	// ID欄を同じユニークな値に変えてデータを追加します
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	id := uuidObj.String()
	return h.AddMultipleItemsWithID(id, widgets)
}

func (h *UbicFoodHandler) AddMultipleItemsWithID(id string, widgets []UbicFoodWidget) (string, error) {
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

func (h *UbicFoodHandler) UpdateData(id, dataType, data string) error {
	// id,dataTypeに対応する行のDataをdataに変えます
	table := h.table
	return table.Update("ID", id).
		Range("DataType", dataType).
		Set("Data", data).
		Run()
}

func (h *UbicFoodHandler) UpdateIntDataTo(id, dataType string, data int) error {
	// id,dataTypeに対応する行のintDataをdataに変えます
	table := h.table
	return table.Update("ID", id).
		Range("DataType", dataType).
		Set("IntData", data).
		Run()
}

func (h *UbicFoodHandler) UpdateIntDataBy(id, dataType string, add int) error {
	// id,dataTypeに対応する行のintDataをaddだけ加算します
	table := h.table
	return table.Update("ID", id).
		Range("DataType", dataType).
		Add("IntData", add).
		Run()
}

func (h *UbicFoodHandler) UpdateIntDataByWithoutMinus(id string, dataType string, add int) error {
	table := h.table
	return table.Update("ID", id).
		Range("DataType", dataType).
		Add("IntData", add).
		If("$ >= ?", "IntData", -add).
		Run()
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

func (h *UbicFoodHandler) GetByMultipleIDs(ids []string) ([]UbicFoodWidget, error) {
	// id <- ids についてIDがidであるデータ全てを返します。
	table := h.table
	var keys []dynamo.Keyed
	for _, id := range ids {
		keys = append(keys, dynamo.Keys{id})
	}
	var res []UbicFoodWidget
	err := table.Batch("ID").
		Get(keys...).
		All(&res)
	if err != nil {
		return []UbicFoodWidget{}, err
	}
	return res, nil
}

func (h *UbicFoodHandler) GetByDataLikeWithDataKindAndType(like string, dataKind string, dataType string) ([]UbicFoodWidget, error) {
	table := h.table

	var res []UbicFoodWidget
	err := table.Get("DataKind", dataKind).
		Filter("$ = ?", "DataType", dataType).
		Filter("$ = ?", "DataKind", dataKind).
		Filter("contains($, ?)", "Data", like).
		All(&res)
	if err != nil {
		return []UbicFoodWidget{}, err
	}
	return res, nil
}

func (h *UbicFoodHandler) GetByIDAndDataType(id, datatype string) (UbicFoodWidget, error) {
	// IDとDataTypeが一致するデータを返します
	table := h.table

	var res UbicFoodWidget
	err := table.Get("ID", id).
		Range("DataType", dynamo.Equal, datatype).
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
