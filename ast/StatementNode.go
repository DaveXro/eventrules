package ast

import "fmt"

type StatementNode struct {
	name string
}

func (sn *StatementNode) EvaluateNode(data *LiveDataContainer) interface{} {
	return sn
}

func (sn *StatementNode) String() string {
	return fmt.Sprintf("%s", sn.name)
}
