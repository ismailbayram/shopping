package presentation

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ismailbayram/shopping/internal/users/domain"
	"github.com/ismailbayram/shopping/test/mocks"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserViews_Login_Failed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedUS := mocks.NewUserService(t)
	views := NewUserViews(mockedUS)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/users/login/", nil)
	views.Login(ctx)
	assert.Equal(t, "EOF", ctx.Errors[0].Error())

	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	reqBody, _ := json.Marshal(map[string]string{
		"name": "new name",
	})
	ctx.Request = httptest.NewRequest("POST", "/users/login/", bytes.NewReader(reqBody))
	views.Login(ctx)
	assert.Equal(t, 1, len(ctx.Errors))
	errs := ctx.Errors[0].Err.(validator.ValidationErrors)
	assert.Equal(t, 2, len(errs))
	assert.Equal(t, "email", errs[0].Tag())
	assert.Equal(t, "LoginDTO.Email", errs[0].Namespace())
	assert.Equal(t, "required", errs[1].Tag())
	assert.Equal(t, "LoginDTO.Password", errs[1].Namespace())

	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	reqBody, _ = json.Marshal(map[string]string{
		"email":    "nonexisted@gmail.com",
		"password": "123456",
	})
	ctx.Request = httptest.NewRequest("POST", "/users/login/", bytes.NewReader(reqBody))
	mockedUS.On("Login", "nonexisted@gmail.com", "123456").Return("", domain.ErrorUserNotFound).Once()
	views.Login(ctx)
	assert.Equal(t, domain.ErrorUserNotFound.Error(), ctx.Errors[0].Error())
}

func TestUserViews_Login_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedUS := mocks.NewUserService(t)
	views := NewUserViews(mockedUS)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	reqBody, _ := json.Marshal(map[string]string{
		"email":    "nonexisted@gmail.com",
		"password": "123456",
	})
	ctx.Request = httptest.NewRequest("POST", "/users/login/", bytes.NewReader(reqBody))
	mockedUS.On("Login", "nonexisted@gmail.com", "123456").Return("token", nil).Once()
	views.Login(ctx)
	assert.Empty(t, ctx.Errors)
	assert.Equal(t, http.StatusOK, w.Code)
	resp, _ := io.ReadAll(w.Body)
	var response map[string]string
	_ = json.Unmarshal(resp, &response)
	assert.Equal(t, "token", response["token"])
}

func TestUserViews_Register_Failed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedUS := mocks.NewUserService(t)
	views := NewUserViews(mockedUS)

	// with empty body
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = httptest.NewRequest("POST", "/users/register/", nil)
	views.Register(ctx)
	assert.Equal(t, "EOF", ctx.Errors[0].Error())

	// with wrong data
	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	reqBody, _ := json.Marshal(map[string]string{
		"first_name": "ismail",
		"last_name":  "bayrambayrambayrambayrambayrambayrambayrambayrambayrambayrambayram",
		"email":      "iso.bayram@gmail.com",
		"password1":  "diff",
		"password2":  "passpasspasspasspasspasspass",
	})
	ctx.Request = httptest.NewRequest("POST", "/users/register/", bytes.NewReader(reqBody))
	views.Register(ctx)
	assert.Equal(t, 1, len(ctx.Errors))
	errs := ctx.Errors[0].Err.(validator.ValidationErrors)
	assert.Equal(t, 3, len(errs))
	assert.Equal(t, "max", errs[0].Tag())
	assert.Equal(t, "RegisterDTO.LastName", errs[0].Namespace())
	assert.Equal(t, "min", errs[1].Tag())
	assert.Equal(t, "RegisterDTO.Password1", errs[1].Namespace())
	assert.Equal(t, "max", errs[2].Tag())
	assert.Equal(t, "RegisterDTO.Password2", errs[2].Namespace())

	// with unmatched passwords
	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	reqBody, _ = json.Marshal(map[string]string{
		"first_name": "ismail",
		"last_name":  "bayram",
		"email":      "iso.bayram@gmail.com",
		"password1":  "123456",
		"password2":  "654321",
	})
	ctx.Request = httptest.NewRequest("POST", "/users/register/", bytes.NewReader(reqBody))
	views.Register(ctx)
	assert.Equal(t, domain.ErrorPasswordUnmatched.Error(), ctx.Errors[0].Error())

	// with existed user
	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	reqBody, _ = json.Marshal(map[string]string{
		"first_name": "ismail",
		"last_name":  "bayram",
		"email":      "iso@iso.com",
		"password1":  "123456",
		"password2":  "123456",
	})
	ctx.Request = httptest.NewRequest("POST", "/users/register/", bytes.NewReader(reqBody))
	mockedUS.On("Register", "iso@iso.com", "123456", "ismail", "bayram").Return(domain.ErrorUserAlreadyExists).Once()
	views.Register(ctx)
	assert.Equal(t, domain.ErrorUserAlreadyExists.Error(), ctx.Errors[0].Error())
}

