package main

import (
	"encoding/json"
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"

	toml "github.com/pelletier/go-toml/v2"
	"github.com/triole/logseal"
	"go.yaml.in/yaml/v3"
)

func (kc *tKC) initConf() {
	configFile, err := kc.detectConfigFiles(cli.Conf)
	kc.Lg.Debug("load config file", logseal.F{"path": configFile})

	raw, err := os.ReadFile(configFile)
	kc.Lg.IfErrFatal("error reading general config %q, %q", configFile, err)
	raw = []byte(os.ExpandEnv(string(raw)))
	switch filepath.Ext(configFile) {
	case ".json":
		err = json.Unmarshal(raw, &kc.Conf)
	case ".toml":
		err = toml.Unmarshal(raw, &kc.Conf)
	case ".yml", ".yaml":
		err = yaml.Unmarshal(raw, &kc.Conf)
	}
	kc.Lg.IfErrFatal("error unmarshal %q, %q", configFile, err)
}

func (kc *tKC) detectConfigFiles(expr string) (cf string, err error) {
	var configFiles []string
	userHomeDir, _ := os.UserHomeDir()
	arr := []string{
		kc.curdir(),
		path.Join(userHomeDir, ".config", "cloakapi"),
		path.Join(userHomeDir, ".conf", "cloakapi"),
		kc.executablePath(),
	}

	for _, fol := range arr {
		kc.Lg.Trace(
			"try to find config file",
			logseal.F{"expression": expr, "path": fol},
		)
		configFiles = find(fol, expr)
		if len(configFiles) > 0 {
			break
		}
	}

	if len(configFiles) == 0 {
		err = errors.New("no config file matches the expression: " + expr)
	}

	if len(configFiles) > 1 {
		err = errors.New(
			"multiple config files match the expression: " + expr +
				" -> " + strings.Join(configFiles, ","),
		)
	}
	kc.Lg.IfErrFatal("config file detection failed", logseal.F{"error": err})
	cf = configFiles[0]
	return
}

func (kc tKC) executablePath() string {
	e, err := os.Executable()
	kc.Lg.IfErrFatal(
		"can not detect binary path",
		logseal.F{"error": err},
	)
	return path.Dir(e)
}

func (kc tKC) curdir() string {
	curdir, err := os.Getwd()
	kc.Lg.IfErrError("could not detect current directory", logseal.F{"error": err})
	return curdir
}
