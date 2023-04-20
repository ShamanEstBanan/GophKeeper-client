package command

import (
	"GophKeeper-client/internal/client"
	"github.com/urfave/cli/v2"
)

func SignCommand() *cli.Command {
	c := &cli.Command{
		Name:    "sign",
		Aliases: []string{"s"},
		Usage:   "registration in service ",
		Action: func(cCtx *cli.Context) error {
			l := cCtx.Args().First()
			p := cCtx.Args().Get(1)
			return client.SignUp(l, p)
		},
	}
	return c
}

func LoginCommand() *cli.Command {
	c := &cli.Command{
		Name:    "login",
		Aliases: []string{"l"},
		Usage:   "login in system",
		Action: func(cCtx *cli.Context) error {
			l := cCtx.Args().First()
			p := cCtx.Args().Get(1)
			client.Login(l, p)
			return nil
		},
	}
	return c
}
