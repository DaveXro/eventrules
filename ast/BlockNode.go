package ast

import "strings"

type BlockNode struct {
	statements []IAstNode
}

func (bn *BlockNode) EvaluateNode(data *LiveDataContainer) interface{} {
	for _, stmt := range bn.statements {
		stmt.EvaluateNode(data)
	}

	return true
}

func (bn *BlockNode) AddStatement(statement IAstNode) error {

	bn.statements = append(bn.statements, statement)

	return nil
}

func (cn BlockNode) String() string {
	var builder strings.Builder

	for _, str := range cn.statements {
		builder.WriteString(str.String())
		builder.WriteRune('\n')
	}

	return builder.String()
}
