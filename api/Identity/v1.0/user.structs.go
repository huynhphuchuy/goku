package account

import (
	"time"
)

// UserController Struct
type UserController struct {
}

// UserService Struct
type UserService struct {
}

// UserRepository Struct
type UserRepository struct {
}

// User Struct
type User struct {
	ID        int        `json:"-"`
	Name      string     `json:"fullname" binding:"required" valid:"required"`
	Username  string     `json:"username" binding:"required" valid:"required"`
	Password  string     `json:"password,omitempty" binding:"required" valid:"stringlength(5|9),required"`
	Email     string     `json:"email" binding:"required" valid:"email,required"`
	DOB       string     `json:"dob,omitempty"`
	Gender    string     `json:"gender,omitempty"`
	Avatar    string     `json:"avatar,omitempty"`
	Payload   string     `json:"payload,omitempty"`
	Active    bool       `json:"-"`
	CreatedAt *time.Time `json:""`
	UpdatedAt *time.Time `json:""`
}

// UserSignin Struct
type UserSignin struct {
	Username string `json:"username" binding:"required" valid:"required"`
	Password string `json:"password,omitempty" binding:"required" valid:"stringlength(5|9),required"`
}
