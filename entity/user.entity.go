package entity

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID     uuid.UUID `gorm:"NOT NULL;uniqueIndex" json:"id"`
	Name   string    `json:"name" gorm:"unique"`
	Email  string    `gorm:"NOT NULL;uniqueIndex" json:"email"`
	Flat   *Flat     `gorm:"foreignKey:FlatId"`
	FlatId uuid.UUID `json:"flat_id"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	err = u.Validate()
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	err = u.Validate()
	return
}

func (u *User) Validate() (err error) {
	if u.Name == "" {
		return errors.New("Name is required")
	}

	if u.Email == "" {
		return errors.New("Email is required")
	}
	if u.FlatId == uuid.Nil {
		return errors.New("Invalid Flat")
	}
	return
}

type UserRepository interface {
	Create(db *gorm.DB, user User) (err error)
	Update(db *gorm.DB, user User) (err error)
	Delete(db *gorm.DB, id string) (err error)
	Find(db *gorm.DB) (usr []User, err error)
	FindByID(db *gorm.DB, id string) (usr User, err error)
}

type UserUsecase interface {
	Upsert(usr User) (err error)
	Add(tx *gorm.DB, usr User) (err error)
	Edit(tx *gorm.DB, usr User) (err error)
	Remove(id string) (err error)
	GetAll() (result []User, err error)
}
