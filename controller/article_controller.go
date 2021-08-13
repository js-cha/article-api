package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/js-cha/article-api/service"
)

type articleController struct {
	articleService service.ArticleService
}

type ArticleController interface {
	Get(w http.ResponseWriter, r *http.Request)
}

func NewArticleController(s service.ArticleService) articleController {
	return articleController{
		articleService: s,
	}
}

func (c articleController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
	}

	article, error := c.articleService.Get(id)
	if error != nil {
		log.Fatal(error)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(article)
}
