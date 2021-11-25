package ast

import (
	"fmt"
	"strings"

	"github.com/DaveXro/eventrules/ast/commands"
)

type CommandNode struct {
	Command    string
	Parameters []IAstNode
}

func (cn *CommandNode) EvaluateNode(data *LiveDataContainer) interface{} {
	fmt.Println("Executing Command: ", cn.Command)

	switch cn.Command {
	case "GET":
		targetUrl := cn.Parameters[0].EvaluateNode(data).(string)
		response, err := commands.GetHandler(targetUrl)
		if err != nil {
			return false
		}
		return response
	}
	return cn
}

func (cn *CommandNode) AddParameter(parameter IAstNode) error {
	cn.Parameters = append(cn.Parameters, parameter)

	return nil
}

func (cn *CommandNode) String() string {
	var builder strings.Builder

	builder.WriteString(cn.Command)
	builder.WriteRune(' ')
	for _, str := range cn.Parameters {
		builder.WriteString(str.String())
		builder.WriteRune('\n')
	}

	return builder.String()
}
