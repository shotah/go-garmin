package definitions

import (
	"context"
	"fmt"
	"time"

	"github.com/shotah/go-garmin/endpoint"
	"github.com/shotah/go-garmin/garmin"
)

// WellnessEndpoints defines all wellness-related endpoints.
var WellnessEndpoints = []endpoint.Endpoint{
	{
		Name:       "GetDailyStress",
		Service:    "Wellness",
		Cassette:   "wellness_stress",
		Path:       "/wellness-service/wellness/dailyStress",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "date", Type: endpoint.ParamTypeDate, Required: false, Description: "Date to get stress data for (YYYY-MM-DD, defaults to today)"},
		},
		CLICommand:    "wellness",
		CLISubcommand: "stress",
		MCPTool:       "get_stress",
		Short:         "Get stress levels for a date",
		Long:          "Get stress levels throughout the day including max, average, and stress chart values",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Wellness.GetDailyStress(ctx, args.Date("date"))
		},
	},
	{
		Name:       "GetBodyBatteryEvents",
		Service:    "Wellness",
		Cassette:   "wellness_body_battery",
		Path:       "/wellness-service/wellness/bodyBattery/events",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "date", Type: endpoint.ParamTypeDate, Required: false, Description: "Date to get body battery events for (YYYY-MM-DD, defaults to today)"},
		},
		CLICommand:    "wellness",
		CLISubcommand: "body-battery",
		MCPTool:       "get_body_battery",
		Short:         "Get body battery events for a date",
		Long:          "Get body battery drain and charge events throughout the day including sleep, activity, and stress impacts",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Wellness.GetBodyBatteryEvents(ctx, args.Date("date"))
		},
	},
	{
		Name:       "GetDailyHeartRate",
		Service:    "Wellness",
		Cassette:   "wellness_heart_rate",
		Path:       "/wellness-service/wellness/dailyHeartRate",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "date", Type: endpoint.ParamTypeDate, Required: false, Description: "Date to get heart rate data for (YYYY-MM-DD, defaults to today)"},
		},
		CLICommand:    "wellness",
		CLISubcommand: "heart-rate",
		MCPTool:       "get_heart_rate",
		Short:         "Get heart rate data for a date",
		Long:          "Get heart rate data for a day including resting HR, max HR, and time in zones",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Wellness.GetDailyHeartRate(ctx, args.Date("date"))
		},
	},
	{
		Name:       "GetDailySpO2",
		Service:    "Wellness",
		Cassette:   "wellness_extended",
		Path:       "/wellness-service/wellness/daily/spo2",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "date", Type: endpoint.ParamTypeDate, Required: false, Description: "Date to get SpO2 data for (YYYY-MM-DD, defaults to today)"},
		},
		CLICommand:    "wellness",
		CLISubcommand: "spo2",
		MCPTool:       "get_spo2",
		Short:         "Get blood oxygen (SpO2) for a date",
		Long:          "Get blood oxygen (SpO2) readings for a day including average, lowest, and sleep SpO2",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Wellness.GetDailySpO2(ctx, args.Date("date"))
		},
	},
	{
		Name:       "GetDailyRespiration",
		Service:    "Wellness",
		Cassette:   "wellness_extended",
		Path:       "/wellness-service/wellness/daily/respiration",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "date", Type: endpoint.ParamTypeDate, Required: false, Description: "Date to get respiration data for (YYYY-MM-DD, defaults to today)"},
		},
		CLICommand:    "wellness",
		CLISubcommand: "respiration",
		MCPTool:       "get_respiration",
		Short:         "Get respiration data for a date",
		Long:          "Get respiration rate data for a day including waking and sleep respiration averages",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Wellness.GetDailyRespiration(ctx, args.Date("date"))
		},
	},
	{
		Name:       "GetDailyIntensityMinutes",
		Service:    "Wellness",
		Cassette:   "wellness_extended",
		Path:       "/wellness-service/wellness/daily/im",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "date", Type: endpoint.ParamTypeDate, Required: false, Description: "Date to get intensity minutes for (YYYY-MM-DD, defaults to today)"},
		},
		CLICommand:    "wellness",
		CLISubcommand: "intensity-minutes",
		MCPTool:       "get_intensity_minutes",
		Short:         "Get intensity minutes for a date",
		Long:          "Get weekly intensity minutes (moderate and vigorous activity) and progress toward weekly goal",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Wellness.GetDailyIntensityMinutes(ctx, args.Date("date"))
		},
	},
	{
		Name:       "GetDailyEvents",
		Service:    "Wellness",
		Cassette:   "wellness_daily_extra",
		Path:       "/wellness-service/wellness/dailyEvents",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "date", Type: endpoint.ParamTypeDate, Required: false, Description: "Date to get daily events for (YYYY-MM-DD, defaults to today)"},
		},
		CLICommand:    "wellness",
		CLISubcommand: "events",
		MCPTool:       "get_daily_events",
		Short:         "Get daily wellness events",
		Long:          "Get daily wellness events including auto-detected activities for a date",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Wellness.GetDailyEvents(ctx, args.Date("date"))
		},
	},
	{
		Name:       "GetWellnessDailySleep",
		Service:    "Wellness",
		Cassette:   "wellness_daily_extra",
		Path:       "/wellness-service/wellness/dailySleepData/{displayName}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "date", Type: endpoint.ParamTypeDate, Required: false, Description: "Date to get sleep data for (YYYY-MM-DD, defaults to today)"},
			{Name: "display_name", Type: endpoint.ParamTypeString, Required: false, Description: "User display name (defaults to current user)"},
		},
		CLICommand:    "wellness",
		CLISubcommand: "sleep",
		MCPTool:       "get_wellness_sleep",
		Short:         "Get sleep via wellness-service",
		Long:          "Get sleep data using the wellness-service dailySleepData path (alternative to `garmin sleep`)",
		DependsOn:     "GetSocialProfile",
		ArgProvider: func(result any) map[string]any {
			profile, ok := result.(*garmin.SocialProfile)
			if !ok || profile == nil {
				return nil
			}
			return map[string]any{"display_name": profile.DisplayName}
		},
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			displayName, err := client.ResolveDisplayName(ctx, args.String("display_name"))
			if err != nil {
				return nil, err
			}
			return client.Wellness.GetDailySleep(ctx, displayName, args.Date("date"))
		},
	},
	{
		Name:       "GetDailySummaryChart",
		Service:    "Wellness",
		Cassette:   "wellness_daily_extra",
		Path:       "/wellness-service/wellness/dailySummaryChart/{displayName}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "date", Type: endpoint.ParamTypeDate, Required: false, Description: "Date to get steps chart for (YYYY-MM-DD, defaults to today)"},
			{Name: "display_name", Type: endpoint.ParamTypeString, Required: false, Description: "User display name (defaults to current user)"},
		},
		CLICommand:    "wellness",
		CLISubcommand: "steps",
		MCPTool:       "get_steps_chart",
		Short:         "Get daily steps chart",
		Long:          "Get intraday steps/activity chart data (typically 15-minute intervals) for a date",
		DependsOn:     "GetSocialProfile",
		ArgProvider: func(result any) map[string]any {
			profile, ok := result.(*garmin.SocialProfile)
			if !ok || profile == nil {
				return nil
			}
			return map[string]any{"display_name": profile.DisplayName}
		},
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			displayName, err := client.ResolveDisplayName(ctx, args.String("display_name"))
			if err != nil {
				return nil, err
			}
			return client.Wellness.GetDailySummaryChart(ctx, displayName, args.Date("date"))
		},
	},
	{
		Name:       "GetDailyFloors",
		Service:    "Wellness",
		Cassette:   "wellness_daily_extra",
		Path:       "/wellness-service/wellness/floorsChartData/daily",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "date", Type: endpoint.ParamTypeDate, Required: false, Description: "Date to get floors data for (YYYY-MM-DD, defaults to today)"},
		},
		CLICommand:    "wellness",
		CLISubcommand: "floors",
		MCPTool:       "get_floors",
		Short:         "Get floors climbed chart",
		Long:          "Get floors ascended and descended chart data for a date",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Wellness.GetDailyFloors(ctx, args.Date("date"))
		},
	},
	{
		Name:       "GetBodyBatteryReports",
		Service:    "Wellness",
		Cassette:   "wellness_daily_extra",
		Path:       "/wellness-service/wellness/bodyBattery/reports/daily",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "range", Type: endpoint.ParamTypeDateRange, Required: false, Description: "Date range for body battery reports (defaults to last 7 days)"},
		},
		CLICommand:    "wellness",
		CLISubcommand: "body-battery-reports",
		MCPTool:       "get_body_battery_reports",
		Short:         "Get body battery reports",
		Long:          "Get daily body battery charged/drained reports over a date range",
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
			return client.Wellness.GetBodyBatteryReports(ctx, start, end)
		},
	},
	{
		Name:       "GetSleepScoreStats",
		Service:    "Wellness",
		Cassette:   "wellness_daily_extra",
		Path:       "/wellness-service/stats/daily/sleep/score",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "range", Type: endpoint.ParamTypeDateRange, Required: false, Description: "Date range for sleep score stats (defaults to last 7 days)"},
		},
		CLICommand:    "wellness",
		CLISubcommand: "sleep-score",
		MCPTool:       "get_sleep_score_stats",
		Short:         "Get sleep score stats",
		Long:          "Get daily sleep scores over a date range",
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
			return client.Wellness.GetSleepScoreStats(ctx, start, end)
		},
	},
}
