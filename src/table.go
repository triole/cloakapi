package main

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func (kc *tKC) printTable() {
	switch cli.Action {
	case "ls":
		switch cli.Ls.Entity {
		case "feds":
			kc.printFederatedIDsTable()
			// case "idps":
			// 	kc.printIDPsTable()
			// case "users":
			// 	kc.printUsersTable()
		}
	}
}

func (kc tKC) printFederatedIDsTable() {
	header := []string{
		"user name",
		"first name",
		"last name",
		"email",
		"local id",
		"remote id",
	}
	var content [][]any
	for _, user := range kc.API.Users {
		line := []any{
			derefString(user.Username),
			derefString(user.FirstName),
			derefString(user.LastName),
			derefString(user.Email),
			derefString(user.ID),
		}
		content = append(content, line)
	}
	renderTable(header, content)
}

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
