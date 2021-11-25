package processor

import (
	"fmt"
	"testing"
)

func TestRuleProcessor(t *testing.T) {
	ruleProcessor, err := CreateProcessorFromFile("./sample1.rule", false)

	if err != nil {
		t.Fail()
	} else {
		result, err := ruleProcessor.RunRuleAgainstEventFile("../LeadEvent.json")
		if err != nil {
			t.Fail()
		}
		fmt.Println("Rule Result: ", result)
	}
}
