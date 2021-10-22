package router

import (
	"encoding/json"

	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/database"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/repositories"
)

type userGetResponseData struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

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

	body := userGetResponseData{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}
	jsonBody, err := json.Marshal(body)
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
