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

func init() {
	db := database.NewDynamoDBHandler()
	controller = controllers.NewController(db)
}

func Router(request events.APIGatewayProxyRequest) (response, error) {
	url := request.Path
	method := request.HTTPMethod
	query := request.QueryStringParameters
	routes := []struct {
		url      string
		method   string
		function func(map[string]string) (response, error)
	}{
		{
			"/user",
			"GET",
			findUserByID,
		},
	}

	for _, route := range routes {
		if route.url == url && route.method == method {
			return route.function(query)
		}
	}

	return response{
		StatusCode: 400,
	}, errors.New("Invalid request error")
}

func findUserByID(query map[string]string) (response, error) {
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
