package flat

import (
	"github.com/harisosa/flat-system-backend/entity"
	"gorm.io/gorm"
)

type flatRepository struct{}

//NewUserSettingRepository create new instance for categry repository
func NewFlatRepository() entity.FlatRepository {
	return &flatRepository{}
}

func (ur *flatRepository) Find(db *gorm.DB) (usr []entity.Flat, err error) {
	err = db.Order("name").Preload("Neighborhood").Find(&usr).Error
	if err != nil {
		return
	}
	return
}

func (ur *flatRepository) FindByID(db *gorm.DB, id string) (usr entity.Flat, err error) {
	err = db.Where("id = ?", id).First(&usr).Error
	if err != nil {
		return
	}
	return
}

func (ur *flatRepository) Create(db *gorm.DB, user entity.Flat) (err error) {
	err = db.Create(&user).Error
	return
}

func (ur *flatRepository) Update(db *gorm.DB, flat entity.Flat) (err error) {
	err = db.Where("id = ?", flat.ID).Updates(&flat).Error
	return
}

func (ur *flatRepository) Delete(db *gorm.DB, id string) (err error) {
	usr, err := ur.FindByID(db, id)
	if err != nil {
		return
	}
	err = db.Unscoped().Delete(&usr).Error
	return
}