func TestUserViews_Register_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedUS := mocks.NewUserService(t)
	views := NewUserViews(mockedUS)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	reqBody, _ := json.Marshal(map[string]string{
		"first_name": "ismail",
		"last_name":  "bayram",
		"email":      "iso@iso.com",
		"password1":  "123456",
		"password2":  "123456",
	})
	ctx.Request = httptest.NewRequest("POST", "/users/register/", bytes.NewReader(reqBody))
	mockedUS.On("Register", "iso@iso.com", "123456", "ismail", "bayram").Return(nil).Once()
	views.Register(ctx)
	assert.Empty(t, ctx.Errors)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUserViews_Verify_Failed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedUS := mocks.NewUserService(t)
	views := NewUserViews(mockedUS)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = []gin.Param{
		{Key: "token", Value: "invalid"},
	}
	mockedUS.On("Verify", "invalid").Return(domain.ErrorUserNotFound).Once()
	views.Verify(ctx)
	assert.Equal(t, domain.ErrorUserNotFound.Error(), ctx.Errors[0].Error())
}

func TestUserViews_Verify_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedUS := mocks.NewUserService(t)
	views := NewUserViews(mockedUS)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Params = []gin.Param{
		{Key: "token", Value: "valid"},
	}
	mockedUS.On("Verify", "valid").Return(nil).Once()
	views.Verify(ctx)
	assert.Empty(t, ctx.Errors)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserViews_ChangePassword_Failed(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedUS := mocks.NewUserService(t)
	views := NewUserViews(mockedUS)

	// without token
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	reqBody, _ := json.Marshal(map[string]string{
		"password": "oldpassword",
	})
	ctx.Request = httptest.NewRequest("POST", "/users/change-password/", bytes.NewReader(reqBody))
	views.ChangePassword(ctx)
	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// with empty body
	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	ctx.Set("user", uint(1))
	ctx.Request = httptest.NewRequest("POST", "/users/change-password/", nil)
	views.ChangePassword(ctx)
	assert.Equal(t, 1, len(ctx.Errors))
	assert.Equal(t, "EOF", ctx.Errors[0].Error())

	// with wrong body
	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	ctx.Set("user", uint(1))
	reqBody, _ = json.Marshal(map[string]string{
		"password": "oldpassword",
	})
	ctx.Request = httptest.NewRequest("POST", "/users/change-password/", bytes.NewReader(reqBody))
	views.ChangePassword(ctx)
	assert.Equal(t, 1, len(ctx.Errors))
	errs := ctx.Errors[0].Err.(validator.ValidationErrors)
	assert.Equal(t, 2, len(errs))
	assert.Equal(t, "min", errs[0].Tag())
	assert.Equal(t, "ChangePasswordDTO.NewPassword1", errs[0].Namespace())
	assert.Equal(t, "min", errs[1].Tag())
	assert.Equal(t, "ChangePasswordDTO.NewPassword2", errs[1].Namespace())

	// with unmatched passwords
	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	ctx.Set("user", uint(1))
	reqBody, _ = json.Marshal(map[string]string{
		"password":      "oldpassword",
		"new_password1": "p1123123",
		"new_password2": "asdasdasdasd",
	})
	ctx.Request = httptest.NewRequest("POST", "/users/change-password/", bytes.NewReader(reqBody))
	views.ChangePassword(ctx)
	assert.Equal(t, 1, len(ctx.Errors))
	assert.Equal(t, domain.ErrorPasswordUnmatched.Error(), ctx.Errors[0].Error())

	// with wrong old password
	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	ctx.Set("user", uint(1))
	reqBody, _ = json.Marshal(map[string]string{
		"password":      "wrong",
		"new_password1": "newpassword",
		"new_password2": "newpassword",
	})
	user := domain.User{ID: 1}
	user.SetPassword("oldpassword")
	mockedUS.On("GetByID", uint(1)).Return(user, nil).Once()
	ctx.Request = httptest.NewRequest("POST", "/users/change-password/", bytes.NewReader(reqBody))
	views.ChangePassword(ctx)
	assert.Equal(t, 1, len(ctx.Errors))
	assert.Equal(t, domain.ErrorWrongPassword.Error(), ctx.Errors[0].Error())

	// with service error.
	w = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(w)
	ctx.Set("user", uint(1))
	reqBody, _ = json.Marshal(map[string]string{
		"password":      "oldpassword",
		"new_password1": "newpassword",
		"new_password2": "newpassword",
	})
	mockedUS.On("GetByID", uint(1)).Return(user, nil).Once()
	mockedUS.On("ChangePassword", user, "newpassword").Return(domain.ErrorGeneral).Once()
	ctx.Request = httptest.NewRequest("POST", "/users/change-password/", bytes.NewReader(reqBody))
	views.ChangePassword(ctx)
	assert.Equal(t, 1, len(ctx.Errors))
	assert.Equal(t, domain.ErrorGeneral.Error(), ctx.Errors[0].Error())
}

func TestUserViews_ChangePassword_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockedUS := mocks.NewUserService(t)
	views := NewUserViews(mockedUS)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set("user", uint(1))
	reqBody, _ := json.Marshal(map[string]string{
		"password":      "oldpassword",
		"new_password1": "newpassword",
		"new_password2": "newpassword",
	})
	user := domain.User{ID: 1}
	user.SetPassword("oldpassword")
	mockedUS.On("GetByID", uint(1)).Return(user, nil).Once()
	mockedUS.On("ChangePassword", user, "newpassword").Return(nil).Once()
	ctx.Request = httptest.NewRequest("POST", "/users/change-password/", bytes.NewReader(reqBody))
	views.ChangePassword(ctx)
	assert.Empty(t, ctx.Errors)
	assert.Equal(t, http.StatusOK, w.Code)
}
