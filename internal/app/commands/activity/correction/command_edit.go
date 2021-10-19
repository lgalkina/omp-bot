package correction

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/activity"
	"log"
)

// updateCorrection Update can be done only on optional fields
type updateCorrection struct {
	ID uint64 `json:"id"`
	Comments string `json:"comments"`
}

func (c *ActivityCorrectionCommander) Edit(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	var msgText string

	var updateCorrection updateCorrection
	err := json.Unmarshal([]byte(args), &updateCorrection)
	if err != nil {
		log.Printf("ActivityCorrectionCommander.Edit: fail to unmarshall json for correction update: %v", err)
		msgText = fmt.Sprint("Failed to parse json to update correction, check command format: " +
			"id must be a number (> 0) and comments must be a string. " +
			"Ex.: /edit__activity__correction { \"id\": 1, \"comments\": \"test3\"}\n")
	} else {
		if updateCorrection.ID > 0 {
			err = c.correctionService.Update(updateCorrection.ID,
				activity.Correction{Comments: updateCorrection.Comments})
			if err != nil {
				log.Printf("ActivityCorrectionCommander.Edit: fail to update correction: %v", err)
				msgText = fmt.Sprintf("Failed to update correction: %v", err)
			} else {
				msgText = "Correction was successfully updated"
			}
		} else {
			log.Printf("ActivityCorrectionCommander.Edit: wrong id = %d", updateCorrection.ID)
			msgText = "Provide valid id of required correction, id must be > 0."
		}
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		msgText,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ActivityCorrectionCommander.Edit: error sending reply message to chat - %v", err)
	}
}
