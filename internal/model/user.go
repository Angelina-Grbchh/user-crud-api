package model

type User struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int32  `json:"age"`
}
