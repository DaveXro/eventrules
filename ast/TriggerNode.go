package ast

import (
	"fmt"
	"strings"
)

type TriggerNode struct {
	EventName string
}

func (tn *TriggerNode) EvaluateNode(data *LiveDataContainer) interface{} {

	eventType, err := data.EventMap.StringField("EventType")

	if err != nil {
		panic("error getting event type")
	}

	eventCategory, err := data.EventMap.StringField("EventCategory")

	if err != nil {
		panic("error getting event category")
	}

	typeCategory := strings.Split(tn.EventName, ":")

	if len(typeCategory) != 2 {
		return false
	}

	match := strings.EqualFold(eventType, typeCategory[0]) && strings.EqualFold(eventCategory, typeCategory[1])
	if !match {
		panic("Event Name doesn't match event")
	}

	return true
}

func (tn *TriggerNode) String() string {
	return fmt.Sprintf("ON %s", tn.EventName)
}
