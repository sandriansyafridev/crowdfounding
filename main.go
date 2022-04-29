package main

import (
	"github.com/sandriansyafridev/crowdfounding/app"
)

var (
	db, sqlDB = app.NewDB()
)

func main() {
	app.NewMigration(db)

	defer sqlDB.Close()
}
