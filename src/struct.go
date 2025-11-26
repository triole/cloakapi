package main

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/triole/logseal"
)

type tKC struct {
	Conf    tConf
	Lg      logseal.Logseal
	Session tSession
}

type tConf struct {
	URL          string `toml:"url" yaml:"url"`
	Realm        string `toml:"realm" yaml:"realm"`
	ClientID     string `toml:"client_id" yaml:"client_id"`
	ClientSecret string `toml:"client_secret" yaml:"client_secret"`
	Proxy        string `toml:"proxy" yaml:"proxy"`
	Insecure     bool   `toml:"insecure" yaml:"insecure"`
}

type tSession struct {
	Client *gocloak.GoCloak
	CTX    context.Context
	Token  *gocloak.JWT
}
