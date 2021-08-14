package service

import (
	"github.com/js-cha/article-api/model"
	"github.com/js-cha/article-api/repository"
)

type ArticleService struct {
	articleRepository *repository.ArticleRepository
}

func NewArticleService(r *repository.ArticleRepository) *ArticleService {
	return &ArticleService{
		articleRepository: r,
	}
}

func (a *ArticleService) Get(id int) (article model.Article, err error) {
	return a.articleRepository.Get(id)
}
