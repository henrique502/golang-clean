package opsgenie

import (
	"github.com/henrique502/golang-clean/domain/alert"
)

func (c *Client) GetAlertList(nextURL *string, query map[string]string) ([]alert.Alert, *string, error) {
	type Response struct {
		PaginateResponse
		Data []Alert `json:"data"`
	}

	r := &Response{}

	err := c.GetList("/v2/alerts", nextURL, query, r)
	if err != nil {
		return nil, nil, err
	}

	data := []alert.Alert{}
	for _, element := range r.Data {
		a := alert.Alert{
			ID:              element.ID,
			Message:         element.Message,
			Priority:        element.Priority,
			Source:          element.Source,
			ReportAckTime:   element.Report.AckTime,
			ReportCloseTime: element.Report.CloseTime,
			IntegrationID:   element.Integration.ID,
			CreatedAt:       element.CreatedAt,
			UpdatedAt:       element.UpdatedAt,
		}

		data = append(data, a)
	}

	return data, &r.Paging.Next, nil
}
