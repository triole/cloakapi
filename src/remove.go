package main

import (
	"github.com/triole/logseal"
)

func (kc tKC) removeUser(userID string) {
	user, err := kc.getUserByID(userID)
	kc.Lg.IfErrWarn("won't remove user", logseal.F{"error": err})
	if err == nil && user != nil {
		kc.Lg.Info(
			"remove user",
			logseal.F{
				"username": deref(user.Username),
				"email":    deref(user.Email),
				"id":       deref(user.ID),
			},
		)
		kc.Session.Client.DeleteUser(
			kc.Session.CTX, kc.Session.Token.AccessToken, kc.Conf.Realm, *user.ID,
		)
	}
}
