package models

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

type UsersModel struct {
	GDB *gorm.DB
}

type User struct {
	gorm.Model
	Username string
	Password string
}

func NewUsersModel(gdb *gorm.DB) *UsersModel {
	return &UsersModel{
		GDB: gdb,
	}
}

func (um *UsersModel) List() (*[]User, error) {
	users := &[]User{}
	err := um.GDB.Find(users).Error
	if err != nil {
		return nil, fmt.Errorf("error fetching list of users: %s", strings.ToLower(err.Error()))
	}
	return users, nil
}

func (um *UsersModel) Get(id uint) (*User, error) {
	user := &User{}
	err := um.GDB.First(user, id).Error
	if err != nil {
		return nil, fmt.Errorf("error fetching single user: %s", strings.ToLower(err.Error()))
	}
	return user, nil
}

func (um *UsersModel) Create(user *User) (*User, error) {
	err := um.GDB.Create(user).Error
	if err != nil {
		return nil, fmt.Errorf("error creating user: %s", strings.ToLower(err.Error()))
	}
	return user, nil
}

func (um *UsersModel) Update(id uint, userupdates *User) (*User, error) {
	u := &User{}
	err := um.GDB.First(u, id).Error
	if err != nil {
		return nil, fmt.Errorf("user not found: %s", err.Error())
	}
	u.Username = userupdates.Username
	u.Password = userupdates.Password
	err = um.GDB.Save(u).Error
	if err != nil {
		return nil, fmt.Errorf("error updating user: %s", err.Error())
	}
	return u, nil
}

func (um *UsersModel) Delete(id uint) error {
	ou := &User{}
	err := um.GDB.First(ou, id).Error
	if err != nil {
		return fmt.Errorf("user not found: %s", err.Error())
	}
	err = um.GDB.Delete(ou).Error
	if err != nil {
		return fmt.Errorf("error deleting user: %s", err.Error())
	}
	return nil
}

//
