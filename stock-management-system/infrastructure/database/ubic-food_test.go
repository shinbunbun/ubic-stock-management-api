package database

import (
	"testing"
)

func TestAddItem(t *testing.T) {
	// AddItem関数のテスト
	h := newDummyHandler()
	w := UbicFoodWidget{
		ID:       "",
		DataType: "type",
		Data:     "data",
		DataKind: "kind",
		IntData:  0,
	}
	t.Run("Add Data", func(t *testing.T) {
		id, err := h.AddItem(w)
		checkError(t, err, nil)
		checkDatabase(t, h, id, w)
	})
	t.Run("Add Second Time", func(t *testing.T) {
		id, err := h.AddItem(w)
		checkError(t, err, nil)
		checkDatabase(t, h, id, w)
	})
}

func TestGetByDataAndDataType(t *testing.T) {
	// GetByDataAndDataType関数のテスト
	h := newDummyHandler()

	dataType := "type"
	data := "data"

	w := UbicFoodWidget{
		ID:       "",
		DataType: dataType,
		Data:     data,
		DataKind: "kind",
		IntData:  0,
	}
	_, err := h.AddItem(w)
	checkError(t, err, nil)

	t.Run("Succesful Get", func(t *testing.T) {
		got, err := h.GetByDataAndDataType(data, dataType)
		checkError(t, err, nil)
		checkWidget(t, got, w)
	})
	t.Run("Failed to Get", func(t *testing.T) {
		_, err := h.GetByDataAndDataType(data, dataType+"1")
		checkError(t, err, NotFoundError)
	})
}

func TestGetByData(t *testing.T) {
	// GetByData関数のテスト
	h := newDummyHandler()
	data1 := "data1"
	data2 := "data2"

	w1 := UbicFoodWidget{
		ID:       "",
		DataType: "type1",
		Data:     data1,
		DataKind: "kind",
		IntData:  0,
	}

	w2 := UbicFoodWidget{
		ID:       "",
		DataType: "type2",
		Data:     data2,
		DataKind: "kind",
		IntData:  0,
	}

	_, err := h.AddItem(w1)
	checkError(t, err, nil)
	_, err = h.AddItem(w1)
	checkError(t, err, nil)
	_, err = h.AddItem(w2)
	checkError(t, err, nil)

	t.Run("Successful get 2 item", func(t *testing.T) {
		ws, err := h.GetByData(data1)
		checkError(t, err, nil)
		if len(ws) != 2 {
			t.Fatalf("Invalid Data length")
		}
	})
	t.Run("Successful get 1 item", func(t *testing.T) {
		ws, err := h.GetByData(data2)
		checkError(t, err, nil)
		if len(ws) != 1 {
			t.Fatalf("Invalid Data length")
		}
	})
}

func TestGetByKind(t *testing.T) {
	// GetByData関数のテスト
	h := newDummyHandler()
	kind1 := "kind1"
	kind2 := "kind2"

	w1 := UbicFoodWidget{
		ID:       "",
		DataType: "type1",
		Data:     "data1",
		DataKind: kind1,
		IntData:  0,
	}

	w2 := UbicFoodWidget{
		ID:       "",
		DataType: "type2",
		Data:     "data2",
		DataKind: kind2,
		IntData:  0,
	}

	_, err := h.AddItem(w1)
	checkError(t, err, nil)
	_, err = h.AddItem(w1)
	checkError(t, err, nil)
	_, err = h.AddItem(w2)
	checkError(t, err, nil)

	t.Run("Successful get 2 item", func(t *testing.T) {
		ws, err := h.GetByDataKind(kind1)
		checkError(t, err, nil)
		if len(ws) != 2 {
			t.Fatalf("Invalid Data length")
		}
	})
	t.Run("Successful get 1 item", func(t *testing.T) {
		ws, err := h.GetByDataKind(kind2)
		checkError(t, err, nil)
		if len(ws) != 1 {
			t.Fatalf("Invalid Data length")
		}
	})
}

func checkDatabase(t testing.TB, h *UbicFoodHandler, id string, want UbicFoodWidget) {
	// IDの値がidであるデータがwantであるかを判定。複数データがある時はエラー
	t.Helper()
	w, err := h.GetByID(id)
	checkError(t, err, nil)
	if len(w) != 1 {
		t.Fatalf("want to get a single widgets by id")
	}
	checkWidget(t, w[0], want)
}

func checkWidget(t testing.TB, got, want UbicFoodWidget) {
	// ID欄を除いて同じかを判定
	t.Helper()
	got.ID = "0"
	want.ID = "0"
	if got != want {
		t.Errorf("got %q widget, want %q", got, want)
	}
}

func checkError(t testing.TB, got, want error) {
	t.Helper()
	if want != nil && got != nil {
		if want.Error() != got.Error() {
			t.Errorf("got %q error, want %q", got, want)
		}
	} else {
		if want != nil {
			t.Fatalf("want to got a error %q", want)
		}
		if got != nil {
			t.Errorf("got a error %q", got)
		}
	}
}

var (
	dummyHandler *UbicFoodHandler = nil
)

func newDummyHandler() *UbicFoodHandler {
	if dummyHandler == nil {
		h := NewDynamoDBHandler()
		table := h.conn.Table("UBIC-FOOD-test")
		dummyHandler = &UbicFoodHandler{
			table: &table,
		}
	}

	dummyHandler.cleanAllItems()

	return dummyHandler
}

func (h *UbicFoodHandler) cleanAllItems() {

	table := dummyHandler.table

	var widgets []UbicFoodWidget
	err := table.Scan().All(&widgets)
	if err != nil {
		panic(err)
	}

	di := make(map[string]bool)
	for _, widget := range widgets {
		di[widget.ID] = true
	}

	for id := range di {
		err := h.DeleteByID(id)
		if err != nil {
			panic(err)
		}
	}
}
