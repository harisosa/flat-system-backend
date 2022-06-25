package user

import (
	"github.com/google/uuid"
	"github.com/harisosa/flat-system-backend/entity"
	"github.com/harisosa/flat-system-backend/helper"
	"gorm.io/gorm"
)

type userUsecase struct {
	db   *gorm.DB
	repo entity.UserRepository
	help helper.Helper
}

//NewCategoryUsecase create new instance for categry usecase
func NewUserUsecase(
	db *gorm.DB,
	rp entity.UserRepository,
	help helper.Helper) entity.UserUsecase {
	return &userUsecase{db, rp, help}
}
func (uuc *userUsecase) Upsert(usr entity.User) (err error) {
	err = uuc.db.Transaction(func(tx *gorm.DB) (err error) {

		if usr.ID != uuid.Nil {
			err = uuc.Edit(tx, usr)
		} else {
			err = uuc.Add(tx, usr)
		}
		return
	})
	return
}

func (uuc *userUsecase) Add(tx *gorm.DB, usr entity.User) (err error) {
	return uuc.repo.Create(tx, usr)
}

func (uuc *userUsecase) Edit(tx *gorm.DB, usr entity.User) (err error) {
	return uuc.repo.Update(tx, usr)
}

func (uuc *userUsecase) Remove(id string) (err error) {
	return uuc.repo.Delete(uuc.db, id)
}

func (uuc *userUsecase) GetAll() (result []entity.User, err error) {
	return uuc.repo.Find(uuc.db)
}

func (uuc *userUsecase) GetByID(id string) (result entity.User, err error) {
	result, err = uuc.repo.FindByID(uuc.db, id)
	return result, err
}
