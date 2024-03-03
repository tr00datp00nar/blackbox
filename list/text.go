package list

import (
	_ "embed"
)

//go:embed text/en/list.md
var _list string

//go:embed text/en/availability.md
var _availability string

//go:embed templates/header.tmpl
var headerTmpl string

//go:embed templates/results.tmpl
var bodyTmpl string
