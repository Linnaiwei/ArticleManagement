package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	render(
			c,
		gin.H{
			"title": "home page",
			"payload": articles,
		},
		"index.html",
		)
}

func getArticle(c *gin.Context) {
	articleId := c.Param("article_id")
	id, err1 := strconv.Atoi(articleId)
	article, err := getArticleById(id)
	if err == nil && err1 == nil {
		c.HTML(
			http.StatusOK,
			"article.html",
			gin.H{
				"Content": article.Content,
			},
		)
	} else {
		c.HTML(
				http.StatusBadRequest,
				"400.html",
				gin.H{
					"Content": "Article does not exist!",
				},
			)
	}

}


func getArticleById (id int) (*article, error) {
	for _, article := range getAllArticles(){
		if article.Id == id {
			return &article, nil
		}
	}
	return nil, errors.New("article not found")
}
