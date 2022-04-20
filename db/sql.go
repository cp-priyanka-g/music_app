package db

import (
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

func NewSql() *sqlx.DB {

	var db *sqlx.DB

	username := "root"

	password := "root"

	host := "localhost"

	port := "3306"

	name := "music_app"

	db = sqlx.MustConnect("mysql", username+":"+password+"@("+host+":"+port+")/"+name)

	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	db.SetConnMaxLifetime(time.Minute * 1)

	return db
}
