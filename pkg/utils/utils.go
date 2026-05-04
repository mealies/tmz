package utils

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mealies/tmz/pkg/data"
)

type Countries map[string][]string
type TimezoneAbbreviation struct {
	Value  string   `json:"value"`
	Abbr   string   `json:"abbr"`
	Offset float64  `json:"offset"`
	IsDST  bool     `json:"isdst"`
	Text   string   `json:"text"`
	UTC    []string `json:"utc"`
}

type TimezoneAbbreviations struct {
	Zones []TimezoneAbbreviation `json:"zones"`
}

func LoadCountries() (Countries, error) {
	b, err := data.DataFS.ReadFile("countries.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read countries.json: %w", err)
	}

	var countries Countries
	if err := json.Unmarshal(b, &countries); err != nil {
		return nil, fmt.Errorf("failed to unmarshal countries.json: %w", err)
	}

	return countries, nil
}

func LoadTimezoneAbbreviations() (*TimezoneAbbreviations, error) {
	b, err := data.DataFS.ReadFile("tz-abbr.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read tz-abbr.json: %w", err)
	}

	if len(b) == 0 {
		return &TimezoneAbbreviations{}, nil
	}

	var abbr TimezoneAbbreviations
	if err := json.Unmarshal(b, &abbr); err != nil {
		return nil, fmt.Errorf("failed to unmarshal tz-abbr.json: %w", err)
	}

	return &abbr, nil
}

func ParseTime(timeStr string) (time.Time, error) {
	if timeStr == "" {
		return time.Now(), nil
	}

	formats := []string{
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"15:04:05",
		"15:04",
	}

	var lastErr error
	for _, f := range formats {
		t, err := time.ParseInLocation(f, timeStr, time.Local)
		if err == nil {
			if f == "15:04:05" || f == "15:04" {
				now := time.Now()
				t = time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), 0, time.Local)
			}
			return t, nil
		}
		lastErr = err
	}
	return time.Time{}, fmt.Errorf("invalid time format: %w", lastErr)
}
