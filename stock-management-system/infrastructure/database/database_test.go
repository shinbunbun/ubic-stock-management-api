package database

//func TestAddItem(t *testing.T) {
//	// AddItem関数のテスト
//	h := NewDynamoDBHandler()
//	w := Widget{
//		ID:       "",
//		DataType: "type",
//		Data:     "data",
//		DataKind: "kind",
//		IntData:  0,
//	}
//	t.Run("Add Data", func(t *testing.T) {
//		id, err := h.AddItem(w)
//		CheckError(t, err, nil)
//		CheckDatabase(t, h, id, w)
//	})
//	t.Run("Add Second Time", func(t *testing.T) {
//		id, err := h.AddItem(w)
//		CheckError(t, err, nil)
//		CheckDatabase(t, h, id, w)
//	})
//}
//
//func TestGetByDataAndDataType(t *testing.T) {
//	// GetByDataAndDataType関数のテスト
//	h := NewDynamoDBHandler()
//
//	dataType := "type"
//	data := "data"
//
//	w := Widget{
//		ID:       "",
//		DataType: dataType,
//		Data:     data,
//		DataKind: "kind",
//		IntData:  0,
//	}
//	_, err := h.AddItem(w)
//	CheckError(t, err, nil)
//
//	t.Run("Succesful Get", func(t *testing.T) {
//		got, err := h.GetByDataAndDataType(data, dataType)
//		CheckError(t, err, nil)
//		CheckWidget(t, got, w)
//	})
//	t.Run("Failed to Get", func(t *testing.T) {
//		_, err := h.GetByDataAndDataType(data, dataType+"1")
//		CheckError(t, err, NotFoundError)
//	})
//}
//
//func CheckDatabase(t testing.TB, h *DynamoDBHandler, id string, want Widget) {
//	// IDの値がidであるデータがwantであるかを判定。複数データがある時はエラー
//	t.Helper()
//	w, err := h.GetByID(id)
//	CheckError(t, err, nil)
//	if len(w) != 1 {
//		t.Fatalf("want to get a single widgets by id")
//	}
//	CheckWidget(t, w[0], want)
//}
//
//func CheckWidget(t testing.TB, got, want Widget) {
//	// ID欄を除いて同じかを判定
//	t.Helper()
//	got.ID = "0"
//	want.ID = "0"
//	if got != want {
//		t.Errorf("got %q widget, want %q", got, want)
//	}
//}
//
//func CheckError(t testing.TB, got, want error) {
//	t.Helper()
//	if want != nil && got != nil {
//		if want.Error() != got.Error() {
//			t.Errorf("got %q error, want %q", got, want)
//		}
//	} else {
//		if want != nil {
//			t.Fatalf("want to got a error %q", want)
//		}
//		if got != nil {
//			t.Errorf("got a error %q", got)
//		}
//	}
//}
