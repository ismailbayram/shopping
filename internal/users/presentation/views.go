package presentation

import (
	"github.com/gin-gonic/gin"
	"github.com/ismailbayram/shopping/internal/users/models"
	"net/http"
)

type UserService interface {
	GetByID(uint) (models.User, error)
	GetByToken(string) (models.User, error)
	Login(string, string) (string, error)
	Register(string, string, string, string) error
	Verify(string) error
	ChangePassword(models.User, string) error
}

type UserViews struct {
	Service UserService
}

func NewUserViews(service UserService) UserViews {
	return UserViews{
		Service: service,
	}
}

func (view *UserViews) Login(ctx *gin.Context) {
	loginDTO := LoginDTO{}
	if err := ctx.ShouldBindJSON(&loginDTO); err != nil {
		_ = ctx.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	token, err := view.Service.Login(loginDTO.Email, loginDTO.Password)
	if err != nil {
		_ = ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (view *UserViews) Register(ctx *gin.Context) {
	registerDTO := RegisterDTO{}
	if err := ctx.ShouldBindJSON(&registerDTO); err != nil {
		_ = ctx.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	if registerDTO.Password1 != registerDTO.Password2 {
		_ = ctx.Error(models.ErrorPasswordUnmatched).SetType(gin.ErrorTypePublic)
		return
	}

	err := view.Service.Register(registerDTO.Email, registerDTO.Password1, registerDTO.FirstName, registerDTO.LastName)
	if err != nil {
		_ = ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{})
}

func (view *UserViews) Verify(ctx *gin.Context) {
	token := ctx.Param("token")
	if err := view.Service.Verify(token); err != nil {
		_ = ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}

func (view *UserViews) ChangePassword(ctx *gin.Context) {
	userID := ctx.GetUint("user")
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	changePasswordDTO := ChangePasswordDTO{}
	if err := ctx.ShouldBindJSON(&changePasswordDTO); err != nil {
		_ = ctx.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	if changePasswordDTO.NewPassword1 != changePasswordDTO.NewPassword2 {
		_ = ctx.Error(models.ErrorPasswordUnmatched).SetType(gin.ErrorTypePublic)
		return
	}

	user, _ := view.Service.GetByID(userID)
	if err := user.CheckPassword(changePasswordDTO.Password); err != nil {
		_ = ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	err := view.Service.ChangePassword(user, changePasswordDTO.NewPassword1)
	if err != nil {
		_ = ctx.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
}
