package dcnearshore

import (
	"fmt"
	"log"

	"github.com/frangar97/testapi/internal/config"
	"github.com/frangar97/testapi/internal/entities"
	"github.com/frangar97/testapi/internal/handlers"
	"github.com/frangar97/testapi/internal/repository"
	"github.com/frangar97/testapi/internal/service"
	"github.com/frangar97/testapi/pkg/database"
	"github.com/labstack/echo/v4"
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

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories, cfg.Secret)
	handlers := handlers.NewHandlers(services, cfg.Secret)

	mux := echo.New()

	registerRoutes(mux, handlers)

	mux.Start(fmt.Sprintf(":%s", cfg.Port))
}

func migrateModels(db *gorm.DB) error {
	err := db.AutoMigrate(entities.User{}, entities.Device{}, entities.Firmware{})

	if err != nil {
		return err
	}

	return nil
}

func registerRoutes(mux *echo.Echo, handlers handlers.Handlers) {
	//user routes
	mux.POST("/api/user", handlers.CreateUserHandler)
	mux.POST("/api/user/login", handlers.LoginUserHandler)

	//device routes
	mux.POST("/api/devices", handlers.CreateDeviceHandler)
	mux.GET("/api/devices", handlers.GetAllDevicesHandler)
	mux.GET("/api/devices/:id", handlers.GetDeviceByIdHandler)
}
