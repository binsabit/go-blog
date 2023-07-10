package db

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"

	"github.com/volatiletech/sqlboiler/v4/boil"
	boiler "github.con/binsabi/go-blog/boil/psql/models"
)

var (
	ErrorAlreadyExists      = errors.New("user already exists")
	ErrorNotFound           = errors.New("user not found")
	ErrorInvalidCredentials = errors.New("invalid creadentials")
)

func CreateUser(ctx context.Context, db *sql.DB, username, password string) error {
	user := boiler.User{
		Username: username,
		Password: genereteHash(password),
	}

	if Exists(ctx, db, username) {
		return ErrorAlreadyExists
	}

	err := user.Insert(ctx, db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func Exists(ctx context.Context, db *sql.DB, username string) bool {
	row, _ := boiler.Users(boiler.UserWhere.Username.EQ(username)).One(ctx, db)
	if row != nil {
		return true
	}
	return false
}

func CheckCredentials(ctx context.Context, db *sql.DB, username, password string) (*boiler.User, error) {
	test := boiler.User{
		Username: username,
		Password: genereteHash(password),
	}

	if !Exists(ctx, db, username) {
		return nil, ErrorNotFound
	}

	user, err := boiler.Users(boiler.UserWhere.Username.EQ(username)).One(ctx, db)

	if err != nil {
		return nil, ErrorInvalidCredentials
	}

	if string(user.Password) != string(test.Password) {
		return nil, ErrorInvalidCredentials
	}
	return user, nil
}

func genereteJWT(user boiler.User) string { return "" }

func genereteHash(text string) []byte {
	h := sha256.New()
	h.Write([]byte(text))

	res := h.Sum(nil)

	return res
}
