package main

import (
	"GophKeeper-client/internal/client"
	"GophKeeper-client/internal/config"
	"GophKeeper-client/internal/entity"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	config.GetConfig()
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "sign",
				Aliases: []string{"s"},
				Usage:   "registration in service ",
				Action: func(cCtx *cli.Context) error {
					l := cCtx.Args().First()
					p := cCtx.Args().Get(1)
					return client.SignUp(l, p)
				},
			},
			{
				Name:    "login",
				Aliases: []string{"l"},
				Usage:   "login in system",
				Action: func(cCtx *cli.Context) error {
					l := cCtx.Args().First()
					p := cCtx.Args().Get(1)
					client.Login(l, p)
					return nil
				},
			},
			{
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
			},
			{
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

									cardInfo := entity.RecordPassword{}
									fmt.Printf("Enter the login: ")
									_, err = fmt.Scanln(&cardInfo.Login)
									fmt.Printf("Enter the password: ")
									_, err = fmt.Scanln(&cardInfo.Password)

									recordType := "password"
									err = client.Create(name, recordType, recordType)
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

									recordType := "bank_card"

									err = client.Create(name, recordType, cardInfo)
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

									recordType := "text_string"
									err = client.Create(name, recordType, data)
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
									var name string
									fmt.Printf("Enter the name of record: ")
									_, err := fmt.Scanln(&name)

									data := entity.RecordByteString{}
									fmt.Printf("Enter the path to byte-file:")
									_, err = fmt.Scanln(&data.Data)

									recordType := "byte_string"
									err = client.Create(name, recordType, data)
									if err != nil {
										return err
									}
									return nil
								},
							},
						},
						Action: func(cCtx *cli.Context) error {
							fmt.Println("Choose type of record: password, some text string, some byte string, back card info", cCtx.Args().First())
							return nil
						},
					},
					{
						Name:  "view",
						Usage: "add a new template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("new task template: ", cCtx.Args().First())
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
							fmt.Println("removed task template: ", cCtx.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
