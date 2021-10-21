package token

import (
	"time"

	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/config"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
)

func GenerateToken(id, email string) (string, error) {

	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "ubic-food-stock-management-system"
	claims["sub"] = id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	claims["iat"] = time.Now()
	claims["email"] = email
	claims["role"] = "user"

	// 電子署名
	return token.SignedString([]byte(config.SignatureKey()))
}

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SignatureKey()), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
