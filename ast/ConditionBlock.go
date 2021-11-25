package ast

import "strings"

type ConditionBlock struct {
	Conditions []IAstNode
	Action     IAstNode
}

func (cb *ConditionBlock) EvaluateNode(data *LiveDataContainer) interface{} {
	shouldContinue := true

	for _, cond := range cb.Conditions {
		result := cond.EvaluateNode(data)

		if !result.(bool) {
			shouldContinue = false
			break
		}
	}

	if shouldContinue {
		return cb.Action.EvaluateNode(data)
	}
	return false
}

func (cb *ConditionBlock) AddCondition(statement IAstNode) error {

	cb.Conditions = append(cb.Conditions, statement)

	return nil
}

func (cb ConditionBlock) String() string {
	var builder strings.Builder

	for _, str := range cb.Conditions {
		builder.WriteString(str.String())
		builder.WriteRune('\n')
	}

	builder.WriteString("Action:\n")
	if cb.Action != nil {
		builder.WriteString(cb.Action.String())
	}
	return builder.String()
}
