package eventrules

import (
	"fmt"
	"testing"

	"github.com/DaveXro/eventrules/processor"
)

func TestRulesEngine(t *testing.T) {
	dns := map[string]string{"EventStore": "https://httpbin.org/post"}

	ConfigureRulesEngine("RulesEngineTest", dns)

	ruleProcessor, err := processor.CreateProcessorFromFile("./sample.rule", false)

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
