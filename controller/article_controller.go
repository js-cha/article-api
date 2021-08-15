package controller

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/js-cha/article-api/model"
	"github.com/js-cha/article-api/service"
)

type ArticleController struct {
	articleService *service.ArticleService
}

func NewArticleController(s *service.ArticleService) *ArticleController {
	return &ArticleController{
		articleService: s,
	}
}

func (c *ArticleController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		BadRequestResponse(w, "Invalid article ID")
		return
	}

	article, error := c.articleService.Get(id)
	if error != nil {
		switch error {
		case sql.ErrNoRows:
			NotFoundResponse(w, "Article not found")
			return
		default:
			InternalServerErrorResponse(w, error.Error())
			return
		}
	}

	OKResponse(w, article)
}

func (c *ArticleController) Add(w http.ResponseWriter, r *http.Request) {
	var a model.Article
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&a); err != nil && err != io.EOF {
		BadRequestResponse(w, "Invalid request body")
		return
	}
	defer r.Body.Close()
	id, err := c.articleService.Add(a)
	if err != nil {
		InternalServerErrorResponse(w, err.Error())
		return
	}

	CreatedResponse(w, map[string]int64{"id": id})
}
