package main

func (kc *tKC) getFedID(userName string) (remID, remIDP string) {
	for _, fed := range kc.API.FedIDs {
		if userName == *fed.UserName {
			return *fed.UserID, *fed.IdentityProvider
		}
	}
	return
}
