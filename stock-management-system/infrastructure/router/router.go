package router

import (
	"fmt"

	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/config"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/controllers"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/database"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/mail"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/repositories"
	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/token"
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
		{
			"/token",
			"GET",
			tokenEndPoint,
		},
		{
			"/login",
			"GET",
			login,
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

	switch err {
	case repositories.UserNotFoundErr:
		return response{
			StatusCode: 404,
		}, nil
	case nil:
	case err:
		return response{
			StatusCode: 500,
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

	if !mail.ValidEmailAddress(email) {
		return response{
			StatusCode: 400,
			Body:       "invalid email address",
		}, nil
	}

	codeRepository := &repositories.CodeRepository{
		UbicFoodHandler: *database.NewDynamoDBHandler().NewUbicFoodHandler(),
	}
	code, err := codeRepository.AddCodeToDB(email)
	if err != nil {
		return response{
			StatusCode: 500,
			Body:       "Failed to register email address",
		}, err
	}

	message := "以下のリンクへアクセスしてUBIC在庫管理システムへの登録を完了してください。\n" + config.GetEndpointURL() + "/dev/complete-register?code=" + code
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

func tokenEndPoint(request event) (response, error) {
	query := request.QueryStringParameters
	code, ok := query["code"]
	if !ok {
		return response{
			StatusCode: 400,
		}, nil
	}

	completeRepository := repositories.CompleteRepository{
		UbicFoodHandler: *database.NewDynamoDBHandler().NewUbicFoodHandler(),
	}
	email, err := completeRepository.CheckCode(code)
	if err != nil {
		return response{
			StatusCode: 400,
			Body:       "Invalid code: " + err.Error(),
		}, nil
	}

	id, isRegistered := completeRepository.IsUserRegistered(email)
	if !isRegistered {
		id, err = completeRepository.RegisterUser(email)
		if err != nil {
			return response{
				StatusCode: 500,
				Body:       "Failed to register user: " + err.Error(),
			}, nil
		}

		err = completeRepository.DeleteCode(code)
		if err != nil {
			return response{
				StatusCode: 500,
				Body:       "Failed to delete temporary code: " + err.Error(),
			}, nil
		}
	}

	token, err := token.GenerateToken(id, email)
	if err != nil {
		return response{
			StatusCode: 500,
			Body:       "Failed to generate token " + err.Error(),
		}, nil
	}

	return response{
		Body:       token,
		StatusCode: 200,
	}, nil
}

func login(request event) (response, error) {
	query := request.QueryStringParameters
	email, ok := query["email"]
	if !ok {
		return response{
			StatusCode: 400,
		}, nil
	}

	if !mail.ValidEmailAddress(email) {
		return response{
			StatusCode: 400,
			Body:       "invalid email address",
		}, nil
	}

	codeRepository := &repositories.CodeRepository{
		UbicFoodHandler: *database.NewDynamoDBHandler().NewUbicFoodHandler(),
	}
	code, err := codeRepository.AddCodeToDB(email)
	if err != nil {
		return response{
			StatusCode: 500,
			Body:       "Failed to register email address",
		}, err
	}

	message := "以下のリンクへアクセスしてJWTを取得してください \n" + config.GetEndpointURL() + "/dev/complete-register?code=" + code
	subject := "UBIC在庫管理システムログイン確認メール"
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
