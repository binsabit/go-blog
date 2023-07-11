package app

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"github.com/volatiletech/null/v8"
	boiler "github.con/binsabi/go-blog/boil/psql/models"
	"github.con/binsabi/go-blog/db"
	lib_json "github.con/binsabi/go-blog/lib/json"
)

func (app *Application) GetBlogs(ctx *fiber.Ctx) error {
	blogs, err := db.GetAllBlogs(context.Background(), app.Storage)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"blogs": blogs})
}

func (app *Application) GetBlogsForUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("user_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user_id parameter",
		})
	}
	blogs, err := db.GetAllBlogsForUser(context.Background(), app.Storage, id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"blogs": blogs})

}

func (app *Application) GetBlog(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("blog_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid blog_id parameter",
		})
	}
	blog, err := db.GetBlogWithID(context.Background(), app.Storage, id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"blog": blog})

}

func (app *Application) CreateBlog(ctx *fiber.Ctx) error {
	var blog boiler.Blog

	err := lib_json.DecodeJSONBody(ctx, &blog)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	userID := app.GetIntValueFromJWT(ctx, "user_id")
	blog.UserID = null.Int{Int: userID, Valid: true}
	err = db.CreateBlog(context.Background(), app.Storage, &blog)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fiber.StatusCreated)

}

func (app *Application) UpdateBlog(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("blog_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid blog_id parameter",
		})
	}
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	err = lib_json.DecodeJSONBody(ctx, &input)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	err = db.UpdateBlog(context.Background(), app.Storage, input.Title, input.Content, id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.SendStatus(fasthttp.StatusOK)

}

func (app *Application) DeleteBlog(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("blog_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid blog_id parameter",
		})
	}
	err = db.DeleteBlog(context.Background(), app.Storage, id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.SendStatus(fiber.StatusOK)
}
