package config

import (
	"os"
	"strings"
)

func AWSRegion() string {
	// Regionの値を返す関数
	return "ap-northeast-1"
}

func DataTable() string {
	// データを保持しているテーブルの名前を返します
	return "UBIC-FOOD"
}

func DataTableTest() string {
	return "UBIC-FOOD-test"
}

func DynamoDBEndpoint() string {
	// DynamoDBのエンドポイントを指す文字列を返します
	str := os.Getenv("DYNAMO_DB")
	if str != "" {
		return str
	}
	return "http://localhost:8000"
}

func SenderEmailAddress() string {
	return os.Getenv("MAIL_SENDER")
}

func PrivateKey() string {
	return strings.Replace(os.Getenv("SIGNINGKEY"), "\\n", "\n", -1)
}

func PublicKey() string {
	return strings.Replace(os.Getenv("PUBLIC_KEY"), "\\n", "\n", -1)
}

func GetEndpointURL() string {
	return os.Getenv("ENDPOINT_URL")
}

func GetS3BucketName() string {
	return os.Getenv("S3_BUCKET")
}
