package neighborhood

import (
	"github.com/harisosa/flat-system-backend/entity"
	"gorm.io/gorm"
)

type neighborhoodRepository struct{}

//NewUserSettingRepository create new instance for categry repository
func NewNeighborhoodRepository() entity.NeighborhoodRepository {
	return &neighborhoodRepository{}
}

func (ur *neighborhoodRepository) Find(db *gorm.DB) (nbr []entity.Neighborhood, err error) {
	err = db.Order("name").Find(&nbr).Error
	if err != nil {
		return
	}
	return
}

func (ur *neighborhoodRepository) FindByID(db *gorm.DB, id string) (nbr entity.Neighborhood, err error) {
	err = db.Where("id = ?", id).First(&nbr).Error
	if err != nil {
		return
	}
	return
}

func (ur *neighborhoodRepository) FindByRange(db *gorm.DB, location int64) (nbr entity.Neighborhood, err error) {
	err = db.Where("range_from > ?  AND ? < range_to", location, location).First(&nbr).Error
	if err != nil {
		return
	}
	return
}

func (ur *neighborhoodRepository) Create(db *gorm.DB, user entity.Neighborhood) (err error) {
	err = db.Create(&user).Error
	return
}

func (ur *neighborhoodRepository) Delete(db *gorm.DB, id string) (err error) {
	nbr, err := ur.FindByID(db, id)
	if err != nil {
		return
	}
	err = db.Unscoped().Delete(&nbr).Error
	return
}
