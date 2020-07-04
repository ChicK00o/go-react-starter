package blackboard

import (
	"backend/log"
)

type Blackboard struct {
	logger           log.Logger
	DataHolder       *DisplayDataHolder
	UpdateChannel    chan bool
	ListenerAttached bool
}

type DisplayDataHolder struct {
	Message        string      `json:"message"`
	Time           string      `json:"time"`
	GoRoutineCount int         `json:"go_routine_count"`
	Data           interface{} `json:"data"`
}

var board *Blackboard = nil

func NewBlackboard(l log.Logger) *Blackboard {
	if board == nil {
		dataHolder := &DisplayDataHolder{}
		board = &Blackboard{
			logger:        l,
			DataHolder:    dataHolder,
			UpdateChannel: make(chan bool),
		}
	}
	return board
}

func (b *Blackboard) UpdateDisplay() {
	if b.ListenerAttached {
		b.UpdateChannel <- true
	}
}
