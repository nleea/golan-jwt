package routes

import (
	"database/sql"
	"learn/jwt/boostrap"

	"github.com/gin-gonic/gin"
)

func Setup(env *boostrap.Env, db *sql.DB, gin *gin.Engine) {

	publicRouter := gin.Group("")
	NewRegisterRouter(env, db, publicRouter)

}
