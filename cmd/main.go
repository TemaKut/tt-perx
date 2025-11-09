package main

import (
	"fmt"
	"github.com/TemaKut/tt-perx/cmd/factory"
	"github.com/TemaKut/tt-perx/internal/app/config"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := cli.App{
		Flags: []cli.Flag{
			&cli.UintFlag{
				Name: CliFlagNParallelTasks,
			},
		},

		Action: func(c *cli.Context) error {
			ctx, cancel := signal.NotifyContext(c.Context, syscall.SIGINT, syscall.SIGTERM)
			defer cancel()

			cfg := config.NewConfig()

			if v := c.Uint(CliFlagNParallelTasks); v > 0 {
				cfg.Service.Math.NParallelTasks = v
			}

			_, cleanup, err := factory.InitApp(cfg)
			if cleanup != nil {
				defer cleanup()
			}

			if err != nil {
				return fmt.Errorf("error init app: %w", err)
			}

			<-ctx.Done()

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(fmt.Errorf("error run app. %w", err))
	}
}
