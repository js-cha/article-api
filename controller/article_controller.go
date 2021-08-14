package controller

import (
	"database/sql"
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
		BadRequestResponse(w)
		return
	}

	article, error := c.articleService.Get(id)
	if error != nil {
		switch error {
		case sql.ErrNoRows:
			NotFoundResponse(w)
			return
		default:
			InternalServerErrorResponse(w)
			return
		}
	}

	OKResponse(w, article)
}
