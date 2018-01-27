package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/sudhanshuraheja/tanker/pkg/config"
	"github.com/sudhanshuraheja/tanker/pkg/logger"
	"github.com/sudhanshuraheja/tanker/pkg/postgres"
	"github.com/sudhanshuraheja/tanker/pkg/server"
)

func main() {
	config.Init()
	logger.Init()

	logger.Infoln("Tanker is running")
	Init()
}

// Init : start the cli wrapper
func Init() *cli.App {
	app := cli.NewApp()
	app.Name = config.Name()
	app.Version = config.Version()
	app.Usage = "this service saves files and makes them available for distribution"

	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start the service",
			Action: func(c *cli.Context) error {
				server.StartAPIServer()
				return nil
			},
		},
		{
			Name:  "migrate",
			Usage: "run database migrations",
			Action: func(c *cli.Context) error {
				return postgres.RunDatabaseMigrations()
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback the latest database migration",
			Action: func(c *cli.Context) error {
				return postgres.RollbackDatabaseMigration()
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

	return app
}
