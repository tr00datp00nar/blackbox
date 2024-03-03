package update

import _ "embed"

//go:embed text/en/update.md
var _update string

//go:embed text/en/notion.md
var _notion string

//go:embed text/en/quickbooks.md
var _quickbooks string

//go:embed text/en/gcal.md
var _gcal string

//go:embed text/en/notionEvents.md
var _notionEvents string
