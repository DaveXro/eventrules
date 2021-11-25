package ast

import (
	"fmt"
)

type AltClauseNode struct {
	Lhs IAstNode
	Rhs IAstNode
	Op  string
}

func (acn *AltClauseNode) EvaluateNode(data *LiveDataContainer) interface{} {
	rhs := acn.Rhs.EvaluateNode(data)
	lhs := acn.Lhs.EvaluateNode(data)

	switch acn.Op {
	case "OR":
		return rhs.(bool) || lhs.(bool)
	case "AND":
		return rhs.(bool) && lhs.(bool)
	}

	return false
}

func (acn *AltClauseNode) String() string {
	return fmt.Sprintf("%s %s %s", acn.Lhs.String(), acn.Op, acn.Rhs.String())
}
