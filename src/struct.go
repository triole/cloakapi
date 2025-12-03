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
	API     tApiData
}

type tConf struct {
	URL          string `toml:"url" yaml:"url"`
	Realm        string `toml:"realm" yaml:"realm"`
	ClientID     string `toml:"client_id" yaml:"client_id"`
	ClientSecret string `toml:"client_secret" yaml:"client_secret"`
	Proxy        string `toml:"proxy" yaml:"proxy"`
	Insecure     bool   `toml:"insecure" yaml:"insecure"`
}

type tApiData struct {
	Users       []*gocloak.User                            `json:"users,omitempty" toml:"users,omitempty" yaml:"users,omitempty"`
	UsersError  error                                      `json:"-" toml:"-" yaml:"-"`
	FedIDs      []*gocloak.FederatedIdentityRepresentation `json:"federated_ids,omitempty" toml:"federated_ids,omitempty" yaml:"federated_ids,omitempty"`
	FedIDsError error                                      `json:"-" toml:"-" yaml:"-"`
	IDPs        []*gocloak.IdentityProviderRepresentation  `json:"idps,omitempty" toml:"idps,omitempty" yaml:"idps,omitempty"`
	IDPsError   error                                      `json:"-" toml:"-" yaml:"-"`
}

type tSession struct {
	Client *gocloak.GoCloak
	CTX    context.Context
	Token  *gocloak.JWT
}
