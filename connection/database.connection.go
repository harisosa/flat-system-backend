package connection

import (
	"log"
	"os"

	"github.com/harisosa/flat-system-backend/config"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PgConnectionPool *gorm.DB

//Pgsql  stuct
type Pgsql struct {
	logger *log.Logger
}

func NewPgsql() *Pgsql {
	return &Pgsql{log.New(os.Stdout, "Database Connection ", log.LstdFlags)}

}

//GormPg Global Variable make connection pool

func (p *Pgsql) ConnectionPool() {
	if PgConnectionPool == nil {
		if err := godotenv.Load(); err != nil {
			p.logger.Print("No .env file found")
			return
		}
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  config.ENV.DatabaseDSN,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{
			SkipDefaultTransaction: true,
		})

		if err != nil {
			p.logger.Print(err)
			return
		}
		PgConnectionPool = db
	}
}
