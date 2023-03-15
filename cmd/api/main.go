package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/diyliv/interweb/config"
	"github.com/diyliv/interweb/internal/controller"
	"github.com/diyliv/interweb/internal/repository"
	"github.com/diyliv/interweb/internal/telegram"
	"github.com/diyliv/interweb/pkg/logger"
	"github.com/diyliv/interweb/pkg/storage/postgres"
)

func main() {
	logger := logger.InitLogger()
	cfg := config.ReadConfig("config", "yaml", "./config")
	psqlConn, err := postgres.ConnPostgres(cfg)
	if err != nil {
		panic(err)
	}
	repo := repository.NewRepository(psqlConn, logger)
	ctrl := controller.NewController(repo, logger)
	bot, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		panic(err)
	}
	telega := telegram.NewTelegram(bot, cfg, logger, ctrl)
	telega.StartBot()
}
