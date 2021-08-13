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

func (a *App) Initialize(dataSourceName string) {
	var err error
	a.DB, err = sql.Open("sqlite3", fmt.Sprintf("./db/%s", dataSourceName))
	if err != nil {
		log.Fatal(err)
	}

	articleRepository := repository.NewArticleRepository(a.DB)
	articleService := service.NewArticleService(articleRepository)
	articleController := controller.NewArticleController(articleService)

	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/article/{id}", articleController.Get)
}

func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.Router))
}
