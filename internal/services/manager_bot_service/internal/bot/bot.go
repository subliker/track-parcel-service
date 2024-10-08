package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
)

type BotConfig struct {
	Token string
}

func Run(bcfg BotConfig) {
	bot, err := tgbotapi.NewBotAPI(bcfg.Token)
	if err != nil {
		logger.Zap.Fatalf("error creating bot: %s", err)
	}

	bot.Debug = true

	logger.Zap.Infof("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			logger.Zap.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
