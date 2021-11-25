package ast

import "fmt"

type ExpressionNode struct {
	Lhs IAstNode
	Rhs IAstNode
	Op  string
}

func (en *ExpressionNode) EvaluateNode(data *LiveDataContainer) interface{} {
	rhs := en.Rhs.EvaluateNode(data)

	switch rhs.(type) {
	case int:
		return en.calculateInts(data)
	case float32:
		return en.calculateFloats(data)
	}
	return en
}

func (en *ExpressionNode) String() string {
	return fmt.Sprintf("(%s %s %s)", en.Lhs.String(), en.Op, en.Rhs.String())
}

func (en *ExpressionNode) calculateInts(data *LiveDataContainer) int {
	lhs := en.Lhs.EvaluateNode(data).(int)
	rhs := en.Rhs.EvaluateNode(data).(int)

	switch en.Op {
	case "+":
		return lhs + rhs
	case "-":
		return lhs - rhs
	case "*":
		return lhs * rhs
	case "/":
		return lhs / rhs
	}

	return 0
}

func (en *ExpressionNode) calculateFloats(data *LiveDataContainer) float32 {
	lhs := en.Lhs.EvaluateNode(data).(float32)
	rhs := en.Rhs.EvaluateNode(data).(float32)

	switch en.Op {
	case "+":
		return lhs + rhs
	case "-":
		return lhs - rhs
	case "*":
		return lhs * rhs
	case "/":
		return lhs / rhs
	}

	return 0.0
}
