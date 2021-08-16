package main

import (
	"github.com/js-cha/article-api/app"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := app.App{}
	app.Initialize("article.db")
	app.Run(":8080")
}
