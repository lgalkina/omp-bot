package correction

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const PageSize uint64 = 2

func (c *ActivityCorrectionCommander) createPaginationButtons(cursor uint64, limit uint64) tgbotapi.InlineKeyboardMarkup {
	inlineKeyboardMarkup := tgbotapi.NewInlineKeyboardMarkup()
	addNextButton, inlineKeyboardButton := c.createNextPageButton(cursor, limit)
	if addNextButton {
		inlineKeyboardMarkup.InlineKeyboard = append(inlineKeyboardMarkup.InlineKeyboard, inlineKeyboardButton)
	}
	addPreviousButton, inlineKeyboardButton := c.createPreviousPageButton(cursor, limit)
	if addPreviousButton {
		inlineKeyboardMarkup.InlineKeyboard = append(inlineKeyboardMarkup.InlineKeyboard, inlineKeyboardButton)
	}
	return inlineKeyboardMarkup
}

func (c *ActivityCorrectionCommander) createPreviousPageButton(cursor uint64, limit uint64) (bool, []tgbotapi.InlineKeyboardButton) {
	if cursor > 0 {
		previousCursor := cursor - limit
		if previousCursor < 0 {
			previousCursor = 0
		}
		return true, c.createPaginationButton(previousCursor, PageSize, "Previous page")
	}
	return false, nil
}

func (c *ActivityCorrectionCommander) createNextPageButton(cursor uint64, limit uint64) (bool, []tgbotapi.InlineKeyboardButton) {
	nextCursor := cursor + limit
	if c.correctionService.GetCorrectionsCount() > int(nextCursor) {
		return true, c.createPaginationButton(nextCursor, PageSize, "Next page")
	}
	return false, nil
}

func (c *ActivityCorrectionCommander) createPaginationButton(cursor uint64, limit uint64, button string) []tgbotapi.InlineKeyboardButton {
	serializedData, _ := json.Marshal(
		CallbackListData{
			Cursor: cursor,
			Limit:  limit,
		})

	callbackPath := path.CallbackPath{
		Domain:       "activity",
		Subdomain:    "correction",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	return tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(button, callbackPath.String()))
}