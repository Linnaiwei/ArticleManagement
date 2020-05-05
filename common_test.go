package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var tmpArticleList []article

func TestMain(m *testing.M){
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func getRouter(withTemplate bool) *gin.Engine {
	r := gin.Default()
	if withTemplate{
		r.LoadHTMLFiles("templates/*")
	}
	return r
}

func testHTTPResponse (t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if !f(w){
		t.Fail()
	}
}

func saveLists () {
	tmpArticleList = articleList
}

func restoreLists() {
	articleList = tmpArticleList
}
