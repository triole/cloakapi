package main

import (
	"github.com/triole/logseal"
)

func main() {
	parseArgs()
	kc := initKC()
	kc.initConf()
	kc.login()

	switch cli.Action {
	case "ls":
		switch cli.Ls.Entity {
		case getCommand(commands.List.FedIDs):
			kc.fetchFederatedIDs()
		case getCommand(commands.List.IdentityProviders):
			kc.fetchIDPs()
		case getCommand(commands.List.UserAttributes):
			kc.fetchUsers()
		case getCommand(commands.List.Users):
			kc.fetchUsers()
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
	}

	switch cli.Output {
	case "json":
		pprintJSON(kc.API)
	case "toml":
		pprintTOML(kc.API)
	case "yaml":
		pprintYAML(kc.API)
	case "table":
		kc.printTable()
	}
}

func initKC() (kc tKC) {
	kc.Lg = logseal.Init(cli.LogLevel, cli.LogFile, cli.LogNoColors, cli.LogJSON)
	return
}
