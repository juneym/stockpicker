package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"usautoparts/stockpicker/config"
)

func main() {

	app := cli.NewApp()

	app.Name = "stockpicker"
	app.Usage = "Picks stock information from Google Finance"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Usage:  "Load confirmation from `FILE`",
			Value:  "./samples/config.json",
			EnvVar: "STOCKPICKER_CONFIG",
		},
	}

	app.Action = func(c *cli.Context) error {

		config, err := config.Load(c.String("config"))
		if err != nil {
			fmt.Printf("Error occurred while loading the configuration file %s: %s \n", c.String("config"), err.Error())
			return cli.NewExitError("Cannot load the config file", 1)
		}

		fmt.Println("Configuration: ")
		fmt.Printf("%#v \n\n", config)

		return nil
	}

	app.Run(os.Args)
}
