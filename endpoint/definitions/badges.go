package definitions

import (
	"context"
	"fmt"

	"github.com/shotah/go-garmin/endpoint"
	"github.com/shotah/go-garmin/garmin"
)

// BadgeEndpoints defines badge and challenge endpoints.
var BadgeEndpoints = []endpoint.Endpoint{
	{
		Name:          "GetEarnedBadges",
		Service:       "Badges",
		Cassette:      "badges",
		Path:          "/badge-service/badge/earned",
		HTTPMethod:    "GET",
		CLICommand:    "badges",
		CLISubcommand: "earned",
		MCPTool:       "get_earned_badges",
		Short:         "Get earned badges",
		Long:          "Get badges the current user has earned",
		Handler: func(ctx context.Context, c any, _ *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Badges.ListEarned(ctx)
		},
	},
	{
		Name:          "GetAvailableBadges",
		Service:       "Badges",
		Cassette:      "badges",
		Path:          "/badge-service/badge/available",
		HTTPMethod:    "GET",
		CLICommand:    "badges",
		CLISubcommand: "available",
		MCPTool:       "get_available_badges",
		Short:         "Get available badges",
		Long:          "Get badges available to the current user (including exclusive badges)",
		Handler: func(ctx context.Context, c any, _ *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Badges.ListAvailable(ctx)
		},
	},
	{
		Name:       "GetCompletedBadgeChallenges",
		Service:    "Badges",
		Cassette:   "badges",
		Path:       "/badgechallenge-service/badgeChallenge/completed",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "start", Type: endpoint.ParamTypeInt, Required: false, Description: "Starting index (0-based, defaults to 0)"},
			{Name: "limit", Type: endpoint.ParamTypeInt, Required: false, Description: "Maximum number of challenges to return (defaults to 20)"},
		},
		CLICommand:    "badges",
		CLISubcommand: "challenges-completed",
		MCPTool:       "get_completed_badge_challenges",
		Short:         "Get completed badge challenges",
		Long:          "Get completed badge challenges with pagination",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Badges.ListCompletedChallenges(ctx, args.IntOrDefault("start", 0), args.IntOrDefault("limit", 20))
		},
	},
	{
		Name:       "GetAvailableBadgeChallenges",
		Service:    "Badges",
		Cassette:   "badges",
		Path:       "/badgechallenge-service/badgeChallenge/available",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "start", Type: endpoint.ParamTypeInt, Required: false, Description: "Starting index (0-based, defaults to 0)"},
			{Name: "limit", Type: endpoint.ParamTypeInt, Required: false, Description: "Maximum number of challenges to return (defaults to 20)"},
		},
		CLICommand:    "badges",
		CLISubcommand: "challenges-available",
		MCPTool:       "get_available_badge_challenges",
		Short:         "Get available badge challenges",
		Long:          "Get available badge challenges with pagination",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Badges.ListAvailableChallenges(ctx, args.IntOrDefault("start", 0), args.IntOrDefault("limit", 20))
		},
	},
	{
		Name:       "GetNonCompletedBadgeChallenges",
		Service:    "Badges",
		Cassette:   "badges",
		Path:       "/badgechallenge-service/badgeChallenge/non-completed",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "start", Type: endpoint.ParamTypeInt, Required: false, Description: "Starting index (0-based, defaults to 0)"},
			{Name: "limit", Type: endpoint.ParamTypeInt, Required: false, Description: "Maximum number of challenges to return (defaults to 20)"},
		},
		CLICommand:    "badges",
		CLISubcommand: "challenges-open",
		MCPTool:       "get_non_completed_badge_challenges",
		Short:         "Get open badge challenges",
		Long:          "Get non-completed (open) badge challenges with pagination",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Badges.ListNonCompletedChallenges(ctx, args.IntOrDefault("start", 0), args.IntOrDefault("limit", 20))
		},
	},
	{
		Name:       "GetVirtualChallengesInProgress",
		Service:    "Badges",
		Cassette:   "badges",
		Path:       "/badgechallenge-service/virtualChallenge/inProgress",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "start", Type: endpoint.ParamTypeInt, Required: false, Description: "Starting index (0-based, defaults to 0)"},
			{Name: "limit", Type: endpoint.ParamTypeInt, Required: false, Description: "Maximum number of challenges to return (defaults to 20)"},
		},
		CLICommand:    "badges",
		CLISubcommand: "virtual",
		MCPTool:       "get_virtual_challenges_in_progress",
		Short:         "Get virtual challenges in progress",
		Long:          "Get in-progress virtual challenges with pagination",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Badges.ListVirtualChallengesInProgress(ctx, args.IntOrDefault("start", 0), args.IntOrDefault("limit", 20))
		},
	},
	{
		Name:       "GetAdHocHistoricalChallenges",
		Service:    "Badges",
		Cassette:   "badges",
		Path:       "/adhocchallenge-service/adHocChallenge/historical",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "start", Type: endpoint.ParamTypeInt, Required: false, Description: "Starting index (0-based, defaults to 0)"},
			{Name: "limit", Type: endpoint.ParamTypeInt, Required: false, Description: "Maximum number of challenges to return (defaults to 20)"},
		},
		CLICommand:    "badges",
		CLISubcommand: "adhoc",
		MCPTool:       "get_adhoc_historical_challenges",
		Short:         "Get historical ad-hoc challenges",
		Long:          "Get historical ad-hoc challenges with pagination",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			return client.Badges.ListAdHocHistorical(ctx, args.IntOrDefault("start", 0), args.IntOrDefault("limit", 20))
		},
	},
}
