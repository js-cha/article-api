package repository

import (
	"database/sql"
	"strings"

	"github.com/js-cha/article-api/model"
)

type ArticleRepository struct {
	DB *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{
		DB: db,
	}
}

func (a *ArticleRepository) Get(id int) (article model.Article, err error) {
	var tags sql.NullString
	err = a.DB.QueryRow(`
		SELECT a.id, a.title, a.date, a.body, group_concat(t.name) as tags 
		FROM article AS a 
		LEFT JOIN tag AS t 
		ON a.id = t.article_id
		WHERE a.id = ?
		GROUP BY a.id
	`, id).Scan(&article.ID, &article.Title, &article.Date, &article.Body, &tags)
	article.Tags = strings.Split(tags.String, ",")
	return
}
