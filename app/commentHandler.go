package app

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	boiler "github.con/binsabi/go-blog/boil/psql/models"
	"github.con/binsabi/go-blog/db"
	lib_json "github.con/binsabi/go-blog/lib/json"
)

func (app *Application) GetComment(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("comment_id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid comment_id parameter",
		})
	}

	comment, err := db.GetComment(context.Background(), app.Storage, id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"comment": comment,
	})

}
func (app *Application) GetComments(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("blog_id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid blog_id parameter",
		})
	}

	comments, err := db.GetCommentsForBlog(context.Background(), app.Storage, id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"comments": comments,
	})
}
func (app *Application) CreateComment(ctx *fiber.Ctx) error {
	var comment boiler.Comment

	err := lib_json.DecodeJSONBody(ctx, &comment)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	err = db.MakeComment(context.Background(), app.Storage, &comment)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fasthttp.StatusCreated)
}
func (app *Application) UpdateComment(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("comment_id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid comment_id parameter",
		})
	}

	var input struct {
		Content string `json:"content"`
	}

	err = lib_json.DecodeJSONBody(ctx, &input)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = db.EditComment(context.Background(), app.Storage, id, input.Content)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.SendStatus(fasthttp.StatusOK)
}

func (app *Application) DeleteComment(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("comment_id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid comment_id parameter",
		})
	}

	err = db.DeleteComment(context.Background(), app.Storage, id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.SendStatus(fasthttp.StatusOK)
}
