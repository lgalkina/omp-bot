package correction

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c *ActivityCorrectionCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__activity__correction — print list of commands\n\n"+

			"/get__activity__correction — get a correction, " +
				"provide id of correction as command argument. Ex.: /get__activity__correction 1\n\n"+

			"/list__activity__correction — get a list of corrections with pagination\n\n"+

			"/delete__activity__correction — delete an existing correction, " +
				"provide id of correction as command argument. Ex.: /delete__activity__correction 1\n\n"+

			"/new__activity__correction — create a new correction with required fields: " +
				"userID, object, action, originalData, and revisedData; " +
				"and optional field comments. " +
				"Ex.: /new__activity__correction { \"userID\": 1, \"object\": \"order1\", \"action\": \"update\"," +
				"\"data\": {\"originalData\": \"test11\", \"revisedData\": \"test12\"}, \"comments\": \"test1\"}\n\n"+

			"/edit__activity__correction — edit fields of correction with given id, " +
				"only optional field comments is available for editing. " +
				"Ex.: /edit__activity__correction { \"id\": 1, \"comments\": \"test3\"}\n",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("ActivityCorrectionCommander.Help: error sending reply message to chat - %v", err)
	}
}
