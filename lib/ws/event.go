package ws

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// EventHandler -
type EventHandler func(*Event)

// Event -
type Event struct {
	ID   string      `json:"id,omitempty"`
	Name string      `json:"method"`
	Data interface{} `json:"params"`
}

// NewEvent -
func NewEvent(rawData []byte) (*Event, error) {
	e := new(Event)

	err := json.Unmarshal(rawData, e)
	if err != nil {
		return nil, err
	}

	return e, nil
}

// Raw -
func (e *Event) Raw() []byte {
	raw, _ := json.Marshal(e)
	return raw
}

// DataRaw -
func (e *Event) DataRaw() []byte {
	raw, _ := json.Marshal(e.Data)
	return raw
}
