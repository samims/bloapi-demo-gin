package api

import (
	"blog/db/models"
)

// Article interface
type Article interface {
	Add(data models.Article) (models.Article, error)
	List() ([]models.Article, error)
}

type article struct {
	db []models.Article
}

func (a *article) Add(data models.Article) (models.Article, error) {
	a.db = append(a.db, data)
	return data, nil
}

func (a *article) List() ([]models.Article, error) {
	return a.db, nil
}

// NewArticle init
func NewArticle(db []models.Article) Article {
	return &article{
		db: db,
	}
}
