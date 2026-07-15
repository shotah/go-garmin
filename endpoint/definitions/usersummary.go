package definitions

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/shotah/go-garmin/endpoint"
	"github.com/shotah/go-garmin/garmin"
)

var hydrationLogBodyConfig = &endpoint.BodyConfig{
	Type: reflect.TypeFor[garmin.HydrationLogRequest](),
	Description: `JSON object to log hydration intake. Required fields: calendarDate, timestampLocal, valueInML.

- calendarDate (string, YYYY-MM-DD): Day to attribute the intake to
- timestampLocal (string): Local timestamp, e.g. "2026-07-15T12:00:00.000"
- valueInML (number): Milliliters to add (positive) or subtract (negative)

The live Connect API expects PUT.`,
	Example: `{
  "calendarDate": "2026-07-15",
  "timestampLocal": "2026-07-15T12:00:00.000",
  "valueInML": 250
}`,
}

// UserSummaryEndpoints defines usersummary-service endpoints (daily totals, hydration, stats).
var UserSummaryEndpoints = []endpoint.Endpoint{
	{
		Name:       "GetDailyUserSummary",
		Service:    "UserSummary",
		Cassette:   "usersummary",
		Path:       "/usersummary-service/usersummary/daily/{displayName}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "date", Type: endpoint.ParamTypeDate, Required: false, Description: "Date for summary (YYYY-MM-DD, defaults to today)"},
			{Name: "display_name", Type: endpoint.ParamTypeString, Required: false, Description: "User display name (defaults to current user)"},
		},
		CLICommand:    "summary",
		CLISubcommand: "daily",
		MCPTool:       "get_daily_user_summary",
		Short:         "Get daily user summary",
		Long:          "Get Connect home-screen daily totals: steps, calories, distance, floors, intensity, stress, and body battery",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			displayName, err := client.ResolveDisplayName(ctx, args.String("display_name"))
			if err != nil {
				return nil, err
			}
			return client.UserSummary.GetDaily(ctx, displayName, args.Date("date"))
		},
	},
	{
		Name:       "GetDailyHydration",
		Service:    "UserSummary",
		Cassette:   "usersummary",
		Path:       "/usersummary-service/usersummary/hydration/daily/{date}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "date", Type: endpoint.ParamTypeDate, Required: false, Description: "Date for hydration data (YYYY-MM-DD, defaults to today)"},
		},
		CLICommand:    "summary",
		CLISubcommand: "hydration",
		MCPTool:       "get_daily_hydration",
		Short:         "Get daily hydration",
		Long:          "Get hydration intake and goal for a single day",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.UserSummary.GetHydration(ctx, args.Date("date"))
		},
	},
	{
		Name:          "LogHydration",
		Service:       "UserSummary",
		Cassette:      "none",
		Path:          "/usersummary-service/usersummary/hydration/log",
		HTTPMethod:    "PUT",
		Body:          hydrationLogBodyConfig,
		CLICommand:    "summary",
		CLISubcommand: "log-hydration",
		MCPTool:       "log_hydration",
		Short:         "Log hydration intake",
		Long:          "Log or adjust hydration intake (positive adds ml, negative subtracts). Use --file, --json, or stdin. Live API uses PUT.",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			req, ok := args.Body.(*garmin.HydrationLogRequest)
			if !ok {
				return nil, fmt.Errorf("invalid hydration log body type: %T", args.Body)
			}
			return client.UserSummary.LogHydration(ctx, req)
		},
	},
	{
		Name:       "GetStepsDailyStats",
		Service:    "UserSummary",
		Cassette:   "usersummary",
		Path:       "/usersummary-service/stats/steps/daily/{start}/{end}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "range", Type: endpoint.ParamTypeDateRange, Required: false, Description: "Date range for daily steps (defaults to last 7 days; ranges >28 days are chunked)"},
		},
		CLICommand:    "summary",
		CLISubcommand: "steps-daily",
		MCPTool:       "get_steps_daily_stats",
		Short:         "Get daily steps stats",
		Long:          "Get daily steps and distance for a date range. API limit is 28 days per call; longer ranges are automatically chunked",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			end := time.Now()
			start := end.AddDate(0, 0, -7)
			if args.HasParam("start") {
				start = args.Date("start")
			}
			if args.HasParam("end") {
				end = args.Date("end")
			}
			return client.UserSummary.GetStepsDaily(ctx, start, end)
		},
	},
	{
		Name:       "GetStepsWeeklyStats",
		Service:    "UserSummary",
		Cassette:   "usersummary",
		Path:       "/usersummary-service/stats/steps/weekly/{end}/{weeks}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "end", Type: endpoint.ParamTypeDate, Required: false, Description: "End date for weekly stats (YYYY-MM-DD, defaults to today)"},
			{Name: "weeks", Type: endpoint.ParamTypeInt, Required: false, Description: "Number of weeks (default 4)"},
		},
		CLICommand:    "summary",
		CLISubcommand: "steps-weekly",
		MCPTool:       "get_steps_weekly_stats",
		Short:         "Get weekly steps stats",
		Long:          "Get weekly aggregated steps and distance ending on the given date",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.UserSummary.GetStepsWeekly(ctx, args.Date("end"), args.IntOrDefault("weeks", 4))
		},
	},
	{
		Name:       "GetStressDailyStats",
		Service:    "UserSummary",
		Cassette:   "usersummary",
		Path:       "/usersummary-service/stats/stress/daily/{start}/{end}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "range", Type: endpoint.ParamTypeDateRange, Required: false, Description: "Date range for daily stress stats (defaults to last 7 days)"},
		},
		CLICommand:    "summary",
		CLISubcommand: "stress-daily",
		MCPTool:       "get_stress_daily_stats",
		Short:         "Get daily stress stats",
		Long:          "Get daily stress summary values for a date range",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			end := time.Now()
			start := end.AddDate(0, 0, -7)
			if args.HasParam("start") {
				start = args.Date("start")
			}
			if args.HasParam("end") {
				end = args.Date("end")
			}
			return client.UserSummary.GetStressDaily(ctx, start, end)
		},
	},
	{
		Name:       "GetStressWeeklyStats",
		Service:    "UserSummary",
		Cassette:   "usersummary",
		Path:       "/usersummary-service/stats/stress/weekly/{end}/{weeks}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "end", Type: endpoint.ParamTypeDate, Required: false, Description: "End date for weekly stress stats (YYYY-MM-DD, defaults to today)"},
			{Name: "weeks", Type: endpoint.ParamTypeInt, Required: false, Description: "Number of weeks (default 4)"},
		},
		CLICommand:    "summary",
		CLISubcommand: "stress-weekly",
		MCPTool:       "get_stress_weekly_stats",
		Short:         "Get weekly stress stats",
		Long:          "Get weekly aggregated stress values ending on the given date",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.UserSummary.GetStressWeekly(ctx, args.Date("end"), args.IntOrDefault("weeks", 4))
		},
	},
	{
		Name:       "GetHydrationStats",
		Service:    "UserSummary",
		Cassette:   "usersummary",
		Path:       "/usersummary-service/stats/hydration/daily/{start}/{end}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "range", Type: endpoint.ParamTypeDateRange, Required: false, Description: "Date range for hydration stats (defaults to last 7 days)"},
		},
		CLICommand:    "summary",
		CLISubcommand: "hydration-stats",
		MCPTool:       "get_hydration_stats",
		Short:         "Get hydration stats",
		Long:          "Get daily hydration intake stats for a date range",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			end := time.Now()
			start := end.AddDate(0, 0, -7)
			if args.HasParam("start") {
				start = args.Date("start")
			}
			if args.HasParam("end") {
				end = args.Date("end")
			}
			return client.UserSummary.GetHydrationStats(ctx, start, end)
		},
	},
	{
		Name:       "GetIntensityMinutesDailyStats",
		Service:    "UserSummary",
		Cassette:   "usersummary",
		Path:       "/usersummary-service/stats/im/daily/{start}/{end}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "range", Type: endpoint.ParamTypeDateRange, Required: false, Description: "Date range for daily intensity minutes (defaults to last 7 days)"},
		},
		CLICommand:    "summary",
		CLISubcommand: "im-daily",
		MCPTool:       "get_intensity_minutes_daily_stats",
		Short:         "Get daily intensity minutes stats",
		Long:          "Get daily moderate and vigorous intensity minutes for a date range",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			end := time.Now()
			start := end.AddDate(0, 0, -7)
			if args.HasParam("start") {
				start = args.Date("start")
			}
			if args.HasParam("end") {
				end = args.Date("end")
			}
			return client.UserSummary.GetIntensityMinutesDaily(ctx, start, end)
		},
	},
	{
		Name:       "GetIntensityMinutesWeeklyStats",
		Service:    "UserSummary",
		Cassette:   "usersummary",
		Path:       "/usersummary-service/stats/im/weekly/{start}/{end}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "range", Type: endpoint.ParamTypeDateRange, Required: false, Description: "Date range for weekly intensity minutes (defaults to last 28 days)"},
		},
		CLICommand:    "summary",
		CLISubcommand: "im-weekly",
		MCPTool:       "get_intensity_minutes_weekly_stats",
		Short:         "Get weekly intensity minutes stats",
		Long:          "Get weekly moderate and vigorous intensity minutes aggregates for a date range",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			end := time.Now()
			start := end.AddDate(0, 0, -28)
			if args.HasParam("start") {
				start = args.Date("start")
			}
			if args.HasParam("end") {
				end = args.Date("end")
			}
			return client.UserSummary.GetIntensityMinutesWeekly(ctx, start, end)
		},
	},
}
