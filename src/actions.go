package main

import (
	"context"
	"crypto/tls"

	"github.com/Nerzal/gocloak/v13"
	"github.com/triole/logseal"
)

func (kc *tKC) login() {
	var err error
	kc.Session.Client = gocloak.NewClient(kc.Conf.URL)

	if kc.Lg.Logrus.Level > 5 {
		kc.Session.Client.RestyClient().SetDebug(true)
	}

	if kc.Conf.Insecure {
		kc.Session.Client.RestyClient().SetTLSClientConfig(
			&tls.Config{InsecureSkipVerify: true},
		)
	}

	if kc.Conf.Proxy != "" {
		kc.Session.Client.RestyClient().SetProxy(kc.Conf.Proxy)
	}
	kc.Lg.Debug(
		"proxy",
		logseal.F{
			"is_set": kc.Session.Client.RestyClient().IsProxySet(),
			"proxy":  kc.Conf.Proxy,
		},
	)

	kc.Session.CTX = context.Background()
	kc.Session.Token, err = kc.Session.Client.LoginClient(
		kc.Session.CTX,
		kc.Conf.ClientID,
		kc.Conf.ClientSecret,
		kc.Conf.Realm,
	)
	kc.Lg.IfErrFatal("login failed", logseal.F{"error": err})
}

// listAllUsers fetches a list of all users from Keycloak
func (kc *tKC) listUsers() {
	params := gocloak.GetUsersParams{}
	users, err := kc.Session.Client.GetUsers(
		kc.Session.CTX, kc.Session.Token.AccessToken, kc.Conf.Realm, params,
	)
	kc.Lg.IfErrError("could not retrieve user list", logseal.F{"error": err})
	if err == nil {
		pprint(users)
	}
}

func (kc *tKC) listIDPs() {
	idps, err := kc.Session.Client.GetIdentityProviders(
		kc.Session.CTX, kc.Session.Token.AccessToken, kc.Conf.Realm,
	)
	kc.Lg.IfErrError("error", logseal.F{"error": err})
	if err == nil {
		pprint(idps)
	}
}
