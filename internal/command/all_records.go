package command

import (
	"GophKeeper-client/internal/client"
	"fmt"
	"github.com/urfave/cli/v2"
)

func AllRecordsCommands() *cli.Command {
	c := &cli.Command{
		Name:    "records",
		Aliases: []string{"a"},
		Usage:   "complete a task on the list",
		Subcommands: []*cli.Command{
			{
				Name:  "all",
				Usage: "add a new template",
				Action: func(cCtx *cli.Context) error {
					err := client.GetAllRecords()
					if err != nil {
						return fmt.Errorf("can't get records. Try again. Error: %w", err)
					}
					return nil
				},
			},
			{
				Name:  "type",
				Usage: "add a new template",
				Action: func(cCtx *cli.Context) error {
					recordType := cCtx.Args().First()
					err := client.GetAllRecordsByType(recordType)
					if err != nil {
						return fmt.Errorf("can't get records. Try again. Error: %w", err)
					}
					return nil
				},
			},
		},
	}
	return c
}
