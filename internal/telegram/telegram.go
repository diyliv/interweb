package telegram

import (
	"errors"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"

	"github.com/diyliv/interweb/config"
	"github.com/diyliv/interweb/internal/errs"
	"github.com/diyliv/interweb/internal/interfaces"
)

type telegram struct {
	tgbot  *tgbotapi.BotAPI
	cfg    *config.Config
	logger *zap.Logger
	ctrl   interfaces.Controller
}

func NewTelegram(tgbot *tgbotapi.BotAPI, cfg *config.Config, logger *zap.Logger, ctrl interfaces.Controller) *telegram {
	return &telegram{
		tgbot:  tgbot,
		cfg:    cfg,
		logger: logger,
		ctrl:   ctrl,
	}
}

func (t *telegram) StartBot() {
	t.logger.Info("Starting telegram bot")
	u := tgbotapi.NewUpdate(0)
	updates := t.tgbot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			switch strings.HasPrefix(update.Message.Text, "/") {
			case true:
				cmd := strings.Split(update.Message.Text, " ")
				if len(cmd) == 2 && cmd[0] == "/find" {
					apiData, err := t.ctrl.FindInfo(update.Message.From.ID, cmd[1])
					if err != nil {
						t.logger.Error("Error while finding info: " + err.Error())
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Error while finding your request.")
						msg.ReplyToMessageID = update.Message.MessageID
						t.tgbot.Send(msg)
					}
					if len(apiData[0].Entries) == 0 {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "No results for your request: "+cmd[1])
						msg.ReplyToMessageID = update.Message.MessageID
						t.tgbot.Send(msg)
					} else {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("API Name: %s\nDescription: %s\nAuth: %s\nHttps: %v\nCors: %s\nLink: %s\nCategory: %s\n",
							apiData[0].Entries[0].API,
							apiData[0].Entries[0].Description,
							apiData[0].Entries[0].Auth,
							apiData[0].Entries[0].Https,
							apiData[0].Entries[0].Cors,
							apiData[0].Entries[0].Link,
							apiData[0].Entries[0].Category))
						msg.ReplyToMessageID = update.Message.MessageID
						t.tgbot.Send(msg)
					}
				} else if len(cmd) == 1 && cmd[0] == "/stats" {
					userInfo, err := t.ctrl.GetUserInfo(update.Message.From.ID)
					if err != nil {
						if errors.Is(err, errs.ErrNotFound) {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You should use bot to get some stats about your requests.")
							msg.ReplyToMessageID = update.Message.MessageID
							t.tgbot.Send(msg)
						}
					} else {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("First request time: %v\nRequests count: %v\n",
							userInfo.UserFirstRequest,
							userInfo.UserRequestsCount))
						msg.ReplyToMessageID = update.Message.MessageID
						t.tgbot.Send(msg)
					}
				} else if len(cmd) == 1 && cmd[0] == "/help" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Usage:\n/find [argument]. Example: /find weather\n/stats - get your personal stats")
					msg.ReplyToMessageID = update.Message.MessageID
					t.tgbot.Send(msg)
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Usage:\n/find [argument]. Example: /find weather\n/stats - get your personal stats")
					msg.ReplyToMessageID = update.Message.MessageID
					t.tgbot.Send(msg)
				}
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I dont understand what are you trying to do.\n /help for details.")
				msg.ReplyToMessageID = update.Message.MessageID
				t.tgbot.Send(msg)

			}
		}
	}
}
