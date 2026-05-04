package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/mealies/tmz/pkg/utils"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <abbreviation> [time]",
	Short: "Get current time in a timezone by its abbreviation",
	Long: `Get the current or specified time for an abbreviated timezone.
   
   tmz get pst 
   
To see what a specified time would be at the specified abbreviated timezone 
   
   tmz get pst 10:30
   tmz get pst "2026-06-01 14:30"

If the abbreviation has multiple options, you will get a prompt to pick from the matching zones
   
   Multiple timezones found for CST:
   [1] (UTC+04:00) Yerevan (Asia/Yerevan)
   [2] (UTC+08:00) Beijing, Chongqing, Hong Kong, Urumqi (Asia/Hong_Kong, Asia/Macau, Asia/Shanghai)
`,
	Args: cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		abbrInput := strings.ToUpper(args[0])
		timeInput := ""
		if len(args) > 1 {
			timeInput = args[1]
		}

		abbrs, err := utils.LoadTimezoneAbbreviations()
		if err != nil {
			return fmt.Errorf("could not load timezone abbreviations: %w", err)
		}

		var matches []utils.TimezoneAbbreviation
		for _, z := range abbrs.Zones {
			if strings.ToUpper(z.Abbr) == abbrInput {
				matches = append(matches, z)
			}
		}

		if len(matches) == 0 {
			return fmt.Errorf("no timezones found for abbreviation: %s", abbrInput)
		}

		var selected utils.TimezoneAbbreviation
		if len(matches) > 1 {
			fmt.Printf("Multiple timezones found for %s:\n", abbrInput)
			for i, m := range matches {
				fmt.Printf("[%d] %s (%s)\n", i+1, m.Text, strings.Join(m.UTC, ", "))
			}
			fmt.Print("Select one (number): ")
			var choice int
			_, err := fmt.Scanln(&choice)
			if err != nil || choice < 1 || choice > len(matches) {
				return fmt.Errorf("invalid selection")
			}
			selected = matches[choice-1]
		} else {
			selected = matches[0]
		}

		if len(selected.UTC) == 0 {
			return fmt.Errorf("no IANA timezone identifier available for %s", selected.Value)
		}

		// Use the first UTC entry as the representative IANA zone
		ianaZone := selected.UTC[0]
		loc, err := time.LoadLocation(ianaZone)
		if err != nil {
			return fmt.Errorf("could not load timezone %s: %w", ianaZone, err)
		}

		localTime, err := utils.ParseTime(timeInput)
		if err != nil {
			return err
		}

		converted := localTime.In(loc)
		fmt.Printf("\nSelected: %s\n", selected.Text)
		fmt.Printf("Time (%s): %s\n", ianaZone, converted.Format("2006-01-02 15:04:05 MST"))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
