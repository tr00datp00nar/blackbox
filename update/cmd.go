package update

import (
	"context"
	"fmt"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `update`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2024`,
	License:     `Apache-2.0`,
	Summary:     help.S(_update),
	Description: help.D(_update),

	Commands: []*Z.Cmd{
		help.Cmd,
		gcalCmd,
		notionCmd,
		quickbooksCmd,
	},
}

var notionCmd = &Z.Cmd{
	Name:        `notion`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2024`,
	License:     `Apache-2.0`,
	Summary:     help.S(_notion),
	Description: help.D(_notion),

	Commands: []*Z.Cmd{
		help.Cmd,
		notionEventsCmd,
	},
}

var quickbooksCmd = &Z.Cmd{
	Name:        `quickbooks`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2024`,
	License:     `Apache-2.0`,
	Summary:     help.S(_quickbooks),
	Description: help.D(_quickbooks),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		// TODO
		return nil
	},
}

var gcalCmd = &Z.Cmd{
	Name:        `gcal`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2024`,
	License:     `Apache-2.0`,
	Summary:     help.S(_gcal),
	Description: help.D(_gcal),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		// TODO
		return nil
	},
}

var notionEventsCmd = &Z.Cmd{
	Name:        `events`,
	Aliases:     []string{``},
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2024`,
	License:     `Apache-2.0`,
	Summary:     help.S(_notionEvents),
	Description: help.D(_notionEvents),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		client := ConnectNotion()
		db := GetScheduledEvents(context.Background(), client)
		fmt.Println(db)
		return nil
	},
}
