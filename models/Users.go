package models

import (
	"errors"
	"html"
	"log"
	"strings"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	IdUser   uint64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Email    string  `json:"email" gorm:"column:email"`
	Password string  `json:"password" gorm:"column:password"`
	Name     string  `json:"name" gorm:"column:name"`
	Company  string  `json:"company" gorm:"column:company"`
	Sector   string  `json:"sector" gorm:"column:sector"`
	Wallet   float64 `json:"wallet" gorm:"column:wallet"`
}

type Recharge struct {
	IdUser uint64  `json:"id"`
	Value  float64 `json:"value"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave(db *gorm.DB) error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Company = html.EscapeString(strings.TrimSpace(u.Company))
	u.Sector = html.EscapeString(strings.TrimSpace(u.Sector))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint64) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}

	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint64) (*User, error) {

	// To hash the password
	err := u.BeforeSave(db)
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password": u.Password,
			"name":     u.Name,
			"email":    u.Email,
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) UpdateWallet(db *gorm.DB, uid uint64) (*User, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"wallet": u.Wallet,
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}

	err := db.Select("name", "wallet").Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
