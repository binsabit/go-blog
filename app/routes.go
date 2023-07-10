package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"github.con/binsabi/go-blog/middlewares"
)

func (app *Application) RegisterRoutes() {

	app.Router.Post("/signup", app.Signup)
	app.Router.Post("/login", app.Login)

	jwt := middlewares.Authenticate(app.Config.JWT)

	blog := app.Router.Group("/blog", jwt)
	blog.Get("", app.GetBlogs)
	blog.Get("/:blog_id", app.GetBlog)
	blog.Post("", app.CreateBlog)
	blog.Put("/:blog_id", app.UpdateBlog)
	blog.Delete("/:blog_id", app.DeleteBlog)

	app.Router.Get("/user/blogs/:user_id", jwt, app.GetBlogsForUser)

	//login - registrations route

	app.Router.Get("/comments/:blog_id", jwt, app.GetComments)
	comment := app.Router.Group("/comment", jwt)
	comment.Post("", app.CreateComment)
	comment.Get("/:comment_id", app.GetComment)
	comment.Put("/:comment_id", app.UpdateComment)
	comment.Delete("/:comment_id", app.DeleteComment)

	//Not found route
	app.Router.Use(func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fasthttp.StatusNotFound)
	})

}
