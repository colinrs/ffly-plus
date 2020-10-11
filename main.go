package main

import (
	"os"

	"ffly-plus/internal/config"
	"ffly-plus/models"
	"ffly-plus/router"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ffly-plus"
	app.Usage = "ffly-plus -c config/config.local.json"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "conf, c",
			Value: "config/config.local.json",
			Usage: "config/config.{local|dev|test|pre|prod}.json",
		},
	}

	app.Action = func(c *cli.Context) error {
		conf := c.String("conf")
		config.Init(conf)
		err := models.Database(config.Conf.MySQL)
		if err != nil {
			return err
		}
		server := router.InitRouter()
		server.GinEngine.Run(":8000")
		return nil
	}
	app.Run(os.Args)
}
