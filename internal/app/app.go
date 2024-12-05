package app

import (
	"os"

	v1 "github.com/hugokishi/hexagonal-go/internal/adapter/handler/private/v1"
	"github.com/hugokishi/hexagonal-go/internal/adapter/repository"
	"github.com/hugokishi/hexagonal-go/internal/core/drivers/db"
	"github.com/hugokishi/hexagonal-go/internal/core/drivers/log"
	"github.com/hugokishi/hexagonal-go/internal/core/drivers/rest"
	"github.com/hugokishi/hexagonal-go/internal/core/models"
	"github.com/hugokishi/hexagonal-go/internal/core/services"
)

var (
	productService *services.ProductService
)

func InitApi() {
	log.InitLogrus()

	dbConn := db.New(os.Getenv("DB_CONNECTION"))

	models.Setup(dbConn)

	store := repository.NewDatabase(dbConn)

	productService = services.NewProductService(store)

	Setup()
}

func Setup() {
	httpServer := rest.New()

	logger := log.NewLogger([]string{}, log.GetLoggerLevel())
	logger.Use(httpServer.Router)

	pv1 := httpServer.Router.Group("api/private/v1")

	v1.NewProductHandler(*productService, pv1)

	go httpServer.Run()
}
