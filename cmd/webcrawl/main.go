package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quandev306/webcrawl/api/route"
	"github.com/quandev306/webcrawl/bootstrap"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second
	r := gin.Default()
	route.Setup(env, timeout, db, r)
	r.Run(env.ServerAddress)
}
