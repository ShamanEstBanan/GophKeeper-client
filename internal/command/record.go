package command

import (
	"GophKeeper-client/internal/client"
	"GophKeeper-client/internal/entity"
	"fmt"
	"github.com/urfave/cli/v2"
)

func RecordCommands() *cli.Command {
	c := &cli.Command{
		Name:    "record",
		Aliases: []string{"r"},
		Usage:   "options for task templates",
		Subcommands: []*cli.Command{
			{
				Name:  "add",
				Usage: "add a new template",
				Subcommands: []*cli.Command{
					{

						Name:  "password",
						Usage: "add a new password",
						Action: func(cCtx *cli.Context) error {
							var name string
							fmt.Printf("Enter the name of record: ")
							_, err := fmt.Scanln(&name)

							passwdInfo := entity.RecordPassword{}
							fmt.Printf("Enter the login: ")
							_, err = fmt.Scanln(&passwdInfo.Login)
							fmt.Printf("Enter the password: ")
							_, err = fmt.Scanln(&passwdInfo.Password)

							err = client.Create(name, entity.RecordTypePassword, passwdInfo)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{

						Name:  "bank_card",
						Usage: "add a new bank card info",
						Action: func(cCtx *cli.Context) error {

							var name string
							fmt.Printf("Enter the name of record: ")
							_, err := fmt.Scanln(&name)

							cardInfo := entity.RecordBankCard{}
							fmt.Printf("Enter the number of card: ")
							_, err = fmt.Scanln(&cardInfo.Number)
							fmt.Printf("Enter the owner's name of card: ")
							_, err = fmt.Scanln(&cardInfo.UserName)
							fmt.Printf("Enter the expire date of card: ")
							_, err = fmt.Scanln(&cardInfo.ExpiredDate)
							fmt.Printf("Enter the CVV of card: ")
							_, err = fmt.Scanln(&cardInfo.CVV)

							err = client.Create(name, entity.RecordTypeBankCard, cardInfo)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{

						Name:  "text_string",
						Usage: "add a new text string",
						Action: func(cCtx *cli.Context) error {
							var name string
							fmt.Printf("Enter the name of record: ")
							_, err := fmt.Scanln(&name)

							data := entity.RecordText{}
							fmt.Printf("Enter the text:")
							_, err = fmt.Scanln(&data.Data)

							err = client.Create(name, entity.RecordTypeTextString, data)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{

						Name:  "byte_string",
						Usage: "add a new byte string",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("Sorry, will be later..")
							return nil
						},
					},
				},
				Action: func(cCtx *cli.Context) error {
					fmt.Println("\nAdd to command type of record: password, bank_card, " +
						"text_string, byte_string\n")
					return nil
				},
			},
			{
				Name:  "view",
				Usage: "add a new template",
				Action: func(cCtx *cli.Context) error {
					var id string
					fmt.Printf("Enter ID of record: ")
					_, err := fmt.Scanln(&id)
					err = client.Get(id)
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:  "edit",
				Usage: "remove an existing template",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("removed task template: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:  "remove",
				Usage: "remove an existing template",
				Action: func(cCtx *cli.Context) error {
					var id string
					fmt.Printf("Enter ID of record to remove: ")
					_, err := fmt.Scanln(&id)
					err = client.Delete(id)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}
	return c
}
