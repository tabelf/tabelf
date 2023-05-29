package actions

import (
	"context"

	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/examples/fs/ent/migrate"
	"github.com/spf13/cobra"

	"tabelf/backend/service/app"
)

func MigrateDB(cmd *cobra.Command, args []string) {
	env, err := cmd.Flags().GetString("env")
	if err != nil {
		Fatalf("get env failed: %v\n", err)
	}

	envConfig := app.LoadConfig(env)
	err = app.InitExtensions(envConfig, env)
	if err != nil {
		Panicf("init extensions failed: %v\n", err)
	}

	client := app.EntClient
	if err := client.Debug().Schema.Create(
		context.Background(),
		schema.WithHooks(func(next schema.Creator) schema.Creator {
			return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
				return next.Create(ctx, tables...)
			})
		}),
		migrate.WithForeignKeys(false),
	); err != nil {
		Fatal(err.Error())
	}
	Track("migrate successfully")
}

func init() {
	cmd := &cobra.Command{
		Use: "migrate",
		Run: MigrateDB,
	}
	rootCmd.AddCommand(cmd)
}
