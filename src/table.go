package main

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func renderTable(header, content []interface{}) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{header})
	// for _, el := range content {
	// 	// t.AppendRow(el.([]string))
	// }
	t.Render()
}
