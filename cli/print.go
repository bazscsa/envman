package cli

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/bitrise-io/envman/envman"
	"github.com/codegangsta/cli"
)

func printEnvs() error {
	environments, err := envman.ReadEnvs(envman.CurrentEnvStoreFilePath)
	if err != nil {
		return err
	}

	if len(environments) == 0 {
		log.Info("[ENVMAN] - Empty envstore")
	} else {
		for _, env := range environments {
			key, value, err := env.GetKeyValuePair()
			if err != nil {
				return err
			}

			options, err := env.GetOptions()
			if err != nil {
				return err
			}

			envString := "- " + key + ": " + value
			fmt.Println(envString)
			if !*options.IsExpand {
				expandString := "  " + "isExpand" + ": " + "false"
				fmt.Println(expandString)
			}
		}
	}

	return nil
}

func print(c *cli.Context) {
	log.Debugln("[ENVMAN] - Work path:", envman.CurrentEnvStoreFilePath)

	if err := printEnvs(); err != nil {
		log.Fatal("[ENVMAN] - Failed to print:", err)
	}
}
