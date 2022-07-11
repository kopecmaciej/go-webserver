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
	Id             uint   `json:"id" gorm:"primaryKey"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"-"`
  CreatedAt      time.Time `json:"createdAt"`
}

func (u *NewUser) CreateUser() error {
	db := lib.GetDB()
	user := User{Username: u.Username, Email: u.Email, CreatedAt: time.Now()}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	user.HashedPassword = string(hashed)
	err := db.Create(&user).Error
	return err
}

func (u *User) GetUser(Id int) (User, error) {
	var user User
	db := lib.GetDB()
	if result := db.First(&user, Id); result.Error != nil {
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
