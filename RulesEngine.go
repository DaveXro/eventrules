package eventrules

import "github.com/DaveXro/eventrules/ast"

func ConfigureRulesEngine(engineId string, dnsSettings map[string]string) {
	ast.EngineId = engineId
	ast.ConfigureDnsSettings(dnsSettings)
}
