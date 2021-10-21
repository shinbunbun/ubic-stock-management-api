package router

import (
	"fmt"

	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/config"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/controllers"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/database"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/mail"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/repositories"
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
		{
			"/register",
			"GET",
			register,
		},
	}
	for _, route := range routes {
		if route.url == url && route.method == method {
			return route.function(request)
		}
	}

	return response{
		StatusCode: 400,
	}, nil
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
	if err == repositories.UserNotFoundErr {
		return response{
			StatusCode: 404,
		}, err
	}
	return response{
		Body:       body,
		StatusCode: 200,
	}, nil
}

func register(request event) (response, error) {
	query := request.QueryStringParameters
	email, ok := query["email"]
	if !ok {
		return response{
			StatusCode: 400,
		}, nil
	}

	fmt.Println(email)

	if !mail.ValidEmailAddress(email) {
		return response{
			StatusCode: 400,
			Body:       "invalid email address",
		}, nil
	}

	registerRepository := &repositories.RegisterRepository{
		UbicFoodHandler: *database.NewDynamoDBHandler().NewUbicFoodHandler(),
	}
	code, err := registerRepository.AddCodeToDB(email)
	if err != nil {
		return response{
			StatusCode: 500,
			Body:       "Failed to register email address",
		}, err
	}

	message := "以下のリンクへアクセスしてUBIC在庫管理システムへの登録を完了してください。\nhttps://urvuod6a7j.execute-api.ap-northeast-1.amazonaws.com/complete-register?code=" + code
	subject := "UBIC在庫管理システム登録確認メール"
	sender := config.SenderEmailAddress()
	fmt.Println(message, email, subject, sender, false)
	err = mail.SendMail(message, email, subject, sender, false)
	if err != nil {
		return response{
			StatusCode: 500,
			Body:       "Failed to send email",
		}, err
	}

	return response{
		Body:       "",
		StatusCode: 200,
	}, nil
}
