package main

import (
	"github.com/magomedcoder/ipmanager/internal/app"
	"github.com/magomedcoder/ipmanager/internal/app/di"
	"github.com/magomedcoder/ipmanager/internal/config"
	"github.com/magomedcoder/ipmanager/internal/server"
	cliV2 "github.com/urfave/cli/v2"
)

func main() {
	_app := app.NewApp(&cliV2.App{
		Name: "IP Manager",
		Flags: []cliV2.Flag{
			&cliV2.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "/etc/ipmanager/ipmanager.yaml",
				Usage:       "IP Manager",
				DefaultText: "/etc/ipmanager/ipmanager.yaml",
			},
		},
	})
	_app.Register(app.Command{
		Name:  "run",
		Usage: "Run server",
		Action: func(ctx *cliV2.Context, conf *config.Config) error {
			return server.Run(di.NewGrpcInjector(conf))
		},
	})
	_app.Run()
}
