package correction

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *ActivityCorrectionCommander) List(inputMessage *tgbotapi.Message) {

	if c.correctionService.GetCorrectionsCount() == 0 {
		outputMsgText := "There is no corrections, use /new__activity__correction command to add corrections"

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("ActivityCorrectionCommander.List: error sending reply message to chat - %v", err)
		}
	} else {
		outputMsgText := "Corrections: \n\n"

		corrections, err := c.correctionService.List(0, PageSize)
		if err != nil {
			log.Printf("ActivityCorrectionCommander.List: fail to list all the corrections: %v", err)
			outputMsgText = fmt.Sprintf("Failed to list all the corrections: %v", err)

			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

			_, err = c.bot.Send(msg)
			if err != nil {
				log.Printf("ActivityCorrectionCommander.List: error sending reply message to chat - %v", err)
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

			msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

			inlineKeyboardMarkup := c.createPaginationButtons(0, PageSize)
			if len(inlineKeyboardMarkup.InlineKeyboard) > 0 {
				msg.ReplyMarkup = inlineKeyboardMarkup
			}

			_, err = c.bot.Send(msg)
			if err != nil {
				log.Printf("ActivityCorrectionCommander.List: error sending reply message to chat - %v", err)
			}
		}
	}
}