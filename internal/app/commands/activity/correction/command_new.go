package correction

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/activity"
	"log"
)

type createCorrection struct {
	UserID uint64 `json:"userID"`
	Object string `json:"object"`
	Action string `json:"action"`
	Data *createData `json:"data"`
	Comments string `json:"comments"`
}

type createData struct {
	OriginalData string `json:"originalData"`
	RevisedData  string `json:"revisedData"`
}

func (c *ActivityCorrectionCommander) New(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	var msgText string

	var createCorrection createCorrection
	err := json.Unmarshal([]byte(args), &createCorrection)
	if err != nil {
		log.Printf("ActivityCorrectionCommander.New: fail to unmarshall json for correction update: %v", err)
		msgText = fmt.Sprint("Failed to parse json to create correction, check command format. " +
			"Ex.: /new__activity__correction { \"userID\": 1, \"object\": \"order1\", \"action\": \"update\"," +
			"\"data:\": {\"originalData\": \"test11\", \"revisedData\": \"test12\"}, \"comments\": \"test1\"}")
	} else {
		id, err := c.correctionService.Create(
			activity.Correction{ UserID: createCorrection.UserID,
								 Object: createCorrection.Object,
								 Action: createCorrection.Action,
								 Data: &activity.Data{ OriginalData: createCorrection.Data.OriginalData,
									 					RevisedData: createCorrection.Data.RevisedData },
								 Comments: createCorrection.Comments })
		if err != nil {
			log.Printf("ActivityCorrectionCommander.New: fail to create correction : %v", err)
			msgText = fmt.Sprintf("Failed to create correction : %v", err)
		} else {
			msgText = fmt.Sprintf("Correction was successfully created with ID = %d", id)
		}
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		msgText,
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("ActivityCorrectionCommander.New: error sending reply message to chat - %v", err)
	}
}
