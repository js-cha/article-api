package service

import (
	"github.com/js-cha/article-api/model"
	"github.com/js-cha/article-api/repository"
)

type articleService struct {
	articleRepository repository.ArticleRepository
}

type ArticleService interface {
	Get(id int) (model.Article, error)
}

func NewArticleService(r repository.ArticleRepository) articleService {
	return articleService{
		articleRepository: r,
	}
}

func (a articleService) Get(id int) (article model.Article, err error) {
	return a.articleRepository.Get(id)
}
