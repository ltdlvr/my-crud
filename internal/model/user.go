package model

type User struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
}

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=50"`
	Age  int    `json:"age" validate:"required,gt=0,lt=120"`
}

type UpdateUserRequest struct {
	Age int `json:"age" validate:"required,gt=0,lt=120"`
}

type UserResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}
