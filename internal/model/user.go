package model

type Login struct {
	MatNum   string `json:"matric_number" validate:"required"`
	Password string `json:"password" validate:"required, min=8"`
}

type LoginResponse struct {
	User  ListCoursesForm `json:"matric_number" `
	Token string          `json:"token" `
}
