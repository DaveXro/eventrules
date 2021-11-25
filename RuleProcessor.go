package eventrules

import (
	"os"

	"github.com/DaveXro/eventrules/ast"
	"github.com/DaveXro/eventrules/ast/commands"
	"github.com/DaveXro/eventrules/parser"

	"github.com/rs/zerolog/log"
)

type RulesProcessorEngine struct {
	Rule          *ast.RuleContainer
	DataContainer ast.LiveDataContainer
}

func CreateProcessorFromFile(ruleFile string, debugParse bool) (*RulesProcessorEngine, error) {
	rawRule, err := os.ReadFile(ruleFile)

	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	ruleTree, err := parser.Parse(ruleFile, rawRule, parser.Debug(debugParse))

	if err != nil {
		log.Fatal().Err(err).Msg("Parsing failed")
		return nil, err
	}

	return &RulesProcessorEngine{Rule: ruleTree.(*ast.RuleContainer)}, nil
}

func (rpe *RulesProcessorEngine) RunRuleAgainstEventFile(againstEventFile string) (bool, error) {

	eventPackage, err := commands.CreatePackageFromFile(againstEventFile)

	if err != nil {
		log.Fatal().Err(err).Msg("Error loading event")
	} else {
		rpe.DataContainer = ast.CreateLiveContainer(eventPackage)
		response := rpe.Rule.Plan.EvaluateNode(&rpe.DataContainer)

		return response.(bool), nil
	}

	return false, nil
}
