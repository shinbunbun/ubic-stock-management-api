package main

import (
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/router"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(router.Router)
}
