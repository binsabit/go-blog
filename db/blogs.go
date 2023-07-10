package db

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	boiler "github.con/binsabi/go-blog/boil/psql/models"
)

func CreateBlog(ctx context.Context, db *sql.DB, blog *boiler.Blog) error {

	err := blog.Insert(ctx, db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func UpdateBlog(ctx context.Context, db *sql.DB, title, content string, id int) error {
	blogPost, err := boiler.FindBlog(ctx, db, id)
	if err != nil {
		return err
	}

	blogPost.Content = content
	blogPost.Title = title

	_, err = blogPost.Update(ctx, db, boil.Infer())

	return err
}

func DeleteBlog(ctx context.Context, db *sql.DB, id int) error {
	blog, err := boiler.FindBlog(ctx, db, id)
	if err != nil {
		return err
	}

	_, err = blog.Delete(ctx, db)

	return err
}

func GetBlogWithID(ctx context.Context, db *sql.DB, id int) (*boiler.Blog, error) {
	blog, err := boiler.FindBlog(ctx, db, id)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func GetAllBlogs(ctx context.Context, db *sql.DB) (boiler.BlogSlice, error) {
	blogs, err := boiler.Blogs().All(ctx, db)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func GetAllBlogsForUser(ctx context.Context, db *sql.DB, userID int) (boiler.BlogSlice, error) {
	blogs, err := boiler.Blogs(qm.WhereIn("user_id=?", userID)).All(ctx, db)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}
