package router

import (
	"encoding/json"

	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/database"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/repositories"
)

func userGet(request event) (response, error) {
	query := request.QueryStringParameters
	id, ok := query["id"]
	if !ok {
		return response{
			StatusCode: 400,
		}, nil
	}

	userRepository := &repositories.UserRepository{
		UbicFoodHandler: database.NewDynamoDBHandler().NewUbicFoodHandler(),
	}
	user, err := userRepository.FindByID(id)
	if err != nil {
		return response{
			StatusCode: 500,
			Body:       "Cannot find user infomation: " + err.Error(),
		}, nil
	}

	jsonBody, err := json.Marshal(user)
	if err != nil {
		return response{
			StatusCode: 500,
			Body:       "Failed to generate response json: " + err.Error(),
		}, nil
	}
	return response{
		StatusCode: 200,
		Body:       string(jsonBody),
	}, nil
}
