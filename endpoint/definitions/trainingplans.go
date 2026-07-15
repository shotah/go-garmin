package definitions

import (
	"context"
	"fmt"

	"github.com/shotah/go-garmin/endpoint"
	"github.com/shotah/go-garmin/garmin"
)

// TrainingPlanEndpoints defines training plan / Garmin Coach endpoints.
var TrainingPlanEndpoints = []endpoint.Endpoint{
	{
		Name:          "ListTrainingPlans",
		Service:       "TrainingPlans",
		Cassette:      "trainingplans",
		Path:          "/trainingplan-service/trainingplan/plans",
		HTTPMethod:    "GET",
		CLICommand:    "plans",
		CLISubcommand: "list",
		MCPTool:       "list_training_plans",
		Short:         "List training plans",
		Long:          "List training plans for the current user",
		Handler: func(ctx context.Context, c any, _ *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.TrainingPlans.List(ctx)
		},
	},
	{
		Name:       "GetTrainingPlanPhased",
		Service:    "TrainingPlans",
		Cassette:   "trainingplans",
		Path:       "/trainingplan-service/trainingplan/phased/{planId}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "plan_id", Type: endpoint.ParamTypeInt, Required: true, Description: "Training plan ID"},
		},
		CLICommand:    "plans",
		CLISubcommand: "phased",
		MCPTool:       "get_training_plan_phased",
		Short:         "Get phased training plan",
		Long:          "Get a phased training plan by ID",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.TrainingPlans.GetPhased(ctx, int64(args.Int("plan_id")))
		},
	},
	{
		Name:       "GetTrainingPlanAdaptive",
		Service:    "TrainingPlans",
		Cassette:   "trainingplans",
		Path:       "/trainingplan-service/trainingplan/fbt-adaptive/{planId}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "plan_id", Type: endpoint.ParamTypeInt, Required: true, Description: "Training plan ID"},
		},
		CLICommand:    "plans",
		CLISubcommand: "adaptive",
		MCPTool:       "get_training_plan_adaptive",
		Short:         "Get adaptive training plan",
		Long:          "Get an FBT adaptive (Garmin Coach) training plan by ID",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.TrainingPlans.GetAdaptive(ctx, int64(args.Int("plan_id")))
		},
	},
}
