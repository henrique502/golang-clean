package database

import (
	"context"

	"github.com/henrique502/golang-clean/domain/alert"
	"github.com/pkg/errors"
)

func (c Client) AlertUpSert(data alert.Alert) error {
	sql := `
    INSERT INTO ` + alert.TableName + `
      (id, priority, source, message, report_ack_time, report_close_time, integration_id, created_at, updated_at)
    VALUES
      ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    ON CONFLICT (id) DO UPDATE SET
      priority = excluded.priority,
      source = excluded.source,
      message = excluded.message,
      report_close_time = excluded.report_close_time,
      report_ack_time = excluded.report_ack_time,
      integration_id = excluded.integration_id,
      created_at = excluded.created_at,
      updated_at = excluded.updated_at;
  `

	_, err := c.conn.Exec(context.Background(), sql,
		data.ID, data.Priority, data.Source, data.Message, data.ReportAckTime,
		data.ReportCloseTime, data.IntegrationID, data.CreatedAt, data.UpdatedAt)

	return errors.Wrap(err, "AlertUpSert sql error")
}
