package main

import (
	"backend/log"
)

type Blackboard struct {
	logger        log.Logger
	dataHolder    *DisplayDataHolder
	updateDisplay chan bool
}

type DisplayDataHolder struct {
	Message string `json:"message"`
	Time    string `json:"time"`
}

func NewBlackboard(l log.Logger) *Blackboard {
	dataHolder := &DisplayDataHolder{}
	return &Blackboard{
		logger:        l,
		dataHolder:    dataHolder,
		updateDisplay: make(chan bool),
	}
}
