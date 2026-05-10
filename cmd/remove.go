package cmd

import (
	"fmt"
	"slices"

	"github.com/mealies/tmz/pkg/utils"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <timezone>",
	Short: "Remove a timezone from the config file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		timezone := args[0]

		config, err := utils.LoadConfig()
		if err != nil {
			return err
		}

		if !slices.Contains(config.Timezones, timezone) {
			fmt.Printf("Timezone %s not found in config\n", timezone)
			return nil
		}

		config.Timezones = slices.DeleteFunc(config.Timezones, func(tz string) bool {
			return tz == timezone
		})

		err = utils.SaveConfig(config)
		if err != nil {
			return err
		}

		fmt.Printf("Removed %s from config\n", timezone)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
