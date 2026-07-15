package definitions

import (
	"context"
	"fmt"
	"time"

	"github.com/shotah/go-garmin/endpoint"
	"github.com/shotah/go-garmin/garmin"
)

// PeriodicHealthEndpoints defines menstrual cycle and pregnancy endpoints.
var PeriodicHealthEndpoints = []endpoint.Endpoint{
	{
		Name:       "GetMenstrualDayView",
		Service:    "PeriodicHealth",
		Cassette:   "periodichealth",
		Path:       "/periodichealth-service/menstrualcycle/dayview/{date}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "date", Type: endpoint.ParamTypeDate, Required: false, Description: "Date for day view (YYYY-MM-DD, defaults to today)"},
		},
		CLICommand:    "health",
		CLISubcommand: "day",
		MCPTool:       "get_menstrual_day_view",
		Short:         "Get menstrual day view",
		Long:          "Get menstrual/pregnancy day view including cycle summary and logged symptoms for a date",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.PeriodicHealth.GetMenstrualDayView(ctx, args.Date("date"))
		},
	},
	{
		Name:       "GetMenstrualCalendar",
		Service:    "PeriodicHealth",
		Cassette:   "periodichealth",
		Path:       "/periodichealth-service/menstrualcycle/calendar/{start}/{end}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "range", Type: endpoint.ParamTypeDateRange, Required: false, Description: "Date range for calendar view (defaults to last 7 days)"},
		},
		CLICommand:    "health",
		CLISubcommand: "calendar",
		MCPTool:       "get_menstrual_calendar",
		Short:         "Get menstrual calendar",
		Long:          "Get menstrual calendar data including cycle summaries and logged symptom/ovulation/note days",
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
			return client.PeriodicHealth.GetMenstrualCalendar(ctx, start, end)
		},
	},
	{
		Name:          "GetPregnancySnapshot",
		Service:       "PeriodicHealth",
		Cassette:      "periodichealth",
		Path:          "/periodichealth-service/menstrualcycle/pregnancysnapshot",
		HTTPMethod:    "GET",
		CLICommand:    "health",
		CLISubcommand: "pregnancy",
		MCPTool:       "get_pregnancy_snapshot",
		Short:         "Get pregnancy snapshot",
		Long:          "Get the pregnancy snapshot summary for the current user",
		Handler: func(ctx context.Context, c any, _ *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.PeriodicHealth.GetPregnancySnapshot(ctx)
		},
	},
}
