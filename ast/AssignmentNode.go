package ast

import "fmt"

type AssignmentNode struct {
	Name string
	Rhs  IAstNode
}

func (an *AssignmentNode) EvaluateNode(data *LiveDataContainer) interface{} {
	rhs := an.Rhs.EvaluateNode(data)

	data.FieldMap[an.Name] = rhs

	return true
}

func (an *AssignmentNode) String() string {
	return fmt.Sprintf("%s = %s", an.Name, an.Rhs.String())
}
