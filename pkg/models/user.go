package models

import (
	"time"

	"go-web-server/lib"

	"golang.org/x/crypto/bcrypt"
)

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id             uint      `json:"id" gorm:"primaryKey"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	Role           string    `json:"role"`
	CreatedAt      time.Time `json:"createdAt"`
}

func (u *NewUser) CreateUser() (User, error) {
	db := lib.GetDB()
	user := User{Username: u.Username, Email: u.Email, Role: "User", CreatedAt: time.Now()}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	user.HashedPassword = string(hashed)
	if result := db.Create(&user); result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (u *User) GetUser() (User, error) {
	var user User
	db := lib.GetDB()
	if result := db.First(&user, u); result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (u *User) GetAllUsers() ([]User, error) {
	var users []User
	db := lib.GetDB()
	if result := db.Find(&users); result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func (u *User) DeleteUser(Id int) error {
	var user User
	db := lib.GetDB()
	err := db.Delete(&user, Id).Error
	return err
}
