package ast

import (
	"strings"

	"github.com/DaveXro/eventrules/ast/commands"
)

type LiveDataContainer struct {
	EventMap commands.IEventFieldMap
	FieldMap map[string]interface{}
}

func CreateLiveContainer(eventMap commands.IEventFieldMap) LiveDataContainer {
	newContainer := LiveDataContainer{EventMap: eventMap}
	newContainer.FieldMap = make(map[string]interface{})

	return newContainer
}

func (ld *LiveDataContainer) FindVariable(name string) interface{} {
	for key, element := range ld.FieldMap {
		if strings.Compare(strings.ToLower(key), strings.ToLower(name)) == 0 {
			return element
		}
	}

	return nil
}
