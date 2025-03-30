package app

import (
	"log"

	"github.com/arvaliullin/wapa-composer/internal/delivery"
	"github.com/arvaliullin/wapa-composer/internal/delivery/handlers"
	"github.com/arvaliullin/wapa-composer/internal/persistence"
)

// ComposerService
type ComposerService struct {
	Config      *ComposerConfig
	HttpService delivery.HttpService
}

func (service *ComposerService) Run() {

	go func() {
		log.Printf("Starting HTTP server on %s", service.Config.Address)

		designRepo := &persistence.DesignRepository{
			DbConnection: service.Config.DbConnection,
		}

		handlers.RegisterDesignHandler(service.HttpService, designRepo)

		service.HttpService.Start(service.Config.Address)
	}()

	select {}
}

func NewComposerService(config *ComposerConfig) *ComposerService {
	return &ComposerService{
		Config:      config,
		HttpService: delivery.NewEchoHttpService(),
	}
}
