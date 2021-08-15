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
		BadRequestResponse(w, "invalid article id")
		return
	}

	article, err := c.articleService.Get(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			NotFoundResponse(w, "article not found")
			return
		default:
			InternalServerErrorResponse(w, err.Error())
			return
		}
	}

	OKResponse(w, article)
}

func (c *ArticleController) Add(w http.ResponseWriter, r *http.Request) {
	var a model.Article
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&a); err != nil && err != io.EOF {
		BadRequestResponse(w, "invalid request body")
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

func (c *ArticleController) GetTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tagName := vars["tagName"]
	date := vars["date"]

	if tagName == "" || date == "" {
		BadRequestResponse(w, "invalid tag or date")
		return
	}

	tag, err := c.articleService.GetTag(tagName, date)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			NotFoundResponse(w, "tag not found")
			return
		default:
			InternalServerErrorResponse(w, err.Error())
			return
		}
	}

	OKResponse(w, tag)
}
