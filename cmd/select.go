package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mealies/tmz/pkg/utils"
	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select a timezone from your config and convert time",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := utils.LoadConfig()
		if err != nil {
			return fmt.Errorf("could not load config: %w", err)
		}

		if len(config.Timezones) == 0 {
			return fmt.Errorf("no timezones saved in config. Use 'tmz add <timezone>' to add some")
		}

		fmt.Println("Saved Timezones:")
		for i, tz := range config.Timezones {
			fmt.Printf("[%d] %s\n", i+1, tz)
		}

		fmt.Print("Select a timezone (number): ")
		var choiceStr string
		fmt.Scanln(&choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil || choice < 1 || choice > len(config.Timezones) {
			return fmt.Errorf("invalid selection")
		}

		selectedTz := config.Timezones[choice-1]
		loc, err := time.LoadLocation(selectedTz)
		if err != nil {
			return fmt.Errorf("could not load timezone %s: %w", selectedTz, err)
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter local time to convert(Press Enter for current time): ")
		timeInput, _ := reader.ReadString('\n')
		timeInput = strings.TrimSpace(timeInput)

		localTime, err := utils.ParseTime(timeInput)
		if err != nil {
			return err
		}

		converted := localTime.In(loc)
		fmt.Printf("\nTime in %s: %s\n", selectedTz, converted.Format("2006-01-02 15:04:05 MST"))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
}
