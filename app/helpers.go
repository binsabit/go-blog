package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (app *Application) GetIntValueFromJWT(ctx *fiber.Ctx, key string) int {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	res := claims[key].(float64)
	return int(res)
}
