package handler

import (
	"cpastartup/helper"
	"cpastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//tangkap inpit dari user
	//map input dari user ke struct ke RegisterUserInput
	//struct di aas kita passing sebagai parameter service

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {

		errors := helper.FormatValidationError(err)
		//untuk menampilkan pesan error lebih rapi
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "token")

	response = helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	//step step
	//user memasukkan input (email & password)
	//input ditangkap oleh handler
	//mapping dari input user ke input struct
	//input struct passing ke service
	//di service mencari dgn bantuan repository user dengan email X
	//mencocokkan dengan password
}
