package definitions

import (
	"context"
	"fmt"
	"reflect"

	"github.com/shotah/go-garmin/endpoint"
	"github.com/shotah/go-garmin/garmin"
)

var lifestyleBehaviourBodyConfig = &endpoint.BodyConfig{
	Type: reflect.TypeFor[garmin.LifestyleBehaviourRequest](),
	Description: `JSON object to create a custom lifestyle behaviour. Required fields: name.

- name (string): Behaviour name
- category (string, optional): Defaults to "CUSTOM" (also CUSTOM_SLEEP_RELATED, …)
- sleepRelated (boolean, optional): Whether the behaviour is sleep-related
- tracked (boolean, optional): Whether the behaviour is tracked

The live Connect API expects POST.`,
	Example: `{
  "name": "Meditation",
  "category": "CUSTOM",
  "sleepRelated": false,
  "tracked": true
}`,
}

// LifestyleEndpoints defines lifestyle logging endpoints.
var LifestyleEndpoints = []endpoint.Endpoint{
	{
		Name:       "GetDailyLifestyleLog",
		Service:    "Lifestyle",
		Cassette:   "lifestyle",
		Path:       "/lifestylelogging-service/dailyLog/{date}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "date", Type: endpoint.ParamTypeDate, Required: false, Description: "Date for lifestyle log (YYYY-MM-DD, defaults to today)"},
		},
		CLICommand:    "lifestyle",
		CLISubcommand: "daily",
		MCPTool:       "get_daily_lifestyle_log",
		Short:         "Get daily lifestyle log",
		Long:          "Get lifestyle logging entries and completion stats for a date",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Lifestyle.GetDaily(ctx, args.Date("date"))
		},
	},
	{
		Name:          "CreateLifestyleBehaviour",
		Service:       "Lifestyle",
		Cassette:      "none",
		Path:          "/lifestylelogging-service/behaviours",
		HTTPMethod:    "POST",
		Body:          lifestyleBehaviourBodyConfig,
		CLICommand:    "lifestyle",
		CLISubcommand: "create-behaviour",
		MCPTool:       "create_lifestyle_behaviour",
		Short:         "Create lifestyle behaviour",
		Long:          "Create a custom lifestyle behaviour tag. Use --file, --json, or stdin. Live API uses POST.",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			req, ok := args.Body.(*garmin.LifestyleBehaviourRequest)
			if !ok {
				return nil, fmt.Errorf("invalid lifestyle behaviour body type: %T", args.Body)
			}
			return client.Lifestyle.CreateBehaviour(ctx, req)
		},
	},
}
