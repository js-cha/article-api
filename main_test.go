package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/js-cha/article-api/app"
)

var a app.App

func TestMain(m *testing.M) {
	a = app.App{}
	a.Initialize("test.db")

	code := m.Run()
	os.Exit(code)
}

func TestGetArticle(t *testing.T) {
	req, _ := http.NewRequest("GET", "/article/1", nil)
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	fmt.Println(rr)
}
