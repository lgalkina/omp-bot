package correction

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

// Delete Process correction delete command
func (c *ActivityCorrectionCommander) Delete(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	var msgText string

	id, err := strconv.Atoi(args)
	if err != nil {
		log.Println("ActivityCorrectionCommander.Delete: wrong args", args)
		msgText = "Provide valid id of required correction, id must be a number (> 0)."
	} else {
		if id >= 0 {
			result, err := c.correctionService.Remove(uint64(id))
			if err != nil {
				log.Printf("ActivityCorrectionCommander.Delete: fail to delete correction with ID = %d: %v", id, err)
				msgText = fmt.Sprintf("Failed to delete correction with ID = %d: %v", id, err)
			} else {
				if result {
					msgText = fmt.Sprintf("Correction with ID = %d was successfully deleted", id)
				} else {
					msgText = fmt.Sprintf("Correction with ID = %d was not deleted", id)
				}
			}
		} else {
			log.Printf("ActivityCorrectionCommander.Delete: wrong id = %d", id)
			msgText = "Provide valid id of required correction, id must be > 0."
		}
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		msgText,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ActivityCorrectionCommander.Delete: error sending reply message to chat - %v", err)
	}
}
