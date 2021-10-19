package router

import (
	"errors"

	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/controllers"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/database"
	"github.com/aws/aws-lambda-go/events"
)

var (
	controller *controllers.Controller
)

type response events.APIGatewayProxyResponse
type event events.APIGatewayProxyRequest

func init() {
	db := database.NewDynamoDBHandler()
	controller = controllers.NewController(db)
}

func Router(request event) (response, error) {
	url := request.Path
	method := request.HTTPMethod
	routes := []struct {
		url      string
		method   string
		function func(event) (response, error)
	}{
		{
			"/user",
			"GET",
			findUserByID,
		},
	}

	for _, route := range routes {
		if route.url == url && route.method == method {
			return route.function(request)
		}
	}

	return response{
		StatusCode: 400,
	}, errors.New("Invalid request error")
}

func findUserByID(request event) (response, error) {
	query := request.QueryStringParameters
	id, ok := query["id"]
	if !ok {
		return response{
			StatusCode: 400,
		}, nil
	}
	body, err := controller.FindUserByID(id)
	if err != nil {
		return response{
			StatusCode: 500,
		}, err
	}
	return response{
		Body:       body,
		StatusCode: 200,
	}, nil
}
