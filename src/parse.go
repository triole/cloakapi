package main

import (
	"fmt"

	"github.com/Nerzal/gocloak/v13"
)

func (kc *tKC) getFedID(userName string) (remID, remIDP string) {
	for _, fed := range kc.API.FedIDs {
		if userName == *fed.UserName {
			return *fed.UserID, *fed.IdentityProvider
		}
	}
	return
}

func (kc *tKC) getUserAttributes(user *gocloak.User) (ret string) {
	ret = fmt.Sprintf("%s", user.Attributes)
	return
}
