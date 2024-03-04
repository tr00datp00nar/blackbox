package notion

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
	Summary:     help.S(_notion),
	Description: help.D(_notion),

	Commands: []*Z.Cmd{
		help.Cmd,
		notionEventsCmd,
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
