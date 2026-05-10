package cmd

import (
	"fmt"
	"slices"

	"github.com/mealies/tmz/pkg/utils"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <timezone>",
	Short: "Add a timezone to the config file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		timezone := args[0]

		countries, err := utils.LoadCountries()
		if err != nil {
			return fmt.Errorf("could not load timezone data: %w", err)
		}
		validTimezones := countries["ALL"]

		if !slices.Contains(validTimezones, timezone) {
			return fmt.Errorf("invalid timezone: %s", timezone)
		}

		config, err := utils.LoadConfig()
		if err != nil {
			return err
		}

		if slices.Contains(config.Timezones, timezone) {
			fmt.Printf("Timezone %s already exists in config\n", timezone)
			return nil
		}

		config.Timezones = append(config.Timezones, timezone)
		err = utils.SaveConfig(config)
		if err != nil {
			return err
		}

		fmt.Printf("Added %s to config\n", timezone)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
