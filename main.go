package main

import (
	"github.com/enkelm/go_api/api"
	"github.com/enkelm/go_api/db"
	"github.com/enkelm/go_api/util"
)

func main() {
	util.InitEnvironmentVariables()
	db.InitMigration()
	api.Init()
}
