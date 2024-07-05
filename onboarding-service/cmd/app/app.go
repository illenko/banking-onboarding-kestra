package app

import (
	"log/slog"

	"github.com/illenko/onboarding-service/internal/configuration"
	"github.com/illenko/onboarding-service/internal/handler"
	"github.com/illenko/onboarding-service/internal/server"
	"github.com/illenko/onboarding-service/internal/service"
)

func Run() {
	configuration.LoadEnv()

	onboardingService := service.NewOnboardingService()
	onboardingHandler := handler.NewOnboardingHandler(onboardingService)

	err := server.New(onboardingHandler).Run(":" + configuration.Get("SERVER_PORT"))
	if err != nil {
		slog.Error("Unable to start the server:", slog.String("error", err.Error()))
		return
	}
}
