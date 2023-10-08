package dcnearshore

import (
	"log"

	"github.com/frangar97/testapi/internal/config"
	"github.com/frangar97/testapi/internal/entities"
	"github.com/frangar97/testapi/pkg/database"
	"gorm.io/gorm"
)

func Run() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewClient(cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.DatabasePort)
	if err != nil {
		log.Fatal(err)
	}

	err = migrateModels(db)
	if err != nil {
		log.Fatal(err)
	}
}

func migrateModels(db *gorm.DB) error {
	err := db.AutoMigrate(entities.User{}, entities.Device{}, entities.Firmware{})

	if err != nil {
		return err
	}

	return nil
}
