package google_calendar

import _ "embed"

//go:embed text/en/gcal.md
var _gcal string

//go:embed text/en/availability.md
var _availability string

//go:embed templates/header.tmpl
var HeaderTmpl string

//go:embed templates/results.tmpl
var BodyTmpl string
