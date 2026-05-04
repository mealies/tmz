# tmz
A Go CLI App to do timezone and time calculations

## Installation

To install `tmz`, ensure you have Go installed on your system, then run:

```bash
go install github.com/mealies/tmz@latest
```

Alternatively, you can clone the repository and build it locally:

```bash
git clone https://github.com/mealies/tmz.git
cd tmz
go build -o tmz main.go
```

## Usage

### Show Command

The `show` command displays the time in one or more specified timezones with the option of specifying a specific datetime.

```bash
# Show current time in New York and London
tmz show America/New_York Europe/London

# Show a specific time in Tokyo
tmz show Asia/Tokyo --time "2023-10-27 15:00"

# Show time in all major timezones
tmz show --all
```

**Flags:**
- `-t, --time string`: Local time to convert (formats: `YYYY-MM-DD HH:MM:SS` or `HH:MM`).
- `-a, --all`: Show time in all timezones defined in the application's database.

### Get Command

The `get` command displays the time for a specified timezone abbreviation with the option of specifying a specific datetime.

```bash
# Show current time in New York and London
tmz get gmt

# Convert a specific time to that timezone]
tmz get pst 15:43

# Convert a specific datetime for Los Angeles
tmz get pst "2023-10-27 15:00"

```