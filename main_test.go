package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/js-cha/article-api/app"
	"github.com/js-cha/article-api/model"
)

// global test app
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
	test.DB.Exec(app.CreateTableArticle)
	test.DB.Exec(app.CreateTableTag)
}

func cleanUpTables() {
	test.DB.Exec("DELETE FROM article")
	test.DB.Exec("ALTER SEQUENCE article_id_seq RESTART WITH 1")
	test.DB.Exec("DELETE FROM tag")
	test.DB.Exec("ALTER SEQUENCE tag_id_seq RESTART WITH 1")
}

func addArticle(title, date, body string) int64 {
	stmt, _ := test.DB.Prepare(`INSERT INTO article (title, date, body) VALUES (?, ?, ?)`)
	result, _ := stmt.Exec(title, date, body)
	id, _ := result.LastInsertId()
	return id
}

func addTags(id int64, tags []string, date string) {
	for _, s := range tags {
		stmt, _ := test.DB.Prepare(`INSERT INTO tag (tag_name, article_id, date) VALUES(?, ?, ?)`)
		stmt.Exec(s, id, date)
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func TestGetInvalidId(t *testing.T) {
	// arrange
	req, _ := http.NewRequest("GET", "/articles/invalid-id", nil)
	res := httptest.NewRecorder()

	// act
	test.Router.ServeHTTP(res, req)

	// assert
	if http.StatusBadRequest != res.Code {
		t.Errorf("Expected: %d. Actual: %d\n", http.StatusBadRequest, res.Code)
	}
}

func TestGetNonExistentArticle(t *testing.T) {
	// arrange
	req, _ := http.NewRequest("GET", "/articles/999", nil)
	res := httptest.NewRecorder()

	// act
	test.Router.ServeHTTP(res, req)

	// assert
	if http.StatusNotFound != res.Code {
		t.Errorf("Expected: %d. Actual: %d\n", http.StatusNotFound, res.Code)
	}
}

func TestGetArticle(t *testing.T) {
	// arrange
	cleanUpTables()
	title := "test article"
	date := "2021-01-01"
	body := "awesome content"
	tags := []string{"fitness", "health", "science"}

	id := addArticle(title, date, body)
	addTags(id, tags, date)

	req, _ := http.NewRequest("GET", "/articles/1", nil)
	res := httptest.NewRecorder()

	// act
	test.Router.ServeHTTP(res, req)

	// assert
	if http.StatusOK != res.Code {
		t.Errorf("Expected: %d. Actual: %d\n", http.StatusOK, res.Code)
	}

	var article model.Article
	json.Unmarshal(res.Body.Bytes(), &article)

	if id != article.ID {
		t.Errorf("Expected: %d. Actual: %d\n", id, article.ID)
	}

	if title != article.Title {
		t.Errorf("Expected: %s. Actual: %s\n", title, article.Title)
	}

	if date != article.Date {
		t.Errorf("Expected: %s. Actual: %s\n", date, article.Date)
	}

	if body != article.Body {
		t.Errorf("Expected: %s. Actual: %s\n", body, article.Body)
	}

	if !reflect.DeepEqual(tags, article.Tags) {
		t.Errorf("Expected: %v. Actual: %v\n", tags, article.Tags)
	}
}

func TestAddArticle(t *testing.T) {
	// arrange
	cleanUpTables()
	body := []byte(`{"title": "test article", "date": "2021-01-01", "body": "test body", "tags": ["health", "fitness", "science"]}`)
	req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// act
	res := httptest.NewRecorder()
	test.Router.ServeHTTP(res, req)

	// assert
	if http.StatusCreated != res.Code {
		t.Errorf("Expected: %d. Actual: %d\n", http.StatusCreated, res.Code)
	}

	var response map[string]int64
	json.Unmarshal(res.Body.Bytes(), &response)

	if response["id"] != 1 {
		t.Errorf("Expected: %d. Actual: %d\n", 1, response["id"])
	}
}

func TestAddArticleInvalidBody(t *testing.T) {
	// arrange
	body := []byte(`{"gibberish: "blah blah blah}`)
	req, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	// act
	res := httptest.NewRecorder()
	test.Router.ServeHTTP(res, req)

	// assert
	if http.StatusBadRequest != res.Code {
		t.Errorf("Expected: %d. Actual: %d\n", http.StatusBadRequest, res.Code)
	}
}

func TestGetTag(t *testing.T) {
	// arrange
	cleanUpTables()
	id := addArticle("test article", "2021-01-01", "awesome content")
	addTags(id, []string{"fitness", "health", "science"}, "2021-01-01")

	req, _ := http.NewRequest("GET", "/tags/health/2021-01-01", nil)
	req.Header.Set("Content-Type", "application/json")

	// act
	res := httptest.NewRecorder()
	test.Router.ServeHTTP(res, req)

	// assert
	var tag model.Tag
	json.Unmarshal(res.Body.Bytes(), &tag)
	if http.StatusOK != res.Code {
		t.Errorf("Expected: %d. Actual: %d\n", http.StatusOK, res.Code)
	}

	if tag.Tag != "health" {
		t.Errorf("Expected: %s. Actual: %s\n", "health", tag.Tag)
	}

	if tag.Count != len(tag.Articles) {
		t.Errorf("Expected: %d. Actual: %d\n", tag.Count, tag.Count)
	}

	if !contains(tag.Articles, strconv.FormatInt(id, 10)) {
		t.Errorf("Expected: %d to be found in %v but not found\n", id, tag.Articles)
	}

	if !contains(tag.Related_Tags, "fitness") {
		t.Errorf("Expected: %s to be found in %v but not found\n", "fitness", tag.Articles)
	}

	if !contains(tag.Related_Tags, "science") {
		t.Errorf("Expected: %s to be found in %v but not found\n", "science", tag.Articles)
	}
}

func TestGetTagNonExistentTAg(t *testing.T) {
	// arrange
	req, _ := http.NewRequest("GET", "/tags/sometag/2050-01-01", nil)
	req.Header.Set("Content-Type", "application/json")

	// act
	res := httptest.NewRecorder()
	test.Router.ServeHTTP(res, req)

	// assert
	if http.StatusNotFound != res.Code {
		t.Errorf("Expected: %d. Actual: %d\n", http.StatusNotFound, res.Code)
	}
}
