package service

import (
	"blog/api"
	"blog/db/models"
)

// Services in the interface for service
type Services interface {
	Article() api.Article
}
type services struct {
	article api.Article
}

func (svc *services) Article() api.Article {
	return svc.article
}

// Init initializes Service
func Init() Services {
	// instance.Init(config.NewConfig())
	// defer instance.Destroy()

	// config := in()
	// log.Fatal(config)
	var db = []models.Article{}

	return &services{
		article: api.NewArticle(
			db,
		),
	}
}
