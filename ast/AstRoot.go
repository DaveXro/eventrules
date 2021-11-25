package ast

type IAstNode interface {
	EvaluateNode(data *LiveDataContainer) interface{}

	String() string
}
