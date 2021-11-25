package ast

import "strings"

type RuleContainer struct {
	Trigger IAstNode
	Plan    IAstNode
}

func (rc RuleContainer) String() string {
	var sb strings.Builder

	sb.WriteString(rc.Trigger.String())
	sb.WriteRune('\n')

	sb.WriteString(rc.Plan.String())
	sb.WriteRune('\n')
	return sb.String()
}
