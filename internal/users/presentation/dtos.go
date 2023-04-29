package presentation

type LoginDTO struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required"`
}

type RegisterDTO struct {
	FirstName string `json:"first_name" binding:"max=64"`
	LastName  string `json:"last_name" binding:"max=64"`
	Email     string `json:"email" binding:"email"`
	Password1 string `json:"password1" binding:"min=6,max=16"`
	Password2 string `json:"password2" binding:"min=6,max=16"`
}

type ChangePasswordDTO struct {
	Password     string `json:"password" binding:"required"`
	NewPassword1 string `json:"new_password1" binding:"min=6,max=16"`
	NewPassword2 string `json:"new_password2" binding:"min=6,max=16"`
}
