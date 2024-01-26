package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"swaggo-handson/helper"
	"swaggo-handson/model"
)

type UserControllerInterface interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
}

type userController struct {
	ctx context.Context
}

func InitUserController(ctx context.Context) UserControllerInterface {
	return &userController{
		ctx: ctx,
	}
}

func (c *userController) Create(ctx *gin.Context) {
	newUser := model.User{}
	err := ctx.BindJSON(&newUser)
	if err != nil {
		errResponse := helper.BuildErrorResponse("error while binding json data", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}
	newUser.ID = getNewID()
	Users = append(Users, &newUser)
	response := helper.BuildResponse(newUser)
	ctx.JSON(http.StatusOK, response)
	return
}
func (c *userController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errResponse := helper.BuildErrorResponse("error while convert id param", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}
	newUser := model.User{}
	err = ctx.BindJSON(&newUser)
	if err != nil {
		errResponse := helper.BuildErrorResponse("error while binding json data", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}
	var user *model.User = nil
	index := 0
	for i, v := range Users {
		if v.ID == id {
			user = v
			index = i
		}
	}
	if user == nil {
		errResponse := helper.BuildErrorResponse(fmt.Sprintf("data dengan id[%d] tidak ditemukan", id), fmt.Errorf("data not found"))
		ctx.AbortWithStatusJSON(http.StatusNotFound, errResponse)
		return
	}
	if newUser.Name != "" {
		user.Name = newUser.Name
	}
	if newUser.Age != 0 {
		user.Age = newUser.Age
	}
	if newUser.IsMarriage != false {
		user.IsMarriage = newUser.IsMarriage
	}
	if newUser.Title != "" {
		user.Title = newUser.Title
	}
	Users[index] = user
	response := helper.BuildResponse(nil)
	ctx.JSON(http.StatusOK, response)
	return
}
func (c *userController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errResponse := helper.BuildErrorResponse("error while convert id param", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}
	var user *model.User = nil
	index := 0
	for i, v := range Users {
		if v.ID == id {
			user = v
			index = i
		}
	}
	if user == nil {
		errResponse := helper.BuildErrorResponse(fmt.Sprintf("data dengan id[%d] tidak ditemukan", id), fmt.Errorf("data not found"))
		ctx.AbortWithStatusJSON(http.StatusNotFound, errResponse)
		return
	}

	Users[index] = Users[len(Users)-1]
	//Users[len(Users)-1] = nil
	Users = Users[:len(Users)-1]
	response := helper.BuildResponse(nil)
	ctx.JSON(http.StatusOK, response)
	return
}
func (c *userController) GetAll(ctx *gin.Context) {
	response := helper.BuildResponse(Users)
	ctx.JSON(http.StatusOK, response)
	return
}
func (c *userController) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		errResponse := helper.BuildErrorResponse("error while convert id param", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errResponse)
		return
	}
	var user *model.User = nil
	for _, v := range Users {
		if v.ID == id {
			user = v
		}
	}
	if user == nil {
		errResponse := helper.BuildErrorResponse(fmt.Sprintf("data dengan id[%d] tidak ditemukan", id), fmt.Errorf("data not found"))
		ctx.AbortWithStatusJSON(http.StatusNotFound, errResponse)
		return
	}
	response := helper.BuildResponse(user)
	ctx.JSON(http.StatusOK, response)
	return
}

var Users []*model.User

func getNewID() int {
	if len(Users) == 0 {
		return 1
	}
	return Users[len(Users)-1].ID + 1
}
