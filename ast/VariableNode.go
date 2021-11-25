package ast

import (
	"fmt"
)

type VariableNode struct {
	Name  string
	Value interface{}
}

func (vn *VariableNode) EvaluateNode(data *LiveDataContainer) interface{} {
	if len(vn.Name) > 0 {
		if vn.Name == "empty" {
			return "NULLCHECK"
		}

		eventField, _ := data.EventMap.Field(vn.Name)
		if eventField != nil {
			vn.Value = eventField
		} else {
			vn.Value = data.FindVariable(vn.Name)
		}

	}
	return vn.Value
}

func (vn *VariableNode) String() string {
	if vn.Value == nil {
		return vn.Name
	}
	return fmt.Sprintf("%s %v", vn.Name, vn.Value)
}
