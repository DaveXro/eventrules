package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type EventPackage struct {
	FieldMap map[string]interface{}
}

func CreatePackageFromFile(filePath string) (*EventPackage, error) {
	response := &EventPackage{}

	err := response.DecodeEventFromFile(filePath)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func CreatePackageFromBytes(bytes []byte) (*EventPackage, error) {
	response := &EventPackage{}
	err := response.DecodeEventFromBytes(bytes)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (pkg *EventPackage) DecodeEventFromFile(filePath string) error {
	jsonFile, err := os.Open(filePath)

	if checkError(err) {
		return err
	}

	defer jsonFile.Close()

	rawBytes, _ := ioutil.ReadAll(jsonFile)
	err = pkg.DecodeEventFromBytes(rawBytes)

	return err
}

func (pkg *EventPackage) DecodeEventFromBytes(rawBytes []byte) error {
	var placeholder interface{}

	err := json.Unmarshal(rawBytes, &placeholder)

	switch placeholder.(type) {
	case []interface{}:
		pkg.FieldMap = make(map[string]interface{})
		pkg.FieldMap["index"] = placeholder
	default:
		pkg.FieldMap = placeholder.(map[string]interface{})
	}

	return err
}

func (pkg *EventPackage) StringField(name string) (string, error) {
	targetEntry := pkg.findEntry(name)

	if targetEntry == nil {
		return "", fmt.Errorf("%s not found", name)
	}

	if str, ok := targetEntry.(string); ok {
		return str, nil
	} else {
		switch actualVal := targetEntry.(type) {
		case int:
			return strconv.Itoa(int(actualVal)), nil
		case float32:
			return strconv.FormatFloat(float64(actualVal), 'f', -1, 32), nil
		case float64:
			return strconv.FormatFloat(actualVal, 'f', -1, 32), nil
		}
	}

	return "", fmt.Errorf("%s cannot be converted to a string", name)
}

func (pkg *EventPackage) IntegerField(name string) (int, error) {
	targetEntry := pkg.findEntry(name)

	if targetEntry == nil {
		return -1, fmt.Errorf("%s not found", name)
	}
	switch actualValue := targetEntry.(type) {
	case string:
		return strconv.Atoi(actualValue)
	case float32:
		return int(actualValue), nil
	case float64:
		return int(actualValue), nil
	}

	intermediary := targetEntry.(float64)
	return int(intermediary), nil
}

func (pkg *EventPackage) FloatField(name string) (float32, error) {
	targetEntry := pkg.findEntry(name)

	if targetEntry == nil {
		return -1, fmt.Errorf("%s not found", name)
	}

	return float32(targetEntry.(float64)), nil
}

func (pkg *EventPackage) BoolField(name string) (bool, error) {
	targetEntry := pkg.findEntry(name)

	if targetEntry == nil {
		return false, fmt.Errorf("%s not found", name)
	}

	return targetEntry.(bool), nil
}

func (pkg *EventPackage) Field(name string) (interface{}, error) {
	targetEntry := pkg.findEntry(name)

	if targetEntry == nil {
		return -1, fmt.Errorf("%s not found", name)
	}

	return targetEntry, nil
}

func (pkg *EventPackage) findEntry(name string) interface{} {
	if strings.Contains(name, ".") {
		var retVal interface{}

		fieldSections := strings.Split(name, ".")
		searchBody := true
		currentMap := pkg.FieldMap

		for idx, s := range fieldSections {
			nextVal := pkg.findEntryInMap(s, currentMap, searchBody)
			if idx < len(fieldSections)-1 {
				currentMap = nextVal.(map[string]interface{})
			} else {
				retVal = nextVal
			}
		}

		return retVal
	} else {
		return pkg.findEntryInMap(name, pkg.FieldMap, true)
	}
}

func (pkg *EventPackage) findEntryInMap(name string, targetMap map[string]interface{}, checkBody bool) interface{} {
	for i := 0; i < 2; i++ {
		for key, element := range targetMap {
			if strings.Compare(strings.ToLower(key), strings.ToLower(name)) == 0 {
				return element
			}
		}
		if checkBody {
			targetMap = pkg.FieldMap["EventBody"].(map[string]interface{})
		} else {
			break
		}
	}

	return nil
}

func checkError(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}

	return false
}
