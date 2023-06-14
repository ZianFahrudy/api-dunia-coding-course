package handler

import (
	"api-dunia-coding/helper"
	"api-dunia-coding/member"
	"api-dunia-coding/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type memberHandler struct {
	memberService member.Service
	jwtService    service.JwtService
}

func NewMemberHandler(memberService member.Service, jwtService service.JwtService) *memberHandler {
	return &memberHandler{memberService, jwtService}
}

func (h *memberHandler) RegisterMember(c *gin.Context) {

	var input member.RegisterMemberInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errMessage := gin.H{
			"errors": errors,
		}

		response := helper.APIResponse("Register Member Failed", http.StatusUnprocessableEntity, errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if len(input.Name) > 45 {
		response := helper.APIResponse("Nama maksimal 45 karakter", http.StatusUnprocessableEntity, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newMember, err := h.memberService.RegisterMember(input)
	if err != nil {

		response := helper.APIResponse("Register Member Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.jwtService.GenerateToken(newMember)

	if err != nil {
		response := helper.APIResponse("Register Member Failed", http.StatusBadRequest, nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := member.FormatMember(newMember, token)

	response := helper.APIResponse("Register Member Success", http.StatusOK, formatter)
	c.JSON(http.StatusOK, response)
}

func (h *memberHandler) Login(c *gin.Context) {

	var input member.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {

		errMessage := gin.H{
			"message": "Email dan Password tidak boleh kosong",
		}

		c.JSON(http.StatusUnprocessableEntity, errMessage)
		return
	}

	_, err = h.memberService.CheckEmailOrPasswordValid(input)

	if err != nil {
		errMessage := gin.H{
			"message": "Email atau Password yang dimasukkan salah ",
		}

		c.JSON(http.StatusUnprocessableEntity, errMessage)
		return
	}

	loggedInUser, err := h.memberService.Login(input)
	if err != nil {
		errMessage := gin.H{
			"message": "Login Gagal",
		}

		c.JSON(http.StatusBadRequest, errMessage)
		return
	}

	token, err := h.jwtService.GenerateToken(loggedInUser)
	if err != nil {
		errMessage := gin.H{
			"message": "Login Gagal",
		}

		c.JSON(http.StatusBadRequest, errMessage)
		return
	}

	formatter := member.FormatMemberLogin("Login Berhasil", token)

	c.JSON(http.StatusOK, formatter)

}

func (h *memberHandler) CheckEmailAvaibility(c *gin.Context) {
	var input member.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{
			"errors": errors,
		}
		response := helper.APIResponse("Email Check Failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return

	}

	isEmailAvailable, err := h.memberService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{
			"errors": "Server Error",
		}
		response := helper.APIResponse("Email Check Failed", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	var metaMessage string

	if isEmailAvailable {
		metaMessage = "Email belum terdaftar"
	} else {
		metaMessage = "Email sudah terdaftar"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, data)
	c.JSON(http.StatusOK, response)

}

func (h *memberHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(member.Member)
	memberID := currentUser.ID

	path := "images/" + file.Filename
	path = fmt.Sprintf("images/%d-%s", memberID, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.memberService.SaveAvatar(memberID, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_uploaded": true,
	}
	response := helper.APIResponse("Success to upload avatar image", http.StatusOK, data)

	c.JSON(http.StatusOK, response)
	return
}
