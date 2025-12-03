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

// keep-sorted start block=yes newline_separated=yes
func (kc *tKC) fetchFederatedIDs() {
	if len(kc.API.Users) == 0 {
		kc.fetchUsers()
	}
	if kc.API.UsersError == nil {
		for _, user := range kc.API.Users {
			feds, err := kc.Session.Client.GetUserFederatedIdentities(
				kc.Session.CTX, kc.Session.Token.AccessToken, kc.Conf.Realm,
				*user.ID,
			)
			if err == nil {
				kc.API.FedIDs = append(kc.API.FedIDs, feds...)
			}
			kc.Lg.IfErrError("could not retrieve user list", logseal.F{"error": err})
		}
	}
}

func (kc *tKC) fetchIDPs() {
	kc.API.IDPs, kc.API.IDPsError = kc.Session.Client.GetIdentityProviders(
		kc.Session.CTX, kc.Session.Token.AccessToken, kc.Conf.Realm,
	)
	kc.Lg.IfErrError("error", logseal.F{"error": kc.API.IDPsError})
}

func (kc *tKC) fetchUsers() {
	params := gocloak.GetUsersParams{}
	kc.API.Users, kc.API.UsersError = kc.Session.Client.GetUsers(
		kc.Session.CTX, kc.Session.Token.AccessToken, kc.Conf.Realm, params,
	)
	kc.Lg.IfErrError(
		"could not retrieve user list", logseal.F{"error": kc.API.UsersError},
	)
	return
}

// keep-sorted end
