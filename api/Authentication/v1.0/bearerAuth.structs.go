package authentication

// BearerAuthController Struct
type BearerAuthController struct {
}

// BearerAuthService Struct
type BearerAuthService struct {
}

// BearerAuthLogin Struct
type BearerAuthLogin struct {
	Username string `json:"username" binding:"required" valid:"required"`
	Password string `json:"password,omitempty" binding:"required" valid:"stringlength(5|9),required"`
}
