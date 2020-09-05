package repository

import (
	"blog/db/models"

	"github.com/astaxie/beego/orm"
)

type ArticleRepo interface {
	Save(a *models.Article) error
	FindAll() ([]models.Article, error)
}

type articleRepo struct {
	db orm.Ormer
}

func (repo *articleRepo) Save(a *models.Article) error {
	id, err := repo.db.Insert(a)
	if err != nil {
		log.Error(err)
		return err
	}
	a.ID = id
	return nil
}

func (repo *articleRepo) FindAll() ([]models.Article, error) {
	qs := repo.db.QueryTable(new(models.Article))

	var articles []models.Article
	_, err := qs.All(&articles)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return articles, nil
}

func NewArticleRepo(db orm.Ormer) ArticleRepo {
	return &articleRepo{
		db: db,
	}
}
