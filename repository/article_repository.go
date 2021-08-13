package repository

import (
	"database/sql"
	"strings"

	"github.com/js-cha/article-api/model"
)

type articleRepository struct {
	DB *sql.DB
}

type ArticleRepository interface {
	Get(id int) (model.Article, error)
}

func NewArticleRepository(db *sql.DB) articleRepository {
	return articleRepository{
		DB: db,
	}
}

func (a articleRepository) Get(id int) (article model.Article, err error) {
	var tags string
	err = a.DB.QueryRow(
		"SELECT a.id, a.title, a.body, a.date, group_concat(t.name) as tags FROM article as a INNER JOIN tag as t ON a.id = t.articleId WHERE a.id = ?", id,
	).Scan(&article.ID, &article.Title, &article.Body, &article.Date, &tags)
	article.Tags = strings.Split(tags, ",")
	return
}
