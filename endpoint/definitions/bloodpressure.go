package definitions

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/shotah/go-garmin/endpoint"
	"github.com/shotah/go-garmin/garmin"
)

var bloodPressureLogBodyConfig = &endpoint.BodyConfig{
	Type: reflect.TypeFor[garmin.BloodPressureLogRequest](),
	Description: `JSON object to log a blood pressure reading. Required fields: measurementTimestampLocal, measurementTimestampGMT, systolic, diastolic, pulse.

- measurementTimestampLocal (string): Local timestamp, e.g. "2026-07-15T12:00:00.000"
- measurementTimestampGMT (string): GMT timestamp, e.g. "2026-07-15T19:00:00.000"
- systolic (number): Systolic pressure
- diastolic (number): Diastolic pressure
- pulse (number): Pulse rate
- sourceType (string, optional): Defaults to "MANUAL"
- notes (string, optional): Free-form notes

The live Connect API expects POST.`,
	Example: `{
  "measurementTimestampLocal": "2026-07-15T12:00:00.000",
  "measurementTimestampGMT": "2026-07-15T19:00:00.000",
  "systolic": 120,
  "diastolic": 80,
  "pulse": 60,
  "sourceType": "MANUAL"
}`,
}

// BloodPressureEndpoints defines blood-pressure endpoints.
var BloodPressureEndpoints = []endpoint.Endpoint{
	{
		Name:       "GetBloodPressureRange",
		Service:    "BloodPressure",
		Cassette:   "bloodpressure",
		Path:       "/bloodpressure-service/bloodpressure/range/{start}/{end}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "range", Type: endpoint.ParamTypeDateRange, Required: false, Description: "Date range for blood pressure data (defaults to last 7 days)"},
		},
		CLICommand:    "bp",
		CLISubcommand: "range",
		MCPTool:       "get_blood_pressure_range",
		Short:         "Get blood pressure range",
		Long:          "Get blood pressure measurements and category stats for a date range",
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
			return client.BloodPressure.GetRange(ctx, start, end)
		},
	},
	{
		Name:          "LogBloodPressure",
		Service:       "BloodPressure",
		Cassette:      "none",
		Path:          "/bloodpressure-service/bloodpressure",
		HTTPMethod:    "POST",
		Body:          bloodPressureLogBodyConfig,
		CLICommand:    "bp",
		CLISubcommand: "log",
		MCPTool:       "log_blood_pressure",
		Short:         "Log blood pressure",
		Long:          "Log a manual blood pressure measurement. Use --file, --json, or stdin. Live API uses POST.",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			req, ok := args.Body.(*garmin.BloodPressureLogRequest)
			if !ok {
				return nil, fmt.Errorf("invalid blood pressure log body type: %T", args.Body)
			}
			return client.BloodPressure.Log(ctx, req)
		},
	},
	{
		Name:       "DeleteBloodPressure",
		Service:    "BloodPressure",
		Cassette:   "none",
		Path:       "/bloodpressure-service/bloodpressure/{date}/{version}",
		HTTPMethod: "DELETE",
		Params: []endpoint.Param{
			{Name: "date", Type: endpoint.ParamTypeDate, Required: true, Description: "Calendar date of the measurement (YYYY-MM-DD)"},
			{Name: "version", Type: endpoint.ParamTypeInt, Required: true, Description: "Measurement version ID to delete"},
		},
		CLICommand:    "bp",
		CLISubcommand: "delete",
		MCPTool:       "delete_blood_pressure",
		Short:         "Delete blood pressure",
		Long:          "Delete a blood pressure measurement by calendar date and version",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			if err := client.BloodPressure.Delete(ctx, args.Date("date"), args.Int("version")); err != nil {
				return nil, err
			}
			return map[string]string{"status": "success"}, nil
		},
	},
}
