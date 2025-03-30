package main

import (
	"flag"

	"github.com/arvaliullin/wapa-composer/internal/app"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Путь к конфигурации сервиса")
	flag.Parse()

	config, err := app.NewComposerConfig(*configPath)
	app.FatalOnError(err, "Ошибка загрузки конфигурации сервиса: %v", err)

	service := app.NewComposerService(config)
	service.Run()
}
