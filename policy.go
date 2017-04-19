package chef

type PolicyService struct {
	client *Client
}

type PolicyListResult struct {
	Uri       string            `json:"uri"`
	Revisions map[string]string `json:"revisions"`
}

type CookbookLock struct {
	Version                 string            `json:"version"`
	Identifier              string            `json:"identifier"`
	DottedDecimalIdentifier string            `json:"dotted_decimal_identifier"`
	CacheKey                string            `json:"cache_key"`
	Origin                  string            `json:"origin"`
	SourceOptions           map[string]string `json:"source_options"`
}

type Policy struct {
	RevisionId           string                  `json:"revision_id"`
	Name                 string                  `json:"name"`
	RunList              []string                `json:"run_list"`
	NamedRunLists        map[string][]string     `json:"named_run_lists,omitempty"`
	IncludedPolicyLocks  map[string]CookbookLock `json:"included_policy_locks,omitempty"`
	DefaultAttributes    interface{}             `json:"default_attributes,omitempty"`
	OverrideAttributes   interface{}             `json:"override_attributes,omitempty"`
	CookbookLocks        map[string]CookbookLock `json:"cookbook_locks,omitempty"`
	SolutionDependencies map[string]string       `json:"solution_dependencies"`
}

// List gets a list of all policies inside of the Chef Server.
// It returns a PolicyListResult object which contains all revisions
// available for the given policy.
func (p *PolicyService) List() (policies *PolicyListResult, err error) {
	err = p.client.magicRequestDecoder("GET", "/policies", nil, &policies)
	return
}
