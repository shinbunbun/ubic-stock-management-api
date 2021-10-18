package router

import (
	"errors"
	"strings"

	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/controllers"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/database"
	"github.com/aws/aws-lambda-go/events"
)

var (
	controller *controllers.Controller
)

func init() {
	db := database.NewDynamoDBHandler()
	controller = controllers.NewController(db)
}

func router(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	url := request.Path
	method := request.HTTPMethod
	query := request.QueryStringParameters
	response := events.APIGatewayProxyResponse{
		Body:       "",
		StatusCode: 200,
	}
	if strings.HasSuffix(url, "user") {
		switch method {
		case "GET":
			id, ok := query["id"]
			if !ok {
				response.StatusCode = 403
				return response, nil
			}
			body, err := controller.FindUserByID(id)
			if err != nil {
				return response, err
			}
			response.Body = body
			return response, nil
		}
	}
	return response, errors.New("Invalid request error")
}
