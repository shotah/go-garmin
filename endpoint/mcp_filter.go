package endpoint

import (
	"fmt"
	"sort"
	"strings"
)

// ToolTier controls how many MCP tools are published within selected services.
type ToolTier string

const (
	TierCore     ToolTier = "core"
	TierExtended ToolTier = "extended"
	TierComplete ToolTier = "complete"
)

// ToolFilter narrows which endpoints RegisterTools publishes.
// Empty Services means all services. Tier defaults to complete when empty.
type ToolFilter struct {
	Services map[string]struct{} // lowercased Endpoint.Service names
	Tier     ToolTier
}

// ParseToolFilter builds a filter from CLI flags.
// toolsFlag is space- and/or comma-separated service names (case-insensitive).
// tierFlag is core|extended|complete (default complete).
func ParseToolFilter(toolsFlag, tierFlag string) (ToolFilter, error) {
	tier := ToolTier(strings.ToLower(strings.TrimSpace(tierFlag)))
	if tier == "" {
		tier = TierComplete
	}
	switch tier {
	case TierCore, TierExtended, TierComplete:
	default:
		return ToolFilter{}, fmt.Errorf("invalid --tool-tier %q (want core|extended|complete)", tierFlag)
	}

	f := ToolFilter{Tier: tier, Services: make(map[string]struct{})}
	for _, tok := range splitToolsFlag(toolsFlag) {
		name := strings.ToLower(tok)
		if name == "" {
			continue
		}
		f.Services[name] = struct{}{}
	}
	return f, nil
}

func splitToolsFlag(s string) []string {
	s = strings.ReplaceAll(s, ",", " ")
	return strings.Fields(s)
}

// Allows reports whether ep should be registered as an MCP tool under this filter.
func (f ToolFilter) Allows(ep *Endpoint) bool {
	if ep == nil || ep.MCPTool == "" || ep.RawOutput {
		return false
	}
	if len(f.Services) > 0 {
		svc := strings.ToLower(strings.TrimSpace(ep.Service))
		if _, ok := f.Services[svc]; !ok {
			return false
		}
	}
	switch f.Tier {
	case TierCore:
		_, ok := mcpTierCore[ep.MCPTool]
		return ok
	case TierExtended:
		_, ok := mcpTierExtended[ep.MCPTool]
		return ok
	case TierComplete, "":
		return true
	default:
		return false
	}
}

// ValidateServices checks --tools names against the registry (fail-fast).
func (f ToolFilter) ValidateServices(r *Registry) error {
	if len(f.Services) == 0 || r == nil {
		return nil
	}
	known := make(map[string]struct{})
	for _, ep := range r.endpoints {
		if ep.Service == "" {
			continue
		}
		known[strings.ToLower(ep.Service)] = struct{}{}
	}
	var unknown []string
	for name := range f.Services {
		if _, ok := known[name]; !ok {
			unknown = append(unknown, name)
		}
	}
	if len(unknown) == 0 {
		return nil
	}
	return fmt.Errorf("unknown --tools service(s): %s (known: %s)",
		strings.Join(sortedKeys(unknown), ", "),
		strings.Join(sortedKeys(mapKeys(known)), ", "))
}

// mcpTierCore is the small recovery / coaching surface (~10 tools).
var mcpTierCore = map[string]struct{}{
	"get_current_date":             {},
	"get_sleep":                    {},
	"get_weight":                   {},
	"get_body_battery":             {},
	"get_hrv":                      {},
	"get_training_readiness":       {},
	"list_activities":              {},
	"get_activity":                 {},
	"get_activity_typed_splits":    {},
	"get_activity_split_summaries": {},
}

// mcpTierExtended is core plus common coaching extras.
var mcpTierExtended = func() map[string]struct{} {
	m := make(map[string]struct{}, len(mcpTierCore)+16)
	for k := range mcpTierCore {
		m[k] = struct{}{}
	}
	for _, k := range []string{
		"get_stress",
		"get_heart_rate",
		"get_body_battery_reports",
		"get_sleep_score_stats",
		"get_intensity_minutes",
		"get_training_status",
		"get_vo2max",
		"get_activity_details",
		"get_daily_user_summary",
	} {
		m[k] = struct{}{}
	}
	return m
}()

func mapKeys(m map[string]struct{}) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

func sortedKeys(keys []string) []string {
	sort.Strings(keys)
	return keys
}
