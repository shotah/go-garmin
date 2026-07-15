package definitions

import (
	"context"
	"fmt"

	"github.com/shotah/go-garmin/endpoint"
	"github.com/shotah/go-garmin/garmin"
)

// PersonalRecordsEndpoints defines personal-record endpoints.
var PersonalRecordsEndpoints = []endpoint.Endpoint{
	{
		Name:       "GetPersonalRecords",
		Service:    "PersonalRecords",
		Cassette:   "personalrecords",
		Path:       "/personalrecord-service/personalrecord/prs/{displayName}",
		HTTPMethod: "GET",
		Params: []endpoint.Param{
			{Name: "display_name", Type: endpoint.ParamTypeString, Required: false, Description: "User display name (defaults to current user)"},
		},
		CLICommand:    "records",
		CLISubcommand: "list",
		MCPTool:       "get_personal_records",
		Short:         "Get personal records",
		Long:          "Get all-time personal records for the user (or another display name)",
		Handler: func(ctx context.Context, c any, args *endpoint.HandlerArgs) (any, error) {
			client, ok := c.(*garmin.Client)
			if !ok {
				return nil, fmt.Errorf("handler received invalid client type: %T, expected *garmin.Client", c)
			}
			displayName, err := client.ResolveDisplayName(ctx, args.String("display_name"))
			if err != nil {
				return nil, err
			}
			return client.PersonalRecords.List(ctx, displayName)
		},
	},
}
