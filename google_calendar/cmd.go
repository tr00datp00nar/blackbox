package google_calendar

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `gcal`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2024`,
	License:     `Apache-2.0`,
	Summary:     help.S(_gcal),
	Description: help.D(_gcal),

	Commands: []*Z.Cmd{
		help.Cmd,
		availableCmd,
	},
}

var availableCmd = &Z.Cmd{
	Name:        `available`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2024`,
	License:     `Apache-2.0`,
	Summary:     help.S(_availability),
	Description: help.D(_availability),

	Commands: []*Z.Cmd{
		help.Cmd,
	},

	Call: func(_ *Z.Cmd, args ...string) error {
		GetAvailability()
		return nil
	},
}
