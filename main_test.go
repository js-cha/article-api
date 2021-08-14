package main_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/js-cha/article-api/app"
)

var test app.App

func TestMain(m *testing.M) {
	test = app.App{}
	// setup db
	test.Initialize("test.db")

	// create tables
	setupTables()

	// run all tests
	code := m.Run()

	// clean up tables
	cleanUpTables()

	os.Exit(code)
}

func setupTables() {
	const createArticleTable = `CREATE TABLE IF NOT EXISTS article (
		id INTEGER PRIMARY KEY, title TEXT NOT NULL, date TEXT NOT NULL, body TEXT NOT NULL
	)`

	const createTagTable = `CREATE TABLE IF NOT EXISTS tag (
		id INTEGER PRIMARY KEY, name TEXT NOT NULL, article_id INTEGER NOT NULL, date TEXT NOT NULL, UNIQUE(name, article_id)
	)`

	test.DB.Exec(createArticleTable)
	test.DB.Exec(createTagTable)

	insertArticle, _ := test.DB.Prepare("INSERT INTO article (title, date, body) VALUES (?, ?, ?)")
	insertArticle.Exec("article title 1", "2021-01-01", "article body 1")
	insertArticle.Exec("article title 2", "2021-01-02", "article body 2")
	insertArticle.Exec("article title 3", "2021-01-03", "article body 3")

	insertTag, _ := test.DB.Prepare("INSERT INTO tag (name, article_id, date) VALUES (?, ?, ?)")
	insertTag.Exec("food", 3, "2021-01-03")
	insertTag.Exec("restaurant", 3, "2021-01-03")
	insertTag.Exec("takeaway", 3, "2021-01-03")
}

func cleanUpTables() {
	test.DB.Exec("DELETE FROM article")
	test.DB.Exec("ALTER SEQUENCE article_id_seq RESTART WITH 1")
	test.DB.Exec("DELETE FROM tag")
	test.DB.Exec("ALTER SEQUENCE tag_id_seq RESTART WITH 1")
}

func TestGetWithInvalidIDReturnsBadRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "/articles/invalid-id", nil)
	res := httptest.NewRecorder()
	test.Router.ServeHTTP(res, req)

	if http.StatusBadRequest != res.Code {
		t.Errorf("Expected: %d. Actual: %d\n", http.StatusBadRequest, res.Code)
	}
}
