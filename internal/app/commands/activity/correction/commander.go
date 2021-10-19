package correction

import (
	"github.com/ozonmp/omp-bot/internal/service/activity/correction"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type ActivityCorrectionCommander struct {
	bot              *tgbotapi.BotAPI
	correctionService *correction.DummyCorrectionService
}

func NewActivityCorrectionCommander(
	bot *tgbotapi.BotAPI,
) *ActivityCorrectionCommander {
	correctionService := correction.NewDummyCorrectionService()

	return &ActivityCorrectionCommander{
		bot:              bot,
		correctionService: correctionService,
	}
}

func (c *ActivityCorrectionCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("ActivityCorrectionCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *ActivityCorrectionCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}

func (c *ActivityCorrectionCommander) HandleCommandResponseSend(msg *tgbotapi.MessageConfig) {
	_, err := c.bot.Send(msg)
	if err != nil {
		c.HandleCommandErrorLog("error sending reply message to chat - %v", err)
	}
}

func (c *ActivityCorrectionCommander) HandleCommandErrorLog(msg string, err error) {
	log.Println("ActivityCorrectionCommander: ", msg, ": ", err)
}