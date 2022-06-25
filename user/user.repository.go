package user

import (
	"github.com/harisosa/flat-system-backend/entity"
	"gorm.io/gorm"
)

type userRepository struct{}

//NewUserSettingRepository create new instance for categry repository
func NewUserRepository() entity.UserRepository {
	return &userRepository{}
}

func (ur *userRepository) Find(db *gorm.DB) (usr []entity.User, err error) {
	err = db.Order("name").Preload("Flat").Find(&usr).Error
	if err != nil {
		return
	}
	return
}

func (ur *userRepository) FindByID(db *gorm.DB, id string) (usr entity.User, err error) {
	err = db.Where("id = ?", id).First(&usr).Error
	if err != nil {
		return
	}
	return
}

func (ur *userRepository) Create(db *gorm.DB, user entity.User) (err error) {
	err = db.Create(&user).Error
	return
}

func (ur *userRepository) Update(db *gorm.DB, user entity.User) (err error) {
	err = db.Where("id = ?", user.ID).Updates(&user).Error
	return
}

func (ur *userRepository) Delete(db *gorm.DB, id string) (err error) {
	usr, err := ur.FindByID(db, id)
	if err != nil {
		return
	}
	err = db.Unscoped().Delete(&usr).Error
	return
}
