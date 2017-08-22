package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	// App meta
	app.Name = Name
	app.Version = Version
	app.Authors = []cli.Author{
		cli.Author{
			Name:  AuthorName,
			Email: AuthorEmail,
		},
	}
	app.Usage = Usage
	app.UsageText = fmt.Sprintf("%s [options] servername", Name)

	// Flags
	var (
		srvName    string
		hostName   string
		port       string
		delay      int
		isInsecure bool
		isSilent   bool
	)

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "host, H",
			Usage:       "target host for scanning",
			Destination: &hostName,
		},
		cli.StringFlag{
			Name:        "port, p",
			Usage:       "ssl port",
			Value:       "443",
			Destination: &port,
		},
		cli.IntFlag{
			Name:        "delay, d",
			Usage:       "delay of expire",
			Value:       0,
			Destination: &delay,
		},
		cli.BoolFlag{
			Name:        "insecure, k",
			Usage:       "permit insecure cert",
			Destination: &isInsecure,
		},
		cli.BoolFlag{
			Name:        "silent, s",
			Usage:       "enable silent mode",
			Destination: &isSilent,
		},
	}

	// Set HelpTemplate
	cli.AppHelpTemplate = HelpTemplate

	// Main action
	app.Action = func(c *cli.Context) error {

		// Args
		if c.NArg() != 1 {
			cli.ShowAppHelp(c)
			return nil
		}
		srvName = c.Args().Get(0)

		if hostName == "" {
			hostName = srvName
		}

		// Dial
		expire, err := Dial(srvName, hostName, port, isInsecure)
		if err != nil {
			if !isSilent {
				fmt.Println(err)
			}
			return cli.NewExitError("", ExitError)
		}

		if !isSilent {
			fmt.Println(expire)
		}

		// Check expire
		limit := time.Now().AddDate(0, 0, delay).UTC()
		if !expire.After(limit) {
			return cli.NewExitError("", ExitExpire)
		}

		return nil
	}

	app.Run(os.Args)
}
