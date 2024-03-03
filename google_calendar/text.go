package google_calendar

import _ "embed"

//go:embed templates/header.tmpl
var HeaderTmpl string

//go:embed templates/results.tmpl
var BodyTmpl string
