package model

type User struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UpdateUserRequest struct {
	Age int `json:"age"`
}

type UserResponse struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}
