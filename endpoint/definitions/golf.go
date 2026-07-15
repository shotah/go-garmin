package definitions

import (
	"context"
	"fmt"

	"github.com/shotah/go-garmin/endpoint"
	"github.com/shotah/go-garmin/garmin"
)

// GolfEndpoints defines Garmin Golf scorecard endpoints.
var GolfEndpoints = []endpoint.Endpoint{
	{
		Name:       "ListGolfScorecards",
		Service:    "Golf",
		Cassette:   "golf",
		Path:       "/gcs-golfcommunity/api/v2/scorecard/summary",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "start", Type: endpoint.ParamTypeInt, Required: false, Description: "Starting index (0-based, defaults to 0)"},
			{Name: "limit", Type: endpoint.ParamTypeInt, Required: false, Description: "Maximum results (defaults to 20)"},
		},
		CLICommand:    "golf",
		CLISubcommand: "list",
		MCPTool:       "list_golf_scorecards",
		Short:         "List golf scorecards",
		Long:          "List recent golf scorecard summaries (rounds) for the authenticated user",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Golf.ListScorecards(ctx, args.IntOrDefault("start", 0), args.IntOrDefault("limit", 20))
		},
	},
	{
		Name:       "GetGolfScorecard",
		Service:    "Golf",
		Cassette:   "golf",
		Path:       "/gcs-golfcommunity/api/v2/scorecard/detail",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "scorecard_id", Type: endpoint.ParamTypeInt, Required: true, Description: "Golf scorecard ID"},
		},
		CLICommand:    "golf",
		CLISubcommand: "get",
		MCPTool:       "get_golf_scorecard",
		Short:         "Get golf scorecard detail",
		Long:          "Get detailed golf scorecard data for a scorecard ID (includes longest-shot distance when available)",
		DependsOn:     "ListGolfScorecards",
		ArgProvider: func(result any) map[string]any {
			summaries, ok := result.(*garmin.GolfScorecardSummaries)
			if !ok || summaries.FirstID() == 0 {
				return nil
			}
			return map[string]any{"scorecard_id": summaries.FirstID()}
		},
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Golf.GetScorecard(ctx, int64(args.Int("scorecard_id")))
		},
	},
	{
		Name:       "GetGolfShotData",
		Service:    "Golf",
		Cassette:   "golf",
		Path:       "/gcs-golfcommunity/api/v2/shot/scorecard/{scorecard_id}/hole",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "scorecard_id", Type: endpoint.ParamTypeInt, Required: true, Description: "Golf scorecard ID"},
			{Name: "holes", Type: endpoint.ParamTypeString, Required: false, Description: "Comma-separated hole numbers (defaults to 1-18)"},
		},
		CLICommand:    "golf",
		CLISubcommand: "shots",
		MCPTool:       "get_golf_shot_data",
		Short:         "Get golf shot data",
		Long:          "Get shot-by-shot golf data for holes on a scorecard (useful with Approach/CT10 tracking)",
		DependsOn:     "ListGolfScorecards",
		ArgProvider: func(result any) map[string]any {
			summaries, ok := result.(*garmin.GolfScorecardSummaries)
			if !ok || summaries.FirstID() == 0 {
				return nil
			}
			return map[string]any{"scorecard_id": summaries.FirstID()}
		},
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Golf.GetShotData(ctx, int64(args.Int("scorecard_id")), args.String("holes"))
		},
	},
}
