package main

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	toml "github.com/pelletier/go-toml/v2"
	"github.com/triole/logseal"
	"go.yaml.in/yaml/v2"
)

func main() {
	parseArgs()
	kc := initKC()
	kc.initConf()
	kc.login()

	var outp any
	var err error
	switch cli.Action {
	case "ls":
		switch cli.Ls.Entity {
		case "feds":
			outp, err = kc.listFederatedIDs()
		case "idps":
			outp, err = kc.listIDPs()
		case "users":
			outp, err = kc.listUsers()
		}
	}
	if err == nil {
		pprint(outp)
	}
}

func initKC() (kc tKC) {
	kc.Lg = logseal.Init(cli.LogLevel, cli.LogFile, cli.LogNoColors, cli.LogJSON)
	return
}

func (kc *tKC) initConf() {
	afp, err := makeAbs(cli.Conf)
	kc.Lg.IfErrFatal("can not retrieve absolute file path", logseal.F{"error": err})

	exists, err := fileExists(afp)
	kc.Lg.IfErrFatal("file exists check failed", logseal.F{"error": err})
	if !exists {
		err = errors.New("file does not exist: " + afp)
		kc.Lg.IfErrFatal("can not read config file", logseal.F{"error": err})
	}

	raw, err := os.ReadFile(afp)
	kc.Lg.IfErrFatal("error reading general config %q, %q", afp, err)
	raw = []byte(os.ExpandEnv(string(raw)))
	switch filepath.Ext(afp) {
	case ".json":
		err = json.Unmarshal(raw, &kc.Conf)
	case ".toml":
		err = toml.Unmarshal(raw, &kc.Conf)
	case ".yml", ".yaml":
		err = yaml.Unmarshal(raw, &kc.Conf)
	}
	kc.Lg.IfErrFatal("error unmarshal %q, %q", afp, err)
}
