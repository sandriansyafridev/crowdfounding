package main

import (
	"github.com/sandriansyafridev/crowdfounding/app"
	"github.com/sandriansyafridev/crowdfounding/repository"
)

var (
	gormDB, sqlDB  = app.NewDB()
	userRepository = repository.NewUserRepositoryImpl(gormDB)
)

func main() {
	app.InitializeMigration(gormDB)
	defer sqlDB.Close()

	r := app.NewRoute()
	r.Run(":8080")

}
