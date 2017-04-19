package chef

type PolicyGroupService struct {
	client *Client
}

type PolicyGroup struct {
	Name string `json:"name"`
}
