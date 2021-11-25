package ast

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ConditionNode struct {
	Lhs IAstNode
	Rhs IAstNode
	Op  string
}

func (cn *ConditionNode) EvaluateNode(data *LiveDataContainer) interface{} {
	rhs := cn.Rhs.EvaluateNode(data)

	switch rhs.(type) {
	case int:
		if res, err := cn.compareInts(data); err == nil {
			return res
		} else {
			fmt.Println("Error: ", err)
		}
		return false
	case string:
		if strings.Compare(rhs.(string), "NULLCHECK") == 0 {
			if res, err := cn.compareNull(data); err == nil {
				return res
			} else {
				fmt.Println("Error: ", err)
			}
			return false
		}

		if res, err := cn.compareStrings(data); err == nil {
			return res
		} else {
			fmt.Println("Error: ", err)
		}
		return false
	case float32:
		if res, err := cn.compareFloats(data); err == nil {
			return res
		} else {
			fmt.Println("Error: ", err)
		}
		return false

	}
	return cn
}

func (cn *ConditionNode) String() string {
	return fmt.Sprintf("%s %s %s", cn.Lhs.String(), cn.Op, cn.Rhs.String())
}

func (cn *ConditionNode) compareNull(data *LiveDataContainer) (bool, error) {
	lhs := cn.Lhs.EvaluateNode(data)

	if lhs == nil && strings.Compare(cn.Op, "==") == 0 {
		return true, nil
	}
	if lhs != nil && strings.Compare(cn.Op, "!=") == 0 {
		return true, nil
	}

	return false, nil
}

func (cn *ConditionNode) compareInts(data *LiveDataContainer) (bool, error) {
	lhs := cleanInt(cn.Lhs.EvaluateNode(data))
	rhs := cleanInt(cn.Rhs.EvaluateNode(data))

	switch cn.Op {
	case "<":
		return lhs < rhs, nil
	case "<=":
		return lhs <= rhs, nil
	case ">":
		return lhs > rhs, nil
	case ">=":
		return lhs >= rhs, nil
	case "==":
		return lhs == rhs, nil
	case "!=":
		return lhs != rhs, nil
	}

	return false, errors.New("no matching operator found")
}

func (cn *ConditionNode) compareStrings(data *LiveDataContainer) (bool, error) {
	lhs := cn.Lhs.EvaluateNode(data).(string)
	rhs := cn.Rhs.EvaluateNode(data).(string)
	comparisonResult := strings.Compare(lhs, rhs)

	switch cn.Op {
	case "<":
		return comparisonResult == -1, nil
	case ">":
		return comparisonResult == 1, nil
	case "==":
		return comparisonResult == 0, nil
	case "!=":
		return comparisonResult != 0, nil
	}

	return false, errors.New("no supported operator found")
}

func (cn *ConditionNode) compareFloats(data *LiveDataContainer) (bool, error) {
	lhs := cleanFloat(cn.Lhs.EvaluateNode(data))
	rhs := cleanFloat(cn.Rhs.EvaluateNode(data))

	switch cn.Op {
	case "<":
		return lhs < rhs, nil
	case "<=":
		return lhs <= rhs, nil
	case ">":
		return lhs > rhs, nil
	case ">=":
		return lhs >= rhs, nil
	case "==":
		return lhs == rhs, nil
	case "!=":
		return lhs != rhs, nil
	}

	return false, errors.New("no matching operator found")
}

func cleanFloat(unformattedNumber interface{}) float32 {

	if formattedNumber, castOk := unformattedNumber.(float64); castOk {
		return float32(formattedNumber)
	}

	return unformattedNumber.(float32)
}

func cleanInt(unformattedNumber interface{}) int {

	if formattedNumber, castOk := unformattedNumber.(float64); castOk {
		return int(formattedNumber)
	}
	if formattedNumber, castOk := unformattedNumber.(float32); castOk {
		return int(formattedNumber)
	}
	if formattedNumber, castOk := unformattedNumber.(string); castOk {
		if parsed, err := strconv.Atoi(formattedNumber); err == nil {
			return parsed
		} else {
			return 0
		}
	}

	return unformattedNumber.(int)
}
