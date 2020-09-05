package api

import (
	"blog/db/models"
	"blog/repository"
	"fmt"
	"log"

	"github.com/astaxie/beego/orm"
)

// Article interface
type Article interface {
	Add(data models.Article) (models.Article, error)
	Get(id int64) (*models.Article, error)
	Update(doc models.Article) (*models.Article, error)
	List() ([]models.Article, error)
}

type article struct {
	articleRepo repository.ArticleRepo
}

func (a *article) Add(data models.Article) (models.Article, error) {
	err := a.articleRepo.Save(&data)
	return data, err
}

func (a *article) Get(id int64) (*models.Article, error) {
	articleQ := models.Article{
		ID: id,
	}
	article, err := a.articleRepo.FindOne(articleQ)
	if err != nil {
		if err == orm.ErrNoRows {
			return nil, fmt.Errorf(`Article not found for %v`, id)
		}
		return nil, err
	}
	return article, nil
}

func (a *article) Update(doc models.Article) (*models.Article, error) {
	articleDoc := models.Article{
		ID: doc.ID,
	}
	article, err := a.articleRepo.FindOne(articleDoc)
	if err != nil {
		if err == orm.ErrNoRows {
			log.Println(fmt.Errorf(`Article not found for id %v`, doc.ID))
			return nil, fmt.Errorf(`Article not found for id %v`, doc.ID)
		}
		return nil, err
	}

	if doc.Title != nil {
		article.Title = doc.Title
	}
	if doc.Description != nil {
		article.Description = doc.Description
	}
	err = a.articleRepo.Update(article, []string{})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return article, nil

}

func (a *article) List() ([]models.Article, error) {
	return a.articleRepo.FindAll()
}

// NewArticle init
func NewArticle(articleRepo repository.ArticleRepo) Article {
	return &article{
		articleRepo: articleRepo,
	}
}
