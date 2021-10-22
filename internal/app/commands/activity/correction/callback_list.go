package correction

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Cursor uint64 `json:"cursor"`
	Limit  uint64 `json:"limit"`
}

func (c *ActivityCorrectionCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("ActivityCorrectionCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}
	outputMsgText := "Corrections: \n"
	corrections, err := c.correctionService.List(parsedData.Cursor, parsedData.Limit)
	if err != nil {
		log.Printf("ActivityCorrectionCommander.List: fail to list the corrections: %v", err)
		outputMsgText = fmt.Sprintf("Failed to list the corrections: %v", err)

		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("ActivityCorrectionCommander.CallbackList: error sending reply message to chat - %v", err)
		}
	} else {
		for _, c := range corrections {
			description, err := c.String()
			if err != nil {
				log.Printf("ActivityCorrectionCommander.List: fail to get correction description: %v", err)
				outputMsgText = fmt.Sprintf("Failed to get correction description: %v", err)
				break
			} else {
				outputMsgText += description
				outputMsgText += "\n"
			}
		}

		msg := tgbotapi.NewMessage(
			callback.Message.Chat.ID,
			outputMsgText,
		)

		inlineKeyboardMarkup := c.createPaginationButtons(parsedData.Cursor, parsedData.Limit)
		if len(inlineKeyboardMarkup.InlineKeyboard) > 0 {
			msg.ReplyMarkup = inlineKeyboardMarkup
		}

		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("ActivityCorrectionCommander.CallbackList: error sending reply message to chat - %v", err)
		}
	}
}

