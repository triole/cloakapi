package main

import (
	"errors"

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

func (kc *tKC) getUserByID(userID string) (user *gocloak.User, err error) {
	for _, usr := range kc.API.Users {
		if userID == *usr.ID {
			user = usr
			break
		}
	}
	if user == nil {
		err = errors.New("no user found matching id: " + userID)
	}
	return
}
