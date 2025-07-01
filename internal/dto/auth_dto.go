package dto

type RegisterRequest struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,password,min=6"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	Email           string `json:"email" validate:"required,email"`
	OTPToken        string `json:"otp_token" validate:"required,len=4"`
	NewPassword     string `json:"new_password" validate:"required,min:6"`
	PasswordConfirm string `json:"password_confirm" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
