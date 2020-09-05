package controllers

import (
	"blog/api/service"
	"blog/db/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Article interface {
	Create(c *gin.Context)
	List(c *gin.Context)
	Get(c *gin.Context)
	Update(c *gin.Context)
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

func (a *article) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	msg := "success"
	if err != nil {
		msg = fmt.Sprintf(`%v is not integer parsable`, c.Param("id"))
	}
	resp, err := a.svc.Article().Get(id)
	if err != nil {
		msg = err.Error()
	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"data":    resp,
	})
}

func (a *article) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	msg := "success"
	if err != nil {
		msg = err.Error()
	}
	var article models.Article
	c.BindJSON(&article)
	article.ID = id
	res, err := a.svc.Article().Update(article)
	if err != nil {
		msg = err.Error()
	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"data":    res,
	})

}

// NewArticle df
func NewArticle(svc service.Services) Article {
	return &article{
		svc: svc,
	}
}
