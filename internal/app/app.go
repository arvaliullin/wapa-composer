package app

import (
	"log"

	"github.com/arvaliullin/wapa-composer/internal/delivery"
)

// ComposerService
type ComposerService struct {
	Config      *ComposerConfig
	HttpService delivery.HttpService
}

func (service *ComposerService) Run() {

	go func() {
		log.Printf("Starting HTTP server on %s", service.Config.Address)
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
