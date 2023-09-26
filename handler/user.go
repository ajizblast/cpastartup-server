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
	//tangkap input dari user
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

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		//untuk menampilkan pesan error lebih rapi
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "token")

	//tambahkan helper response
	response := helper.APIResponse("Success login", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// proses cek email
	// ada input email dari user
	// input email di mapping ke struct input - di handler
	// struct input di passing ke service
	// service akan manggil repository - email sudah ada atau belum
	// repository - db
	// end cek email

	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		//untuk menampilkan pesan error lebih rapi
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email Check Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return //kalau ada return, berarti akan di stop saat ada error
	}

	// check dari service
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.APIResponse("Email Check Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return //kalau ada return, berarti akan di stop saat ada error
	}

	// pakai gin.H utk beri response data
	data := gin.H{
		"is_available": isEmailAvailable,
	}

	// buat variable baru utk cek dan beri response json
	var metaMessage string

	if isEmailAvailable {
		metaMessage = "Email is available"
	} else {
		metaMessage = "Email has been registered"
	}

	// ini kalau balasan berhasil / tidak ada email kembar
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
