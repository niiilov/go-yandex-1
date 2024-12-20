package application

import (
	"github.com/niiilov/go-yandex-1/internal/handlers"
	"log/slog"
	"net/http"
)

type Config struct {
	Addr string
}

func NewConfig(addr string) *Config {
	if addr == "" {
		addr = ":8090"
	}
	slog.Info("Server Has Been Assigned An Address", "Addr", addr)
	return &Config{
		Addr: addr,
	}
}

func (c Config) Run() error {
	http.HandleFunc("/api/v1/calculate", handlers.CalcHandler)
	slog.Info("Server Has Started Successfully")

	return http.ListenAndServe(c.Addr, nil)
}
