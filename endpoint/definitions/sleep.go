// endpoint/definitions/sleep.go
package definitions

import (
	"context"
	"fmt"

	"github.com/shotah/go-garmin/endpoint"
	garmin "github.com/shotah/go-garmin/garmin"
)

// SleepEndpoints defines all sleep-related API endpoints.
var SleepEndpoints = []endpoint.Endpoint{
	{
		Name:       "GetDailySleep",
		Service:    "Sleep",
		Cassette:   "sleep_daily",
		Path:       "/sleep-service/sleep/dailySleepData",
		HTTPMethod: "GET",

		Params: []endpoint.Param{
			{
				Name:     "date",
				Type:     endpoint.ParamTypeDate,
				Required: false,
				Description: "Wake-up calendar day (YYYY-MM-DD). For 'last night' / this morning's sleep, omit or pass today — " +
					"not yesterday. Defaults to today.",
			},
		},

		CLICommand: "sleep",
		MCPTool:    "get_sleep",
		Short:      "Sleep score + stages for a night",
		Long: "Primary sleep tool: overnight sleep score, total duration, deep/light/REM stages, awake time, " +
			"and Body Battery change. Use for 'how did I sleep', 'sleep score', 'last night', sleep quality. " +
			"Prefer this over get_wellness_sleep. For multi-day score trends only, use get_sleep_score_stats.",

		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Sleep.GetDaily(ctx, args.Date("date"))
		},
	},
}
