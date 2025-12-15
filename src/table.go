package main

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func (kc *tKC) printTable() {
	switch cli.Action {
	case "ls":
		switch cli.Ls.Entity {
		// keep-sorted start block=yes
		case getCommand(commands.List.FedIDs):
			kc.printFederatedIDsTable()
		case getCommand(commands.List.IdentityProviders):
			kc.printIDPsTable()
		case getCommand(commands.List.UserAttributes):
			kc.printUserAttributes()
		case getCommand(commands.List.Users):
			kc.printUsersTable()
			// keep-sorted end
		}
	}
}

// keep-sorted start block=yes newline_separated=yes
func (kc tKC) printFederatedIDsTable() {
	header := []string{
		"user name",
		"email",
		"local id",
		"remote id",
		"remote idp",
	}
	var content [][]any
	for _, user := range kc.API.Users {
		userName := deref(user.Username)
		remID, remIDP := kc.getFedID(userName)
		line := []any{
			userName,
			deref(user.Email),
			deref(user.ID),
			remID,
			remIDP,
		}
		content = append(content, line)
	}
	renderTable(header, content)
}

func (kc tKC) printIDPsTable() {
	header := []string{
		"name",
		"alias",
		"internal id",
		"enabled",
		"link only",
	}
	var content [][]any
	for _, idp := range kc.API.IDPs {
		line := []any{
			deref(idp.DisplayName),
			deref(idp.Alias),
			deref(idp.InternalID),
			*idp.Enabled,
			*idp.LinkOnly,
		}
		content = append(content, line)
	}
	renderTable(header, content)
}

func (kc tKC) printUserAttributes() {
	header := []string{
		"user name",
		"email",
	}
	var content [][]any
	for _, user := range kc.API.Users {
		userName := deref(user.Username)
		line := []any{
			userName,
			deref(user.Email),
			fmtYAML(user.Attributes),
		}
		content = append(content, line)
	}
	renderTable(header, content)
}

func (kc tKC) printUsersTable() {
	header := []string{
		"user id",
		"user name",
		"first name",
		"last name",
		"email",
		"user enabled",
		"email verified",
	}
	var content [][]any
	for _, user := range kc.API.Users {
		userName := deref(user.Username)
		line := []any{
			deref(user.ID),
			userName,
			deref(user.FirstName),
			deref(user.LastName),
			deref(user.Email),
			*user.Enabled,
			*user.EmailVerified,
		}
		content = append(content, line)
	}
	renderTable(header, content)
}

// keep-sorted end

func renderTable(header []string, content [][]any) {
	t := table.NewWriter()
	t.SetStyle(table.Style{
		Name: "myNewStyle",
		Box: table.BoxStyle{
			BottomLeft:       "\\",
			BottomRight:      "/",
			BottomSeparator:  "v",
			Left:             "[",
			LeftSeparator:    "{",
			MiddleHorizontal: "-",
			MiddleSeparator:  "+",
			MiddleVertical:   "|",
			PaddingLeft:      " ",
			PaddingRight:     " ",
			Right:            " ]",
			RightSeparator:   "}",
			TopLeft:          "(",
			TopRight:         ")",
			TopSeparator:     "^",
			UnfinishedRow:    " ~~~",
		},
		Options: table.Options{
			DrawBorder:      false,
			SeparateColumns: true,
			SeparateFooter:  true,
			SeparateHeader:  true,
			SeparateRows:    false,
		},
	})
	t.SetOutputMirror(os.Stdout)
	var headerRow table.Row
	for _, el := range header {
		headerRow = append(headerRow, el)
	}
	t.AppendHeader(headerRow)
	for _, el := range content {
		t.AppendRow(el)
	}
	t.Render()
}
