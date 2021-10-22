package token

import (
	"fmt"
	"time"

	"github.com/Yuto/ubic-stock-management-api/stock-management-system/infrastructure/config"
	jwt "github.com/dgrijalva/jwt-go"
)

var prvKey []byte = []byte(config.PrivateKey())
var pubKey []byte = []byte(config.PublicKey())

type JwtClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(id, email string) (string, error) {
	ecdsaKey, err := jwt.ParseECPrivateKeyFromPEM(prvKey)
	if err != nil {
		return "", err
	}

	// headerのセット
	token := jwt.New(jwt.SigningMethodES512)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["iss"] = "ubic-food-stock-management-system"
	claims["sub"] = id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	claims["iat"] = time.Now().Unix()
	claims["email"] = email
	claims["role"] = "user"

	// 電子署名
	return token.SignedString(ecdsaKey)
}

func VerifyToken(tokenString string) (JwtClaims, error) {
	jwt, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		ecdsaKey, err := jwt.ParseECPublicKeyFromPEM(pubKey)
		if err != nil {
			return "", err
		}
		return ecdsaKey, nil
	})
	if claims, ok := jwt.Claims.(*JwtClaims); ok && jwt.Valid {
		return *claims, nil
	} else {
		return *claims, err
	}
}
