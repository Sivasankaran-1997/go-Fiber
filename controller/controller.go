package controller

import (
	"fiberscurd/domain"
	"fiberscurd/middleware"
	"fiberscurd/services"
	"fiberscurd/utils"
	"net/http"

	"strings"

	"github.com/gofiber/fiber"
)

func CreateUser(c *fiber.Ctx) {
	var newUser domain.User
	if err := c.BodyParser(&newUser); err != nil {
		resterr := utils.BadRequest("Invalid JSON")
		c.Status(resterr.Status).JSON(resterr)
		return
	}

	result, resterr := services.CreateUser(newUser)

	if resterr != nil {
		c.Status(resterr.Status).JSON(resterr)
		return
	}
	c.JSON(result)

}

func Login(c *fiber.Ctx) {
	var newUser domain.User
	if err := c.BodyParser(&newUser); err != nil {
		resterr := utils.BadRequest("Invalid JSON")
		c.Status(resterr.Status).JSON(resterr)
		return
	}

	result, resterr := services.LoginUser(newUser)

	if resterr != nil {
		c.Status(resterr.Status).JSON(resterr)
		return
	}
	c.JSON(result)

}

func GetUserByEmail(c *fiber.Ctx) {

	tokenString := c.Get("Authorization")

	if strings.TrimSpace(tokenString) == "" {
		resterr := utils.BadRequest("Empty Token Param")
		c.Status(resterr.Status).JSON(resterr)
		return
	}

	result, resterr := middleware.UsergetValidate(tokenString)

	if resterr != nil {
		c.Status(resterr.Status).JSON(resterr)
		return
	}
	c.JSON(result)

}

func DeleteUserByEmail(c *fiber.Ctx) {

	tokenString := c.Get("Authorization")

	if strings.TrimSpace(tokenString) == "" {
		resterr := utils.BadRequest("Empty Token Param")
		c.Status(resterr.Status).JSON(resterr)
		return
	}

	result, resterr := middleware.UserdeleteValidate(tokenString)

	if resterr != nil {
		c.Status(resterr.Status).JSON(resterr)
		return
	} else {
		getResult, err := services.DeleteUser(result.(string))

		if err != nil {
			c.Status(resterr.Status).JSON(resterr)
			return
		}
		c.JSON(getResult)
	}
}

func UpdateUserByEmail(c *fiber.Ctx) {
	tokenString := c.Get("Authorization")

	var newUser domain.User
	if strings.TrimSpace(tokenString) == "" {
		resterr := utils.BadRequest("Empty Token Param")
		c.Status(resterr.Status).JSON(resterr)
	}

	if err := c.BodyParser(&newUser); err != nil {
		resterr := utils.BadRequest("Invalid JSON")
		c.Status(resterr.Status).JSON(resterr)
		return
	}

	result, resterr := middleware.UserupdateValidate(tokenString)

	if resterr != nil {
		c.Status(resterr.Status).JSON(resterr)
		return
	} else {
		newUser.Email = result.(string)
		isParital := c.Route().Method == http.MethodPatch
		result, resterr := services.UpdateUser(isParital, newUser)

		if resterr != nil {
			c.Status(resterr.Status).JSON(resterr)
			return
		}
		c.JSON(result)
	}

}
