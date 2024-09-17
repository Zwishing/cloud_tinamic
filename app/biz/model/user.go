package model

type LoginRequest struct {
	UserAccount string `json:"userAccount" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Category    string `json:"category" validate:"required,oneof= username email phone"`
}

//type LoginResponse {
//
//}

type RegisterRequest struct {
	UserAccount string `json:"userAccount" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Category    string `json:"category" validate:"required,oneof= username email phone"`
}
