package domain

import (
	"github.com/billowdev/life-journal/pkg/uuidv7"
	"gorm.io/gorm"
)

type USER_STATUS string

const (
	USER_STATUS_ACTIVE   USER_STATUS = "active"
	USER_STATUS_INACTIVE USER_STATUS = "inactive"
	USER_STATUS_BANNED   USER_STATUS = "banned"
)

type User struct {
	BaseModel
	FirstName string `json:"first_name" validate:"required,max=255" gorm:"size:255"`
	LastName  string `json:"last_name" validate:"required,max=255"`
	Email    string      `json:"email" validate:"required,max=150" gorm:"size:50"`
	Username string      `json:"username" validate:"required,max=255" gorm:"size:255"`
	Password string      `json:"password"`
	Status   USER_STATUS `json:"status" validate:"required,max=50" gorm:"size:50"`
}

var TNUser = "users"

func (st *User) TableName() string {
	return TNUser
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID, err = uuidv7.GenerateUUIDv7(); err != nil {
		return err
	}
	return nil
}
