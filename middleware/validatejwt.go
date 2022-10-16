package middleware

import (
	"fmt"

	"fiberscurd/services"
	"fiberscurd/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var User interface{}

func ValidateToken(tokens string) (interface{}, *utils.Resterr) {
	token, _ := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		User = claims["sub"]
	} else if float64(time.Now().Unix()) > claims["exp"].(float64) && token.Valid {
		return nil, utils.BadRequest("JWT Expired")
	} else {
		return nil, utils.InternalErr("Token Not Valid")
	}

	return User, nil
}

func UsergetValidate(tokens string) (interface{}, *utils.Resterr) {

	tokenValue, err := ValidateToken(tokens)

	if err != nil {
		return nil, err
	}

	result, resterr := services.GetByID(tokenValue.(string))

	if resterr != nil {
		return nil, utils.NotFound("User Not Found")
	}

	return result, nil
}

func UserdeleteValidate(tokens string) (interface{}, *utils.Resterr) {

	tokenValue, err := ValidateToken(tokens)

	if err != nil {
		return nil, err
	}

	result, resterr := services.GetByID(tokenValue.(string))

	if resterr != nil {
		return nil, utils.NotFound("User Not Found")
	}

	return result.Email, nil
}

func UserupdateValidate(tokens string) (interface{}, *utils.Resterr) {

	tokenValue, err := ValidateToken(tokens)

	if err != nil {
		return nil, err
	}

	result, resterr := services.GetByID(tokenValue.(string))

	if resterr != nil {
		return nil, utils.NotFound("User Not Found")
	}

	return result.Email, nil
}
