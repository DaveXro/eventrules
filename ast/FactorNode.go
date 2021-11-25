package ast

import (
	"fmt"
)

type TermNode struct {
	Lhs IAstNode
	Rhs IAstNode
	Op  string
}

func (tn *TermNode) EvaluateNode(data *LiveDataContainer) interface{} {
	rhs := tn.Rhs.EvaluateNode(data)

	switch rhs.(type) {
	case int:
		return tn.calculateInts(data)
	case float32:
		return tn.calculateFloats(data)
	}
	return tn
}

func (tn *TermNode) String() string {
	return fmt.Sprintf("%s %s %s", tn.Lhs.String(), tn.Op, tn.Rhs.String())
}

func (tn *TermNode) calculateInts(data *LiveDataContainer) int {
	lhs := tn.Lhs.EvaluateNode(data).(int)
	rhs := tn.Rhs.EvaluateNode(data).(int)

	switch tn.Op {
	case "*":
		return lhs * rhs
	case "/":
		return lhs / rhs
	}

	return 0
}

func (tn *TermNode) calculateFloats(data *LiveDataContainer) float32 {
	lhs := tn.Lhs.EvaluateNode(data).(float32)
	rhs := tn.Rhs.EvaluateNode(data).(float32)

	switch tn.Op {
	case "*":
		return lhs * rhs
	case "/":
		return lhs / rhs
	}

	return 0.0
}
