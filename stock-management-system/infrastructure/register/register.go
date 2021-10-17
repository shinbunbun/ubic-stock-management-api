package register

import "github.com/google/uuid"

func GenerateConfirmCode() (string, error) {
	// UUID生成
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	code := uuidObj.String()
	return code, nil
}
