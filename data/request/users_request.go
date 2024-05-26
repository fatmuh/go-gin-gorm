package request

type CreateUserRequest struct {
	Username string `validate:"required,min=1,max=255" json:"username"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8,max=255" json:"password"`
}

type UpdateUsersRequest struct {
	Id       int    `validate:"required,min=1" json:"id"`
	Username string `validate:"required,min=1,max=255" json:"username"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8,max=255" json:"password"`
}

type LoginRequest struct {
	Username string `validate:"required,min=1,max=255" json:"username"`
	Password string `validate:"required,min=8,max=255" json:"password"`
}
