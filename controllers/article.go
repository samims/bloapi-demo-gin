package controllers

import (
	"blog/api/service"
	"blog/db/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article interface {
	Create(c *gin.Context)
	List(c *gin.Context)
}

type article struct {
	svc service.Services
}

func (a *article) Create(c *gin.Context) {
	var article models.Article

	c.BindJSON(&article)

	res, err := a.svc.Article().Add(article)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": res,
	})

}

func (a *article) List(c *gin.Context) {
	resp, err := a.svc.Article().List()
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"data": resp,
	})
}

// NewArticle df
func NewArticle(svc service.Services) Article {
	return &article{
		svc: svc,
	}
}
