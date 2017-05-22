package main

import (
	"fmt"
	"os"

	"github.com/leeola/fixity"
	"github.com/leeola/fixity/autoload"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "fiki"
	app.Usage = "a wiki implemented ontop of Fixity"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Value:  "~/.config/fixity/config.toml",
			Usage:  "load config from `PATH`",
			EnvVar: "FIXI_CONFIG",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "blob",
			ArgsUsage: "HASH",
			Aliases:   []string{"b"},
			Usage:     "inspect a raw blob from HASH",
			Action:    BlobCmd,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// loadFixity expands the configPath and loads fixity.
func loadFixity(ctx *cli.Context) (fixity.Fixity, error) {
	configPath := ctx.GlobalString("config")

	configPath, err := homedir.Expand(configPath)
	if err != nil {
		return nil, err
	}

	return autoload.LoadFixity(configPath)
}
