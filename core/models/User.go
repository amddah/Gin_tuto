package models

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
) 

type User struct {
	ID  string `gorm:"type:varchar(64);primaryKey" json:"id"`

	Name string `gorm:"type:varchar(255)" json:"name"`
	Email string `gorm:"type:varchar(255);unique" json:"email"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	uuidV7, err := uuid.NewV7()
	if err != nil {
		return err
	}
	u.ID = uuidV7.String()
	fmt.Println("UUID v7 en Base64:", u.ID)
	return err
}