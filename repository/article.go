package repository

import (
	"blog/db/models"
	"log"

	"github.com/astaxie/beego/orm"
)

// ArticleRepo interface
type ArticleRepo interface {
	Save(a *models.Article) error
	FindOne(doc models.Article) (*models.Article, error)
	Update(doc *models.Article, fieldsToUpdate []string) error
	FindAll() ([]models.Article, error)
}

type articleRepo struct {
	db orm.Ormer
}

func (repo *articleRepo) Save(a *models.Article) error {
	id, err := repo.db.Insert(a)
	if err != nil {
		log.Println(err)
		return err
	}
	a.ID = id
	return nil
}

func (repo *articleRepo) FindOne(doc models.Article) (*models.Article, error) {
	qs := repo.db.QueryTable(new(models.Article))
	if doc.ID != 0 {
		qs = qs.Filter("id", doc.ID)
	}
	// if doc.Title != nil {
	// 	qs = qs.Filter("title", *doc.Title)
	// }
	var article models.Article
	err := qs.One(&article)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &article, nil
}

func (repo *articleRepo) Update(doc *models.Article, fieldsToUpdate []string) error {
	_, err := repo.db.Update(doc, fieldsToUpdate...)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (repo *articleRepo) FindAll() ([]models.Article, error) {
	qs := repo.db.QueryTable(new(models.Article))

	var articles []models.Article
	_, err := qs.All(&articles)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return articles, nil
}

// NewArticleRepo repo for new article
func NewArticleRepo(db orm.Ormer) ArticleRepo {
	return &articleRepo{
		db: db,
	}
}
