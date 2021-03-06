package cli

import "github.com/codegangsta/cli"

var (
	commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Create an empty .envstore.yml into the current working directory, or to the path specified by the --path flag.",
			Action:  initEnvStore,
			Flags: []cli.Flag{
				flClear,
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add new, or update an exist environment variable.",
			Action:  add,
			Flags: []cli.Flag{
				flKey,
				flValue,
				flValueFile,
				flNoExpand,
				flAppend,
			},
		},
		{
			Name:    "clear",
			Aliases: []string{"c"},
			Usage:   "Clear the envstore.",
			Action:  clear,
		},
		{
			Name:    "print",
			Aliases: []string{"p"},
			Usage:   "Print out the environment variables in envstore.",
			Action:  print,
		},
		{
			Name:            "run",
			Aliases:         []string{"r"},
			Usage:           "Run the specified command with the environment variables stored in the envstore.",
			SkipFlagParsing: true,
			Action:          run,
		},
	}
)
