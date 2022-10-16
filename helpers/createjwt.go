package helpers

import (
	"fiberscurd/domain"
	"fiberscurd/utils"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(claims *domain.UserJWTsigneDetails) (string, *utils.Resterr) {

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", utils.InternalErr("Not Token Created")
	}

	return token, nil
}
