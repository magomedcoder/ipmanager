package cli

import (
	"github.com/magomedcoder/ipmanager/internal/app/di"
	cliV2 "github.com/urfave/cli/v2"
)

func RunMigrate(ctx *cliV2.Context, app *di.CliProvider) error {
	return app.Migrate.Migrate(ctx.Context)
}
