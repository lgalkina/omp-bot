package activity

import (
	"encoding/json"
	"time"
)

type Correction struct {
	ID uint64  `json:"id"`// required
	Timestamp time.Time  `json:"timestamp"`// required
	UserID uint64  `json:"userID"`// required
	Object string  `json:"object"`// required
	Action string  `json:"action"`// required
	Data *Data  `json:"data"`// required
	Comments string  `json:"comments"`// optional
}

func (c *Correction) String() (string, error) {
	json, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return "", err
	}
	return string(json), nil
}

func (c *Correction) MarshalJSON() ([]byte, error) {
	type Alias Correction
	return json.Marshal(&struct {
		Timestamp string `json:"timestamp"`
		*Alias
	}{
		Timestamp: c.Timestamp.Format(time.RFC850),
		Alias:    (*Alias)(c),
	})
}

type Data struct {
	OriginalData string `json:"originalData"`
	RevisedData string `json:"revisedData"`
}

func (d *Data) String() (string, error) {
	json, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		return "", err
	}
	return string(json), nil
}