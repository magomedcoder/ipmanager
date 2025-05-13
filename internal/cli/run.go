package cli

import (
	"github.com/magomedcoder/ipmanager/internal/cli/handler"
	"github.com/magomedcoder/ipmanager/internal/config"
	cliV2 "github.com/urfave/cli/v2"
)

type AppProvider struct {
	Conf    *config.Config
	Migrate *handler.Migrate
}

func RunMigrate(ctx *cliV2.Context, app *AppProvider) error {
	return app.Migrate.Migrate(ctx.Context)
}

func RunTestData(ctx *cliV2.Context, app *AppProvider) error {
	return app.Migrate.TestData(ctx.Context)
}
