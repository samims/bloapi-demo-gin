package main

import (
	"blog/config"
	"blog/router"

	_ "github.com/lib/pq"
)

func main() {
	db := config.GetDB()
	router.Init(db)
}
