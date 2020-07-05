package blackboard

import (
	"backend/log"
)

type Blackboard struct {
	logger           log.Logger
	Display          *DisplayDataHolder
	UpdateChannel    chan bool
	ListenerAttached bool
	InternalData     *CustomInternalData
}

type DisplayDataHolder struct {
	Message        string             `json:"message"`
	Time           string             `json:"time"`
	GoRoutineCount int                `json:"go_routine_count"`
	Data           *CustomDisplayData `json:"data"`
}

var board *Blackboard = nil

func NewBlackboard(l log.Logger) *Blackboard {
	if board == nil {
		dataHolder := &DisplayDataHolder{
			Data: InitialCustomData(),
		}
		board = &Blackboard{
			logger:        l,
			Display:       dataHolder,
			UpdateChannel: make(chan bool),
			InternalData:  InitialCustomInternalData(),
		}
	}
	return board
}

func (b *Blackboard) UpdateDisplay() {
	if b.ListenerAttached {
		b.UpdateChannel <- true
	}
}
