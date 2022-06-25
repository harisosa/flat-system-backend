package flat

import (
	"errors"

	"github.com/google/uuid"
	"github.com/harisosa/flat-system-backend/entity"
	"gorm.io/gorm"
)

type flatUsecase struct {
	db    *gorm.DB
	repo  entity.FlatRepository
	nrepo entity.NeighborhoodUsecase
}

//NewCategoryUsecase create new instance for categry usecase
func NewFlatUsecase(
	db *gorm.DB,
	rp entity.FlatRepository,
	nrp entity.NeighborhoodUsecase) entity.FlateUsecase {
	return &flatUsecase{db, rp, nrp}
}
func (fuc *flatUsecase) Upsert(flat entity.Flat) (err error) {
	err = fuc.db.Transaction(func(tx *gorm.DB) (err error) {

		nbr, err := fuc.nrepo.GetByLocation(flat.Location)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			flat.Neighborhood = &nbr
			flat.NeighborhoodId = nbr.ID
		}

		if flat.ID != uuid.Nil {
			err = fuc.Edit(tx, flat)
		} else {
			err = fuc.Add(tx, flat)
		}
		return
	})
	return
}

func (uuc *flatUsecase) Add(tx *gorm.DB, flat entity.Flat) (err error) {

	return uuc.repo.Create(tx, flat)
}

func (uuc *flatUsecase) Edit(tx *gorm.DB, flat entity.Flat) (err error) {
	return uuc.repo.Update(tx, flat)
}

func (uuc *flatUsecase) Remove(ids []string) (err error) {
	tx := uuc.db.Begin()
	for _, id := range ids {

		err = uuc.repo.Delete(tx, id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return
}

func (uuc *flatUsecase) GetAll() (result []entity.Flat, err error) {
	return uuc.repo.Find(uuc.db)
}

func (uuc *flatUsecase) GetByID(id string) (result entity.Flat, err error) {
	result, err = uuc.repo.FindByID(uuc.db, id)
	return result, err
}
