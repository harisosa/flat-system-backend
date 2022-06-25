package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Flat struct {
	ID             uuid.UUID     `json:"id" gorm:"unique;NOT NULL"`
	Name           string        `json:"name"`
	Location       int64         `gorm:"NOT NULL" json:"location"`
	Neighborhood   *Neighborhood `gorm:"foreignKey:NeighborhoodId"`
	NeighborhoodId uuid.UUID
}

func (f *Flat) BeforeCreate(tx *gorm.DB) (err error) {
	if f.ID == uuid.Nil {
		f.ID = uuid.New()
	}

	return
}

type FlatRepository interface {
	Create(db *gorm.DB, user Flat) (err error)
	Update(db *gorm.DB, user Flat) (err error)
	Delete(db *gorm.DB, id string) (err error)
	Find(db *gorm.DB) (usr []Flat, err error)
	FindByID(db *gorm.DB, id string) (usr Flat, err error)
}

type FlateUsecase interface {
	Upsert(usr Flat) (err error)
	Add(tx *gorm.DB, usr Flat) (err error)
	Edit(tx *gorm.DB, usr Flat) (err error)
	Remove(ids []string) (err error)
	GetAll() (result []Flat, err error)
}
