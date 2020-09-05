package main

import (
	"blog/router"

	_ "github.com/lib/pq"
)

func main() {
	router.Init()
}
