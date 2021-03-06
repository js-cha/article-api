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

func (a *ArticleService) Add(article model.Article) (int64, error) {
	return a.articleRepository.Add(article)
}

func (a *ArticleService) GetTag(tagName, date string) (tag model.Tag, err error) {
	return a.articleRepository.GetTag(tagName, date)
}
