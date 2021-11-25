package ifrules

import "github.com/invoice-fair/ifrules/ast"

func ConfigureRulesEngine(engineId string, dnsSettings map[string]string) {
	ast.EngineId = engineId
	ast.ConfigureDnsSettings(dnsSettings)
}
