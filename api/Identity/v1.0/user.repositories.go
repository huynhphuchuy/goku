package account

import (
	db "Gogin/internal/platform/database"
)

// Register Repository
func (repository UserRepository) Register(u User) (int, error) {
	user := new(User)
	createUser := db.Session.Create(&u).Scan(&user)
	return user.ID, createUser.Error
}

// Verify Repository
func (repository UserRepository) Verify(email string) error {
	db.Session.Model(&User{}).Where("email = ?", email).Update(User{Active: true})
	return nil
}

// ValidateCredentials Repository
func (repository UserRepository) ValidateCredentials(h UserSignin) bool {
	var count int
	db.Session.Model(&User{}).Where("username = ?", h.Username).Where("password = ?", h.Password).Count(&count)
	return (count != 0)
}

// GetUserByIdentity Repository
func (repository UserRepository) GetUserByIdentity(username string, email string) *User {
	var count int
	var user *User = new(User)
	db.Session.Model(&User{}).Where("username = ?", username).Or("email = ?", email).First(&user).Count(&count)
	if count > 0 {
		return user
	}
	return nil
}
