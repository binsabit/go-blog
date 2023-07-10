package db

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	boiler "github.con/binsabi/go-blog/boil/psql/models"
)

func MakeComment(ctx context.Context, db *sql.DB, comment *boiler.Comment) error {
	err := comment.Insert(ctx, db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(ctx context.Context, db *sql.DB, id int) error {
	comment, err := GetComment(ctx, db, id)
	if err != nil {
		return err
	}
	_, err = comment.Delete(ctx, db)

	return err
}

func EditComment(ctx context.Context, db *sql.DB, id int, content string) error {
	comment, err := GetComment(ctx, db, id)
	if err != nil {
		return err
	}
	comment.Content = content

	_, err = comment.Update(ctx, db, boil.Infer())
	return err
}

func GetCommentsForBlog(ctx context.Context, db *sql.DB, blogID int) (boiler.CommentSlice, error) {
	comments, err := boiler.Comments(qm.Where("blog_id=?", blogID)).All(ctx, db)
	if err != nil {
		return nil, err
	}
	return comments, err
}

func GetComment(ctx context.Context, db *sql.DB, id int) (*boiler.Comment, error) {
	comment, err := boiler.FindComment(ctx, db, id)
	if err != nil {
		return nil, err
	}
	return comment, nil

}
