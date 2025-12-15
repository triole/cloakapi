package main

import (
	"fmt"

	"github.com/triole/logseal"
)

func (kc tKC) removeUser(userID string) {
	user, err := kc.getUserByID(userID)
	kc.Lg.IfErrWarn("won't remove user", logseal.F{"error": err})
	if err == nil && user != nil {
		fmt.Printf("remove user: %s, %s, %s\n", *user.Username, *user.Email, *user.ID)
		kc.Session.Client.DeleteUser(
			kc.Session.CTX, kc.Session.Token.AccessToken, kc.Conf.Realm, *user.ID,
		)
	}
}
