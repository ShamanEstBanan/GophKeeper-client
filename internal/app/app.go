package app

import (
	"GophKeeper-client/internal/command"
	"GophKeeper-client/internal/config"
	"github.com/urfave/cli/v2"
)

func New() *cli.App {
	config.GetConfig()
	app := &cli.App{
		Commands: command.KeeperCommands(),
	}
	return app
}
