package controllers

import (
	"example/go-api/models"
	"example/go-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IdentityController interface {
	RegisterUser(context *gin.Context)
	GenerateToken(context *gin.Context)
}

type identityController struct {
	service services.IdentityService
}

func NewIdentityController(service services.IdentityService) IdentityController {
	return &identityController{
		service: service,
	}
}

func (controller identityController) RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := controller.service.RegisterUser(user); err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}
	context.Status(http.StatusOK)
}

func (controller identityController) GenerateToken(context *gin.Context) {
	var request services.TokenRequest
	request.Email = context.PostForm("username")
	request.Password = context.PostForm("password")
	token, err := controller.service.GetToken(request);
	if err != nil{
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}
	context.JSON(http.StatusCreated, gin.H{"access_token":&token})
}