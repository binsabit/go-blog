package app

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/valyala/fasthttp"
	"github.con/binsabi/go-blog/db"
	lib_json "github.con/binsabi/go-blog/lib/json"
)

type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (app *Application) Signup(ctx *fiber.Ctx) error {
	var newUser Request
	err := lib_json.DecodeJSONBody(ctx, &newUser)

	if err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = db.CreateUser(context.Background(), app.Storage,
		newUser.Username, newUser.Password)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fasthttp.StatusCreated)

}

func (app *Application) Login(ctx *fiber.Ctx) error {
	var newUser Request
	err := lib_json.DecodeJSONBody(ctx, &newUser)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := db.CheckCredentials(context.Background(), app.Storage,
		newUser.Username, newUser.Password)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	claims := jtoken.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(app.Config.JWT.Expires).Unix(),
	}

	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(app.Config.JWT.Secret))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(LoginResponse{
		Token: t,
	})
}
