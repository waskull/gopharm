package core

type LoginParams struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserCreateParams struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"optional"`
	Dni      string `json:"dni" binding:"required"`
}

type EmployeeCreateParams struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Dni      string `json:"dni" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Phone    string `json:"phone" binding:"optional"`
}

type UserGetParams struct {
	ID interface{} `json:"id" binding:"required"`
}

type UserUpdateParams struct {
	Fullname string `json:"fullname" binding:"required"`
	Password string `json:"password" binding:"required"`
	Dni      string `json:"dni" binding:"required"`
	Phone    string `json:"phone" binding:"optional"`
}

type UserUpdateRoleParams struct {
	Role string `json:"role" binding:"required"`
}
