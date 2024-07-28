package main

import (
	"github.com/gin-gonic/gin"
	routes "learn/jwt/api/v1/routes"
	"learn/jwt/boostrap"
)

func main() {

	app := boostrap.InitApp()
	env := app.ENV

	app.DB.Ping()
	defer app.DatabaseDisconnect()

	r := gin.Default()
	routes.Setup(env, app.DB, r)

	r.Run(env.ServerAddress)
}
