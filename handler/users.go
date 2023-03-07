package handler

import (
	"backend_startup/helper"
	"backend_startup/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct
type userHandler struct {
	userServices users.Services
}

func NewUserHandler(userServices users.Services) *userHandler {
	return &userHandler{userServices}
}

// handler
func (h *userHandler) RegisterUser(c *gin.Context) {
	// get data input from user

	var input users.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// format error input
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUsers, err := h.userServices.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
	}

	// token, genrate token
	formatter := users.FormatUser(newUsers, "Token")

	// call halper
	jsonResponse := helper.APIResponse("Account has been register", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, jsonResponse)
}
