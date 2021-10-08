package opsgenie

import (
	"github.com/henrique502/golang-clean/domain/team"
)

func (c *Client) GetTeamList(nextURL *string, query map[string]string) ([]team.Team, *string, error) {
	type Response struct {
		PaginateResponse
		Data []team.Team `json:"data"`
	}

	r := &Response{}

	err := c.GetList("/v2/teams", nextURL, query, r)
	if err != nil {
		return nil, nil, err
	}

	return r.Data, &r.Paging.Next, nil
}
