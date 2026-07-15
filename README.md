<p align="center">
  <img src="assets/banner.svg" alt="go-garmin" width="100%">
</p>

# go-garmin

A Go client library, CLI, and MCP server for the Garmin Connect API.

Use it as:

- a **Go library** in your own programs
- a **CLI** (`garmin …`) that prints JSON
- an **MCP server** (`garmin mcp`) so LLM assistants can query and update your Garmin data

## Features

- OAuth login with MFA support and automatic token refresh
- Declarative endpoint registry: one definition drives CLI commands + MCP tools
- Broad Connect coverage: sleep, wellness, activities, training metrics, workouts, summary stats, badges, blood pressure, lifestyle, and more
- VCR-backed integration fixtures for reliable tests

## Installation

### Prerequisites

- Go 1.25 or later ([install Go](https://go.dev/doc/install))
- `$GOPATH/bin` (usually `~/go/bin`) on your `PATH`

### Install CLI

```bash
go install github.com/shotah/go-garmin/cmd/garmin@latest
garmin --version
```

### Build from source

```bash
git clone https://github.com/shotah/go-garmin.git
cd go-garmin
go build -o bin/garmin ./cmd/garmin
./bin/garmin --help
```

## Quick start

```bash
# 1. Interactive login (email / password / MFA) → saves session.json
garmin login

# 2. Smoke test
garmin sleep
garmin metrics readiness

# 3. Optional: expose the same session to an LLM via MCP
garmin mcp
```

Session file location:

| OS | Path |
|----|------|
| Linux / macOS | `~/.config/garmin/session.json` (or `$XDG_CONFIG_HOME/garmin/session.json`) |
| Windows | `%AppData%\garmin\session.json` |

Access tokens refresh automatically on API calls. Rotated tokens are written back to `session.json`. If the refresh token itself is revoked, run `garmin logout` then `garmin login` again.

## CLI usage

All commands print JSON (suitable for `jq`, scripts, and piping).

### Authentication

```bash
garmin login     # interactive email / password / MFA
garmin logout    # delete saved session
```

### Commands

```bash
# Sleep
garmin sleep [date]

# Wellness
garmin wellness stress [date]
garmin wellness body-battery [date]
garmin wellness heart-rate [date]
garmin wellness spo2 [date]
garmin wellness respiration [date]
garmin wellness intensity-minutes [date]
garmin wellness events [date]
garmin wellness sleep [date] [--display-name=...]
garmin wellness steps [date] [--display-name=...]
garmin wellness floors [date]
garmin wellness body-battery-reports [--start=YYYY-MM-DD] [--end=YYYY-MM-DD]
garmin wellness sleep-score [--start=YYYY-MM-DD] [--end=YYYY-MM-DD]

# Daily totals (user summary)
garmin summary daily [date] [--display-name=...]
garmin summary hydration [date]
garmin summary log-hydration --json='{"calendarDate":"2026-07-15","timestampLocal":"2026-07-15T12:00:00.000","valueInML":250}'
garmin summary steps-daily [--start=YYYY-MM-DD] [--end=YYYY-MM-DD]
garmin summary steps-weekly [end] [--weeks=4]
garmin summary stress-daily [--start=YYYY-MM-DD] [--end=YYYY-MM-DD]
garmin summary stress-weekly [end] [--weeks=4]
garmin summary hydration-stats [--start=YYYY-MM-DD] [--end=YYYY-MM-DD]
garmin summary im-daily [--start=YYYY-MM-DD] [--end=YYYY-MM-DD]
garmin summary im-weekly [--start=YYYY-MM-DD] [--end=YYYY-MM-DD]

# Activities
garmin activities list [--start=0] [--limit=20]
garmin activities get <activity-id>
garmin activities types
garmin activities splits <activity-id>
garmin activities weather <activity-id>
garmin activities details <activity-id>
garmin activities hr-zones <activity-id>
garmin activities power-zones <activity-id>
garmin activities exercise-sets <activity-id>

# Weight and HRV
garmin weight daily [date]
garmin weight range --start=YYYY-MM-DD --end=YYYY-MM-DD
garmin hrv daily [date]
garmin hrv range --start=YYYY-MM-DD --end=YYYY-MM-DD

# Training metrics
garmin metrics readiness [date]
garmin metrics vo2max [date]
garmin metrics endurance [date]
garmin metrics hill [date]
garmin metrics hill-stats --start=YYYY-MM-DD --end=YYYY-MM-DD
garmin metrics training-status [date]
garmin metrics load-balance [date]
garmin metrics acclimation [date]
garmin metrics race-predictions [display-name]
garmin metrics race-predictions-daily [date]
garmin metrics race-predictions-monthly [date]

# Fitness age / fitness stats
garmin fitnessage daily [date]
garmin fitnessage stats --start=YYYY-MM-DD --end=YYYY-MM-DD
garmin fitnessstats get [--start=YYYY-MM-DD] [--end=YYYY-MM-DD] [--aggregation=weekly] [--metrics=calories,distance,duration]
garmin fitnessstats activities [--start=YYYY-MM-DD] [--end=YYYY-MM-DD] [--activity_type=running] [--metrics=name,startLocal,activityType]

# Biometrics / devices / profile
garmin biometric lactate-threshold
garmin biometric ftp
garmin biometric hr-zones
garmin biometric power-weight [date]
garmin devices list
garmin devices settings <device-id>
garmin profile social
garmin profile settings
garmin profile display

# Workouts
garmin workouts list [--start=0] [--limit=20]
garmin workouts get <workout-id>
garmin workouts create --file=workout.json
garmin workouts create --json='{"workoutName": "..."}'
cat workout.json | garmin workouts create
garmin workouts update <workout-id> --file=workout.json
garmin workouts delete <workout-id>
garmin workouts schedule <workout-id> <date>
garmin workouts unschedule <schedule-id>

# Exercise library (strength workouts)
garmin exercises categories
garmin exercises muscles
garmin exercises equipment
garmin exercises list [--category=BENCH_PRESS] [--muscle=CHEST] [--equipment=DUMBBELL]
garmin exercises get <exercise-key>

# Calendar (month is 0-indexed: January=0)
garmin calendar get --year=2026 [--month=0] [--day=28] [--start=1]

# Personal records / training plans
garmin records list
garmin plans list
garmin plans phased <plan-id>
garmin plans adaptive <plan-id>

# Badges and challenges
garmin badges earned
garmin badges available
garmin badges challenges-completed
garmin badges challenges-available
garmin badges challenges-open
garmin badges virtual
garmin badges adhoc

# Blood pressure
garmin bp range --start=YYYY-MM-DD --end=YYYY-MM-DD
garmin bp log --json='{"systolic":120,"diastolic":80,...}'
garmin bp delete --date=YYYY-MM-DD --version=1

# Periodic health
garmin health day [date]
garmin health calendar --start=YYYY-MM-DD --end=YYYY-MM-DD
garmin health pregnancy

# Lifestyle logging
garmin lifestyle daily [date]
garmin lifestyle create-behaviour --json='{"name":"Went drinking",...}'
```

## MCP server (LLM integration)

`garmin mcp` starts a [Model Context Protocol](https://modelcontextprotocol.io/) server over stdio. Hosts like Claude Code, Claude Desktop, and Cursor spawn that process and let the model call tools.

### What the AI actually reads

The model does **not** automatically read this README, `ENDPOINTS.md`, or the Go source.

When the MCP host connects, it asks the server for its tool list. For each registered endpoint, go-garmin exposes:

| What the model sees | Where it comes from |
|---------------------|---------------------|
| **Tool name** | `MCPTool` in `endpoint/definitions/*.go` (e.g. `get_sleep`) |
| **Tool description** | the endpoint `Long` string |
| **Argument schema** | each `Param` name/type/required + `Description` |
| **JSON body hints** (write tools) | `BodyConfig.Description` and `BodyConfig.Example` |
| **Tool results** | pretty-printed JSON returned by the Garmin API handler |

Flow:

1. Host starts `garmin mcp` and loads `session.json`.
2. Host sends `tools/list` → model gets names + descriptions + schemas (~97 tools).
3. Model picks a tool and arguments (e.g. `get_sleep` with `date=2026-07-14`).
4. Host sends `tools/call` → handler hits Garmin Connect → result comes back as JSON text.
5. Model reasons over that JSON to answer you.

Useful implications:

- Better `Long` / param descriptions in endpoint definitions = better tool use.
- `get_current_date` is a local helper (no Garmin call) so the model can resolve “today” / “yesterday”.
- Binary downloads (`RawOutput`) are CLI-only and are **not** registered as MCP tools.
- Auth is invisible to the model: it just gets errors if you are not logged in.

### Prerequisites

1. Install the CLI (see [Installation](#installation))
2. Login once: `garmin login`
3. Confirm: `garmin sleep` returns data

### Claude Code

Add to `~/.claude.json` (global) or `.claude/settings.json` (project):

```json
{
  "mcpServers": {
    "garmin": {
      "command": "garmin",
      "args": ["mcp"]
    }
  }
}
```

### Claude Desktop

Add to `~/Library/Application Support/Claude/claude_desktop_config.json` (macOS) or `%APPDATA%\Claude\claude_desktop_config.json` (Windows):

```json
{
  "mcpServers": {
    "garmin": {
      "command": "garmin",
      "args": ["mcp"]
    }
  }
}
```

### Cursor

Add to Cursor MCP settings:

```json
{
  "garmin": {
    "command": "garmin",
    "args": ["mcp"]
  }
}
```

### Troubleshooting

1. Verify PATH: `which garmin` / `where garmin`
2. Confirm login: `garmin sleep`
3. Use an absolute binary path in MCP config if the host cannot find `garmin`
4. Re-login after Garmin invalidates the refresh token: `garmin logout && garmin login`

### Example prompts

- "How did I sleep last night?"
- "What's my training readiness today?"
- "Show my personal records and hill score trend this week"
- "Log 250ml of water"
- "Create a custom lifestyle behaviour called Went drinking"
- "Create a 45-minute threshold interval workout and schedule it for tomorrow"

### Available tools

The MCP server exposes **97 tools** generated from the endpoint registry:

| Category | Tools |
|----------|-------|
| Utility | `get_current_date` |
| Sleep | `get_sleep` |
| Wellness | `get_stress`, `get_body_battery`, `get_heart_rate`, `get_spo2`, `get_respiration`, `get_intensity_minutes`, `get_daily_events`, `get_wellness_sleep`, `get_steps_chart`, `get_floors`, `get_body_battery_reports`, `get_sleep_score_stats` |
| User summary | `get_daily_user_summary`, `get_daily_hydration`, `log_hydration`, `get_steps_daily_stats`, `get_steps_weekly_stats`, `get_stress_daily_stats`, `get_stress_weekly_stats`, `get_hydration_stats`, `get_intensity_minutes_daily_stats`, `get_intensity_minutes_weekly_stats` |
| Activity | `list_activities`, `get_activity`, `get_activity_types`, `get_activity_splits`, `get_activity_weather`, `get_activity_details`, `get_activity_hr_zones`, `get_activity_power_zones`, `get_activity_exercise_sets`, `get_activity_typed_splits`, `get_activity_split_summaries`, `get_activity_gear` |
| Weight / HRV | `get_weight`, `get_hrv` |
| Metrics | `get_training_readiness`, `get_training_status`, `get_vo2max`, `get_endurance_score`, `get_hill_score`, `get_hill_score_stats`, `get_training_load_balance`, `get_heat_altitude_acclimation`, `get_race_predictions`, `get_race_predictions_daily`, `get_race_predictions_monthly` |
| Fitness age / stats | `get_fitness_age`, `get_fitness_age_stats`, `get_fitness_stats`, `get_fitness_stats_activities` |
| Biometric | `get_lactate_threshold`, `get_cycling_ftp`, `get_heart_rate_zones`, `get_power_to_weight` |
| Devices / profile | `list_devices`, `get_device_settings`, `get_social_profile`, `get_user_settings`, `get_profile_settings` |
| Workouts | `list_workouts`, `get_workout`, `create_workout`, `update_workout`, `delete_workout`, `schedule_workout`, `unschedule_workout` |
| Exercises | `list_exercise_categories`, `list_muscle_groups`, `list_equipment_types`, `list_exercises`, `get_exercise` |
| Calendar / courses | `get_calendar`, `list_courses`, `get_course`, `delete_course` |
| Records / plans | `get_personal_records`, `list_training_plans`, `get_training_plan_phased`, `get_training_plan_adaptive` |
| Badges | `get_earned_badges`, `get_available_badges`, `get_completed_badge_challenges`, `get_available_badge_challenges`, `get_non_completed_badge_challenges`, `get_virtual_challenges_in_progress`, `get_adhoc_historical_challenges` |
| Blood pressure | `get_blood_pressure_range`, `log_blood_pressure`, `delete_blood_pressure` |
| Periodic health | `get_menstrual_day_view`, `get_menstrual_calendar`, `get_pregnancy_snapshot` |
| Lifestyle | `get_daily_lifestyle_log`, `create_lifestyle_behaviour` |

### LLM-powered workout creation

Ask in natural language; the model uses the exercise library (~1,794 exercises) and workout tools to build Connect-compatible workouts.

**Running**

> "Create a 45-minute interval workout with 5-minute warmup, 6x3min at threshold with 2min recovery, and cooldown"

**Strength**

> "Create a push day: bench press 4x8, overhead press 3x10, tricep dips 3x12, 90s rest"

**Planning**

> "Look at my recent activities and create a recovery workout for tomorrow"

#### Workout sport types

| Sport | ID | Features |
|-------|----|----------|
| Running | 1 | Pace zones, HR zones, distance/time |
| Cycling | 2 | Power zones, cadence, distance/time |
| Swimming | 4 | Stroke types, equipment, pool length |
| Strength | 5 | Exercise library, reps, sets, rest |
| Cardio | 6 | HR zones, time targets |
| Yoga | 7 | Time-based flows |
| Pilates | 8 | Time-based sequences |
| HIIT | 9 | Intervals, work/rest |

## Library usage

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/shotah/go-garmin/garmin"
)

func main() {
	client := garmin.New(garmin.Options{})

	if err := client.Login(context.Background(), "email", "password"); err != nil {
		panic(err)
	}

	sleep, err := client.Sleep.GetDaily(context.Background(), time.Now())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Sleep score: %d\n", sleep.SleepScores.Overall.Value)
}
```

Common service entry points on `*garmin.Client`:

`Sleep`, `Wellness`, `Activities`, `Weight`, `HRV`, `Metrics`, `FitnessAge`, `FitnessStats`, `Biometric`, `Devices`, `UserProfile`, `Workouts`, `Exercises`, `Calendar`, `Courses`, `UserSummary`, `PersonalRecords`, `TrainingPlans`, `Badges`, `BloodPressure`, `PeriodicHealth`, `Lifestyle`

## Architecture

Endpoints are defined once under `endpoint/definitions/` and registered in `endpoint/definitions/register.go`. From that registry the project generates:

- CLI commands (`endpoint.CLIGenerator`)
- MCP tools (`endpoint.MCPGenerator`)
- Completeness checks (`endpoint.Validator`)

See [AGENTS.md](AGENTS.md) for the add-an-endpoint workflow, and [ENDPOINTS.md](ENDPOINTS.md) for the Garmin API coverage checklist.

## Development

```bash
make help                 # list targets
make tools                # install goimports-reviser + golangci-lint v2
make install-hooks        # git pre-commit: autofix + lint + validate + test
make check                # same checks as the pre-commit hook
make validate-endpoints   # endpoint registry completeness
make cli                  # build ./bin/garmin

# Fixture recording + VCR integration tests (auth required first)
make auth                 # interactive login → settings.json (required)
make fixtures             # record/update VCR cassettes
make fixtures CASSETTE=metrics
make test-integration     # go test -tags=integration (needs auth + fixtures)
```

## Releasing

Go modules are **not** published to GitHub Packages. Publishing is a git tag:

```bash
git tag v0.1.0
git push origin v0.1.0
```

That makes the library installable as:

```bash
go get github.com/shotah/go-garmin@v0.1.0
go install github.com/shotah/go-garmin/cmd/garmin@v0.1.0
```

Pushing a `v*` tag also runs **GoReleaser**, which attaches multi-platform `garmin` CLI binaries to a [GitHub Release](https://github.com/shotah/go-garmin/releases) (linux/mac/windows). That’s the right place for downloadable binaries; GitHub Packages is a poor fit for Go source modules.

CI runs on every PR/push to `main` (build, lint, endpoint validation, non-integration tests).

## License

MIT
