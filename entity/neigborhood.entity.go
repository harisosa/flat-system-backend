package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Neighborhood struct {
	ID        uuid.UUID `json:"id" gorm:"unique;NOT NULL"`
	Name      string    `json:"name" gorm:"unique"`
	RangeFrom int       `json:"range_from"`
	RangeTo   int       `json:"range_to"`
}

func (n *Neighborhood) BeforeCreate(tx *gorm.DB) (err error) {
	if n.ID == uuid.Nil {
		n.ID = uuid.New()
	}

	return
}

type NeighborhoodRepository interface {
	Create(db *gorm.DB, neighborhood Neighborhood) (err error)
	Delete(db *gorm.DB, id string) (err error)
	Find(db *gorm.DB) (neighborhood []Neighborhood, err error)
	FindByID(db *gorm.DB, id string) (neighborhood Neighborhood, err error)
	FindByRange(db *gorm.DB, location int64) (neighborhood Neighborhood, err error)
}

type NeighborhoodUsecase interface {
	Add(tx *gorm.DB, usr Neighborhood) (err error)
	Remove(ids []string) (err error)
	GetAll() (result []Neighborhood, err error)
	GetByLocation(location int64) (result Neighborhood, err error)
}
