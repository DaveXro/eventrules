package ast

import "fmt"

type ValueNode struct {
	Value interface{}
}

func (vn *ValueNode) EvaluateNode(data *LiveDataContainer) interface{} {
	return vn.Value
}

func (vn *ValueNode) String() string {
	switch vn.Value.(type) {
	case int:
		return fmt.Sprintf("%d", vn.Value.(int))
	case string:
		return fmt.Sprintf("\"%s\"", vn.Value.(string))
	case float32:
		return fmt.Sprintf("%f", vn.Value.(float32))
	case float64:
		return fmt.Sprintf("%f", vn.Value.(float64))
	case bool:
		return fmt.Sprintf("%v", vn.Value.(bool))
	}

	return fmt.Sprintf("%v", vn.Value)
}
