package neighborhood

import (
	"github.com/harisosa/flat-system-backend/entity"
	"gorm.io/gorm"
)

type neighborhoodUsecase struct {
	db   *gorm.DB
	repo entity.NeighborhoodRepository
}

//NewCategoryUsecase create new instance for categry usecase
func NewNeighborhoodUsecase(
	db *gorm.DB,
	rp entity.NeighborhoodRepository) entity.NeighborhoodUsecase {
	return &neighborhoodUsecase{db, rp}
}
func (uuc *neighborhoodUsecase) Add(tx *gorm.DB, flat entity.Neighborhood) (err error) {
	return uuc.repo.Create(tx, flat)
}

func (uuc *neighborhoodUsecase) Remove(ids []string) (err error) {
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

func (uuc *neighborhoodUsecase) GetAll() (result []entity.Neighborhood, err error) {
	return uuc.repo.Find(uuc.db)
}

func (uuc *neighborhoodUsecase) GetByID(id string) (result entity.Neighborhood, err error) {
	result, err = uuc.repo.FindByID(uuc.db, id)
	return result, err
}

func (uuc *neighborhoodUsecase) GetByLocation(location int64) (result entity.Neighborhood, err error) {
	return uuc.repo.FindByRange(uuc.db, location)
}
