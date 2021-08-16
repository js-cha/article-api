package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/js-cha/article-api/controller"
	"github.com/js-cha/article-api/repository"
	"github.com/js-cha/article-api/service"
)

type App struct {
	DB     *sql.DB
	Router *mux.Router
}

const (
	CreateTableArticle = `CREATE TABLE IF NOT EXISTS article (
		id INTEGER PRIMARY KEY, title TEXT NOT NULL, date TEXT NOT NULL, body TEXT NOT NULL
	)`
	CreateTableTag = `CREATE TABLE IF NOT EXISTS tag (
		id INTEGER PRIMARY KEY, tag_name TEXT NOT NULL, article_id INTEGER NOT NULL, date TEXT NOT NULL, UNIQUE(tag_name, article_id)
	)`
)

func (a *App) Initialize(dataSourceName string) {
	var err error
	a.DB, err = sql.Open("sqlite3", fmt.Sprintf("./%s", dataSourceName))
	if err != nil {
		log.Fatal(err)
	}

	a.DB.Exec(CreateTableArticle)
	a.DB.Exec(CreateTableTag)

	articleRepository := repository.NewArticleRepository(a.DB)
	articleService := service.NewArticleService(articleRepository)
	articleController := controller.NewArticleController(articleService)

	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/articles", articleController.Add).Methods("POST")
	a.Router.HandleFunc("/articles/{id}", articleController.Get).Methods("GET")
	a.Router.HandleFunc("/tags/{tagName}/{date}", articleController.GetTag).Methods("GET")
}

func (a *App) Run(port string) {
	fmt.Printf("Running server at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, a.Router))
}
