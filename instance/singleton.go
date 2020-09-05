package instance

import (
	"blog/config"
	"fmt"
	"sync"

	"github.com/astaxie/beego/orm"
)

type instance struct {
	db     orm.Ormer
	config config.Config
}

var singleton = &instance{}
var once sync.Once

func Init(config config.Config) {
	once.Do(func() {
		singleton.config = config

		orm.RegisterDriver("postgres", orm.DRPostgres)
		conString := fmt.Sprintf(`postgres://%v:%v@%v:%v/%v?sslmode=disable`,
			singleton.config.GetDBUser(), singleton.config.GetDBPass(), singleton.config.GetDBHost(), singleton.config.GetDBPort(), singleton.config.GetDBName())

		orm.RegisterDataBase("default", "postgres", conString)

		orm.RunSyncdb("default", false, true)

		singleton.db = orm.NewOrm()
		singleton.db.Using("default")
	})
}

func Destroy() error {
	singleton.db = nil
	return nil
}

func DB() orm.Ormer {
	return singleton.db
}
