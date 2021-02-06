package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

const userPwPepper = "secret-random-string"

// All the databse information goes here, Model in mvc
type Userinfo struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"`
	PasswrodHash string `gorm:"not null"`
}

type UserInfoService struct {
	db *gorm.DB
}

/*To intialise a new database instance.
Any controller that uses database has to create an instance of databse by calling this function */

func NewUserInfoService(connectionInfo string) (*UserInfoService, error) {
	db, err := gorm.Open("postgres", connectionInfo)

	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return &UserInfoService{
		db: db,
	}, nil
}

func (us *UserInfoService) Create(user *Userinfo) error {
	pwBytes := []byte(user.Password + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswrodHash = string(hashedBytes)
	user.Password = ""
	return us.db.Create(user).Error
}

func (us *UserInfoService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&Userinfo{}).Error; err != nil {
		return err
	}
	return nil
}
