package main

import (
	"github.com/niiilov/go-yandex-1/internal/application"
	"log/slog"
)

func main() {
	app := application.NewConfig(":8090")

	err := app.Run()
	if err != nil {
		slog.Error("Error When Starting The Server: ", err)
	}
}
