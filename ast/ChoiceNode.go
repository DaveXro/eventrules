package ast

import "strings"

type ChoiceNode struct {
	Choices []IAstNode
}

func (cn *ChoiceNode) EvaluateNode(data *LiveDataContainer) interface{} {
	for _, stmt := range cn.Choices {
		choiceResult := stmt.EvaluateNode(data)
		if choiceResult.(bool) {
			return true
		}
	}

	return false
}

func (cn *ChoiceNode) AddChoice(choice IAstNode) error {

	cn.Choices = append(cn.Choices, choice)

	return nil
}

func (cn ChoiceNode) String() string {
	var builder strings.Builder

	for _, str := range cn.Choices {
		builder.WriteString(str.String())
		builder.WriteRune('\n')
	}

	return builder.String()
}
