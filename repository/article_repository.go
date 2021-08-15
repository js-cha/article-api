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
		SELECT a.id, a.title, a.date, a.body, group_concat(t.tag_name) as tags 
		FROM article AS a 
		LEFT JOIN tag AS t 
		ON a.id = t.article_id
		WHERE a.id = ?
		GROUP BY a.id
	`, id).Scan(&article.ID, &article.Title, &article.Date, &article.Body, &tags)
	article.Tags = strings.Split(tags.String, ",")
	return
}

func (a *ArticleRepository) Add(article model.Article) (int64, error) {
	stmt, _ := a.DB.Prepare(`INSERT INTO article (title, date, body) VALUES (?, ?, ?)`)
	res, err := stmt.Exec(article.Title, article.Date, article.Body)
	if err != nil {
		return 0, err
	}
	stmt.Close()
	id, _ := res.LastInsertId()

	stmt, _ = a.DB.Prepare(`INSERT INTO tag (tag_name, article_id, date) VALUES (?, ?, ?)`)
	for _, v := range article.Tags {
		_, err := stmt.Exec(v, id, article.Date)
		if err != nil {
			return 0, err
		}
	}
	stmt.Close()

	return id, err
}

func (a *ArticleRepository) GetTag(tagName, date string) (tag model.Tag, err error) {
	var relatedTags string
	var articles string

	err = a.DB.QueryRow(`
		SELECT t1.tag_name, group_concat(t2.tag_name), group_concat(DISTINCT cast(t1.article_id as TEXT)), (SELECT COUNT(DISTINCT t1.tag_name) FROM tag WHERE date = t1.date)
		FROM tag AS t1
		LEFT JOIN tag AS t2
		ON t1.date = t2.date AND t1.tag_name != t2.tag_name
		WHERE t1.tag_name = ?
		AND t1.date = ?
		GROUP BY t1.tag_name
		ORDER BY t1.date DESC
		LIMIT 10;`, tagName, date).Scan(&tag.Tag, &relatedTags, &articles, &tag.Count)

	tag.Articles = strings.Split(articles, ",")
	tag.Related_Tags = strings.Split(relatedTags, ",")

	return
}
