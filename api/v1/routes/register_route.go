package routes

import (
	"database/sql"
	"learn/jwt/api/v1/controller"
	"learn/jwt/boostrap"
	"learn/jwt/repository"
	"learn/jwt/usecases"

	"github.com/gin-gonic/gin"
)

func NewRegisterRouter(env *boostrap.Env, db *sql.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, "user")
	rc := &controller.RegisterController{
		RegisterUsecase: usecases.NewUserUseCase(ur),
		Env:             env,
	}
	group.POST("/register", rc.Register)
}
