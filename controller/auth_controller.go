package controller

import (
	"api-dunia-coding/common"
	"api-dunia-coding/common/exception"
	"api-dunia-coding/config"
	"api-dunia-coding/data/model"
	"api-dunia-coding/domain/entity"
	"api-dunia-coding/domain/repository"
	"api-dunia-coding/middleware"
	"api-dunia-coding/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService service.AuthService
	repository.AuthRepository
	config.Config
}

func NewAuthController(authService *service.AuthService, authRepository repository.AuthRepository, config config.Config) AuthController {
	return AuthController{AuthService: *authService, AuthRepository: authRepository, Config: config}
}

func (controller *AuthController) Route(app *gin.Engine) {
	api := app.Group("/api/v1")

	api.POST("/login", controller.Login)
	api.POST("/register", controller.RegisterMember)
	api.POST("/avatar", middleware.AuthenticateJWT(controller.AuthRepository, controller.Config), controller.UploadAvatar)

}

func (controller *AuthController) RegisterMember(c *gin.Context) {

	var input model.RegisterMemberBody

	errReq := c.ShouldBindJSON(&input)
	exception.PanicIfNeeded(errReq)
	if errReq != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Register Gagal",
			Data:    nil,
		})
		return
	}

	// cek email already exist
	emailIsExist, _ := controller.AuthService.CheckEmailAvailable(c.Copy(), input.Email)

	if emailIsExist {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Email sudah terdaftar silahkan gunakan email yang lain",
			Data:    nil,
		})
		return
	}

	// cek jika input nama lebih dari 45 karakter
	if len(input.Name) > 45 {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Nama maksimal 45 karakter",
			Data:    nil,
		})
		return
	}

	// create new member
	newMember := controller.AuthService.RegisterMember(c.Copy(), input)

	// generate token
	token := common.GenerateToken(newMember.Name, newMember.ID, controller.Config)

	// mapping data register Success
	resultWithToken := map[string]interface{}{
		"username": newMember.Name,
		"token":    token,
	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Register Member Success",
		Data:    resultWithToken,
	})
	return

}

func (controller *AuthController) Login(c *gin.Context) {
	var request model.LoginBody
	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "Login Gagal",
			Data:    nil,
		})
		return
	}

	_, err = controller.AuthService.CheckEmailOrPasswordValid(c.Copy(), request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Email atau password salah",
			Data:    nil,
		})
		return
	}
	response, errs := controller.AuthService.Login(c.Copy(), request)

	if errs != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Login Gagal",
			Data:    nil,
		})
		return
	}
	tokenJwtResult := common.GenerateToken(response.Name, response.ID, controller.Config)

	resultWithToken := map[string]interface{}{
		"member_id":    response.ID,
		"member_name":  response.Name,
		"member_email": response.Email,
		"token":        tokenJwtResult,
	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Login Berhasil",
		Data:    resultWithToken,
	})

}

func (controller *AuthController) UploadAvatar(c *gin.Context) {
	currentUser := c.MustGet(controller.Get("JWT_CURRENT_USER")).(entity.Member)

	file, err := c.FormFile("avatar")

	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to upload avatar image",
			Data: gin.H{
				"is_uploaded": false,
			},
		})
		return
	}

	path := "storage/" + file.Filename
	path = fmt.Sprintf("storage/%d-%s", currentUser.ID, file.Filename)

	err = c.SaveUploadedFile(file, path)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to upload avatar image",
			Data: gin.H{
				"is_uploaded": false,
			},
		})
		return
	}

	_, err = controller.AuthService.SaveAvatar(c.Copy(), currentUser.ID, fmt.Sprintf("http:localhost:8081/%s", path))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GeneralResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to upload avatar image",
			Data: gin.H{
				"is_uploaded": false,
			},
		})
		return
	}

	c.JSON(http.StatusOK, model.GeneralResponse{
		Code:    http.StatusOK,
		Message: "Success to upload avatar image",
		Data: gin.H{
			"is_uploaded": true,
		},
	})
	return
}
