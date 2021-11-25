package ast

import (
	"strings"
)

type JsonBodyNode struct {
	Entries map[string]interface{}
}

func CreateJsonBodyNode() *JsonBodyNode {
	node := &JsonBodyNode{}
	node.Entries = make(map[string]interface{})

	return node
}

func (jn *JsonBodyNode) EvaluateNode(data *LiveDataContainer) interface{} {
	evaluatedBody := make(map[string]interface{})

	for key, value := range jn.Entries {
		if castType, canCast := jn.Entries[key].(IAstNode); canCast {
			evaluatedBody[key] = castType.EvaluateNode(data)
		} else {
			evaluatedBody[key] = value
		}
	}
	return evaluatedBody
}

func (jn *JsonBodyNode) String() string {
	var builder strings.Builder

	builder.WriteString("{\n")

	for key, entry := range jn.Entries {
		builder.WriteRune('\t')

		builder.WriteRune('"')
		builder.WriteString(key)
		builder.WriteRune('"')
		builder.WriteString(" : ")
		builder.WriteString(entry.(IAstNode).String())
		builder.WriteRune('\n')
	}

	builder.WriteString("}\n")

	return builder.String()
}
