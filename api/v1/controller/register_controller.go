package controller

import (
	"learn/jwt/boostrap"
	"learn/jwt/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	RegisterUsecase domain.RegisterUsecase
	Env             *boostrap.Env
}

func (rc *RegisterController) Register(c *gin.Context) {
	var request domain.RegisterRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = rc.RegisterUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	user := domain.User{
		ID:       1,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = rc.RegisterUsecase.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// accessToken, err := rc.RegisterUsecase.CreateAccessToken(&user, rc.Env.AccessTokenSecret, rc.Env.AccessTokenExpiryHour)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	// 	return
	// }

	// refreshToken, err := rc.RegisterUsecase.CreateRefreshToken(&user, rc.Env.RefreshTokenSecret, rc.Env.RefreshTokenExpiryHour)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	// 	return
	// }

	signupResponse := domain.RegisterResponse{
		AccessToken:  "",
		RefreshToken: "",
	}

	c.JSON(http.StatusOK, signupResponse)

}
