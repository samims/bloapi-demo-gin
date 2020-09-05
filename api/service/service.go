package service

import (
	"blog/api"
	"blog/repository"

	"github.com/astaxie/beego/orm"
)

// Services in the interface for service
type Services interface {
	Article() api.Article // This is article service interface
}
type services struct {
	article api.Article
}

func (svc *services) Article() api.Article {
	return svc.article
}

// Init initializes Service
func Init(db orm.Ormer) Services {
	// instance.Init(config.NewConfig())
	// defer instance.Destroy()

	// config := in()
	// log.Fatal(config)

	// ar := repository.NewArticleRepo(db)

	// aS := api.NewArticle(ar)

	return &services{
		article: api.NewArticle(
			repository.NewArticleRepo(db),
		),
	}
}
