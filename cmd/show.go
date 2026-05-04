package cmd

import (
	"fmt"
	"slices"
	"time"

	"github.com/mealies/tmz/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	timeStr string
	all     bool
)

var showCmd = &cobra.Command{
	Use:   "show [timezone...]",
	Short: "Convert local time to specified timezones",
	RunE: func(cmd *cobra.Command, args []string) error {
		countries, err := utils.LoadCountries()
		if err != nil {
			return fmt.Errorf("could not load timezone data: %w", err)
		}
		validTimezones := countries["ALL"]

		localTime, err := utils.ParseTime(timeStr)
		if err != nil {
			return err
		}

		timezones := args
		if all {
			if len(timezones) == 0 {
				timezones = countries["ALL"]
			}
		}

		if len(timezones) == 0 && !all {
			return fmt.Errorf("at least one timezone or --all is required")
		}

		for _, tz := range timezones {
			if !slices.Contains(validTimezones, tz) {
				return fmt.Errorf("invalid timezone: %s (not found in countries.json)", tz)
			}
		}

		fmt.Printf("Local time (%s): %s\n\n", time.Local.String(), localTime.Format("2006-01-02 15:04:05 MST"))

		for _, tz := range timezones {
			loc, err := time.LoadLocation(tz)
			if err != nil {
				fmt.Printf("Error loading timezone %s: %v\n", tz, err)
				continue
			}
			converted := localTime.In(loc)
			fmt.Printf("%-20s: %s\n", tz, converted.Format("2006-01-02 15:04:05 MST"))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringVarP(&timeStr, "time", "t", "", "Local time to convert (e.g., '2006-01-02 15:04')")
	showCmd.Flags().BoolVarP(&all, "all", "a", false, "Show a set of common timezones")
}
