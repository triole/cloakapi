package main

import (
	"github.com/triole/logseal"
)

func main() {
	parseArgs()
	kc := initKC()
	kc.initConf()
	kc.login()

	var vals any

	switch cli.Action {
	case "ls":
		switch cli.Ls.Entity {
		case getCommand(commands.List.AuthFlows):
			vals, _ = kc.fetchAuthFlows()
		case getCommand(commands.List.FedIDs):
			kc.fetchFederatedIDs()
			vals = kc.API.FedIDs
		case getCommand(commands.List.IdentityProviders):
			kc.fetchIDPs()
			vals = kc.API.IDPs
		case getCommand(commands.List.UserAttributes):
			kc.fetchUsers()
			vals = kc.API.Users
		case getCommand(commands.List.Users):
			kc.fetchUsers()
			vals = kc.API.Users
		}
	case "rm":
		switch cli.Rm.Entity {
		case getCommand(commands.List.Users):
			kc.fetchUsers()
			kc.removeUser(cli.Rm.Target)
		}
	case "tpl":
		if cli.Tpl.File != "" || cli.Tpl.String != "" {
			kc.fetchUsers()
			kc.fetchFederatedIDs()
		}
		if cli.Tpl.File != "" {
			by := kc.readFile(cli.Tpl.File)
			kc.execTemplate(string(by))
		}
		if cli.Tpl.String != "" {
			kc.execTemplate(cli.Tpl.String)
		}
	case "var":
		kc.listTemplateVars()
	}

	switch cli.Output {
	case "json":
		pprintJSON(vals)
	case "toml":
		pprintTOML(vals)
	case "yaml":
		pprintYAML(vals)
	case "table":
		kc.printTable()
	}
}

func initKC() (kc tKC) {
	kc.Lg = logseal.Init(cli.LogLevel, cli.LogFile, cli.LogNoColors, cli.LogJSON)
	return
}
