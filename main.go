package main

import (
	_ "embed"
	"log"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/good" // Common Go tools
	"github.com/rwxrob/help"

	"github.com/tr00datp00nar/blackbox/list"
)

func init() {
	Z.Dynamic["uname"] = func() string { return Z.Out("uname", "-a") }
	Z.Dynamic["ls"] = func() string { return Z.Out("ls", "-l", "-h") }
}

func main() {

	// remove log prefixes
	log.SetFlags(0)

	// provide panic trace
	Z.AllowPanic = true

	Cmd.Run()
}

var Cmd = &Z.Cmd{
	Name:        `blackbox`,
	Usage:       `COMMAND [args]`,
	Summary:     help.S(_blackbox),
	Description: help.D(_blackbox),
	Copyright:   `Copyright 2024 Micah Nadler`,
	Version:     `v0.1.0`,
	License:     `Apache-2.0`,
	Source:      `git@github.com:tr00datp00nar/blackbox.git`,
	Issues:      `github.com/tr00datp00nar/blackbox/issues`,

	Commands: []*Z.Cmd{
		help.Cmd,
		good.Cmd,
		list.Cmd,
	},

	Shortcuts: Z.ArgMap{
		`avail`: {`list`, `available`},
	},
}
