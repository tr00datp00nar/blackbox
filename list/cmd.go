package list

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"

	gcal "github.com/tr00datp00nar/blackbox/google_calendar"
)

var Cmd = &Z.Cmd{
	Name:        `list`,
	Aliases:     []string{``},
	Usage:       `COMMAND`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2024`,
	License:     `Apache-2.0`,
	Summary:     help.S(_list),
	Description: help.D(_list),

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
		gcal.GetAvailability()
		return nil
	},
}
