package seeder

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/harisosa/flat-system-backend/entity"
	"gorm.io/gorm"
)

//MongoSeeder instancce for database seeder
type PostgreSeeder interface {
	SeedFile()
	DatabaseMigration()
}

type postgreSeeder struct {
	store *gorm.DB
	l     *log.Logger
}

// NewMongoSeeder will create new an mongo seeder object representation of domain.UserUsecase interface
func NewPostgresSeeder(db *gorm.DB) PostgreSeeder {
	return &postgreSeeder{
		store: db,
		l:     log.New(os.Stdout, "Database Seeder ", log.LstdFlags),
	}
}

func (pg *postgreSeeder) DatabaseMigration() {
	pg.store.AutoMigrate()
	var tables []string
	pg.store.Table("information_schema.tables").Where("table_schema = ?", "public").Pluck("table_name", &tables)

	for _, table := range tables {
		pg.store.Migrator().DropTable(table)
	}
	//Migrate table related to setting
	pg.store.Debug().AutoMigrate(&entity.Neighborhood{}, &entity.Flat{}, &entity.User{})
	pg.store.Migrator().DropConstraint(&entity.Flat{}, "fk_flats_neighborhood")
}

//SeedFile Category
func (pg *postgreSeeder) SeedFile() {

	/*

		pg.l.Println("On Flat Seeder")
		err = pg.store.Create(&flats).Error
		if err != nil {
			pg.l.Println("Seeding Flat data error :" + err.Error())
		}*/

	pg.l.Println("On User Seeder")
	err := pg.store.Create(&users).Error
	if err != nil {
		pg.l.Println("Seeding user data error :" + err.Error())
	}
	pg.l.Println("On Neighbor Seeder")
	err = pg.store.Create(&neighborhoods).Error
	if err != nil {
		pg.l.Println("Seeding Neighbor data error :" + err.Error())
	}

}

var (
	neighborhoods []*entity.Neighborhood = []*entity.Neighborhood{
		&negbor3,
		&negbor4,
	}

	negbor1 = entity.Neighborhood{
		ID:        uuid.New(),
		Name:      "Serpong",
		RangeFrom: 1,
		RangeTo:   5,
	}

	negbor2 = entity.Neighborhood{
		ID:        uuid.New(),
		Name:      "Alam Sutra",
		RangeFrom: 6,
		RangeTo:   10,
	}

	negbor3 = entity.Neighborhood{
		ID:        uuid.New(),
		Name:      "Sumaerccon",
		RangeFrom: 11,
		RangeTo:   15,
	}
	negbor4 = entity.Neighborhood{
		ID:        uuid.New(),
		Name:      "Onyx",
		RangeFrom: 16,
		RangeTo:   20,
	}

	flat1 = entity.Flat{
		ID:           uuid.New(),
		Name:         "Savanna",
		Location:     1,
		Neighborhood: &negbor1,
	}

	flat2 = entity.Flat{
		ID:           uuid.New(),
		Name:         "Nava Park",
		Location:     8,
		Neighborhood: &negbor2,
	}

	flat3 = entity.Flat{
		ID:           uuid.New(),
		Name:         "Vanya Park",
		Location:     9,
		Neighborhood: &negbor2,
		//NeighborhoodId: negbor2.ID,
	}

	users []*entity.User = []*entity.User{
		{
			ID:    uuid.New(),
			Name:  "David",
			Email: "david@gmail.com",
			//FlatId: flat1.ID,
			Flat: &flat1,
		},
		{
			ID:    uuid.New(),
			Name:  "Deddy",
			Email: "deddy@gmail.com",
			//FlatId: flat1.ID
			Flat: &flat1,
		},
		{
			ID:    uuid.New(),
			Name:  "Danny",
			Email: "danny@gmail.com",
			//FlatId: flat2.ID,
			Flat: &flat2,
		},
		{
			ID:    uuid.New(),
			Name:  "Desta",
			Email: "desta@gmail.com",
			//FlatId: flat3.ID,
			Flat: &flat3,
		},
	}
)
