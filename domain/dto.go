package domain

import (
	"fiberscurd/utils"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID       string `json:id"`
	Name     string `json:name bson:"name,omitempty"`
	Email    string `json:email bson:"email,omitempty"`
	Password string `json:password bson:"password,omitempty"`
}

type UserJWTsigneDetails struct {
	Email string
	jwt.RegisteredClaims
}

func (user *User) Vaildate() *utils.Resterr {
	if strings.TrimSpace(user.Name) == "" {
		return utils.BadRequest("Name Required")
	}

	if strings.TrimSpace(user.Email) == "" {
		return utils.BadRequest("Email Required")
	}

	if strings.TrimSpace(user.Password) == "" {
		return utils.BadRequest("PassWord Required")
	}

	return nil
}
