package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/colinrs/ffly-plus/models"

	"github.com/colinrs/ffly-plus/internal/config"
	"github.com/colinrs/ffly-plus/internal/sentinelm"
	"github.com/colinrs/ffly-plus/internal/version"
	"github.com/colinrs/ffly-plus/router"
	"github.com/colinrs/ffly-plus/rpc"
	serverGin "github.com/colinrs/pkgx/server/gin"

	"github.com/arl/statsviz"
	"github.com/urfave/cli"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io:8080
// @BasePath /api/v1
func main() {
	app := cli.NewApp()
	app.Name = "ffly-plus"
	app.Usage = "ffly-plus -c config/config.local.json"
	printVersion := false
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "conf, c",
			Value: "config/config.local.json",
			Usage: "config/config.{local|dev|test|pre|prod}.json",
		},
		cli.BoolFlag{
			Name:        "version, v",
			Required:    false,
			Usage:       "-v",
			Destination: &printVersion,
		},
	}

	app.Action = func(c *cli.Context) error {
		if printVersion {
			fmt.Printf("{%#v}", version.Get())
			return nil
		}

		conf := c.String("conf")
		config.Init(conf)

		err := sentinelm.InitSentinelByCustom()
		if err != nil {
			return err
		}
		err = models.Database(config.Conf.MySQL)
		if err != nil {
			return err
		}
		if err = InitValidator(); err != nil {
			return err
		}
		go func() {
			err = rpc.InitRPCService()
			if err != nil {
				panic(err)
			}
		}()

		server := router.InitRouter()
		go runMonti()
		server.GinEngine.Run(":8000")
		return nil
	}
	app.Run(os.Args)
}

func runMonti() {
	// Register statsviz handlers on the default serve mux.
	statsviz.RegisterDefault()
	http.ListenAndServe(":8001", nil)
}

// InitValidator ...
func InitValidator() error {
	var validateFuns []*serverGin.ValidateFuncs
	validateFuns = append(validateFuns, &serverGin.ValidateFuncs{
		TagName: "ccless",
		Fn:      serverGin.IsLess,
		Message: "{0} not less:{1}",
	})

	return serverGin.InitValidator(validateFuns)
}
