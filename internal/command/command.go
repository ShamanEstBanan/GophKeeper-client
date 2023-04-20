package command

import (
	"github.com/urfave/cli/v2"
)

func KeeperCommands() []*cli.Command {
	commands := []*cli.Command{
		SignCommand(),
		LoginCommand(),
		AllRecordsCommands(),
		RecordCommands(),
	}
	return commands
}
