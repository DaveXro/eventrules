package eventrules

import (
	"fmt"
	"testing"
)

func TestRulesEngine(t *testing.T) {
	dns := map[string]string{"EventStore": "https://httpbin.org/post"}

	ConfigureRulesEngine("RulesEngineTest", dns)

	ruleProcessor, err := CreateProcessorFromFile("./sample.rule", false)

	if err != nil {
		t.Fail()
	} else {
		result, err := ruleProcessor.RunRuleAgainstEventFile("./LeadEvent.json")
		if err != nil {
			t.Fail()
		}
		fmt.Println("Rule Result: ", result)
	}
}
func TestRuleProcessor(t *testing.T) {
	ruleProcessor, err := CreateProcessorFromFile("./sample1.rule", false)

	if err != nil {
		t.Fail()
	} else {
		result, err := ruleProcessor.RunRuleAgainstEventFile("./LeadEvent.json")
		if err != nil {
			t.Fail()
		}
		fmt.Println("Rule Result: ", result)
	}
}
