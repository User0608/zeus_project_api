package database

import (
	"sync"

	"github.com/User0608/zeus_project_api/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var ones sync.Once
var gdb *gorm.DB

func GetDBConnextion(conf *configs.DBConfigs) (*gorm.DB, error) {
	var errr error
	ones.Do(func() {
		schema := schema.NamingStrategy{SingularTable: true}
		db, err := gorm.Open(postgres.Open(conf.GetConnectionString()), &gorm.Config{NamingStrategy: schema})
		errr = err
		gdb = db
	})
	return gdb, errr
}
