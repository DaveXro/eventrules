package ast

import (
	"strings"

	"github.com/invoice-fair/ifrules/ast/commands"

	"github.com/google/uuid"
)

type RaiseEventNode struct {
	EventName   string
	MessageBody IAstNode
}

var EngineId string
var privateDns map[string]string

func ConfigureDnsSettings(services map[string]string) {
	privateDns = services
}

func (ren *RaiseEventNode) EvaluateNode(data *LiveDataContainer) interface{} {
	if privateDns == nil {
		panic("No DNS settings")
	}

	decodedBody := ren.MessageBody.EvaluateNode(data)

	if decodedBody != nil {
		eventEnvelope := make(map[string]interface{})

		typeCategory := strings.Split(ren.EventName, ":")

		eventEnvelope["EventId"] = uuid.New().String()
		eventEnvelope["EventCategory"] = typeCategory[0]
		eventEnvelope["EventType"] = typeCategory[1]
		eventEnvelope["RaisingSystem"] = EngineId
		eventEnvelope["RaisingUser"] = "System"
		eventEnvelope["EventBody"] = decodedBody.(map[string]interface{})

		err := commands.PostHandler(privateDns["EventStore"], eventEnvelope)

		return err == nil
	}
	return false
}

func (ren *RaiseEventNode) String() string {
	var builder strings.Builder

	builder.WriteString("{\n")

	builder.WriteString("\"EventBody\" : ")
	builder.WriteString(ren.MessageBody.String())

	builder.WriteString("}\n")

	return builder.String()
}
