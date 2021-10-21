package token

import (
	"fmt"
	"time"

	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/config"

	jwt "github.com/form3tech-oss/jwt-go"
)

var prvKey []byte = []byte(config.SignatureKey())

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
	return token.SignedString(prvKey)
}

func VerifyToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return prvKey, nil
	})
	return err
}
