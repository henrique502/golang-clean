package opsgenie

import (
	"github.com/henrique502/golang-clean/domain/integration"
)

func (c *Client) GetIntegrationList(nextURL *string, query map[string]string) ([]integration.Integration, *string, error) {
	type Response struct {
		PaginateResponse
		Data []integration.Integration `json:"data"`
	}

	r := &Response{}

	err := c.GetList("/v2/integrations", nextURL, query, r)
	if err != nil {
		return nil, nil, err
	}

	return r.Data, &r.Paging.Next, nil
}
