package correction

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Get Process correction get command
func (c *ActivityCorrectionCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	var msgText string

	id, err := strconv.Atoi(args)
	if err != nil {
		log.Println("ActivityCorrectionCommander.Get: wrong args", args)
		msgText = "Provide valid id of required correction, id must be a number (> 0)."
	} else {
		if id >= 0 {
			correction, err := c.correctionService.Describe(uint64(id))
			if err != nil {
				log.Printf("ActivityCorrectionCommander.Get: fail to get correction with ID = %d: %v", id, err)
				msgText = fmt.Sprintf("Failed to get correction with ID = %d: %v", id, err)
			} else {
				description, err := correction.String()
				if err != nil {
					log.Printf("ActivityCorrectionCommander.List: fail to get correction description: %v", err)
					msgText = fmt.Sprintf("Failed to get correction description: %v", err)
				} else {
					msgText = description
				}
			}
		} else {
			log.Printf("ActivityCorrectionCommander.Get: wrong id = %d", id)
			msgText = "Provide valid id of required correction, id must be > 0."
		}
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		msgText,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ActivityCorrectionCommander.Get: error sending reply message to chat - %v", err)
	}
}