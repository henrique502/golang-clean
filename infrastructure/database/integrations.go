package database

import (
	"context"

	"github.com/henrique502/golang-clean/domain/integration"
	"github.com/pkg/errors"
)

func (c Client) IntegrationUpSert(data integration.Integration) error {
	sql := `
    INSERT INTO ` + integration.TableName + `
      (id, name, type, enabled, created_at, updated_at)
    VALUES
      ($1, $2, $3, $4, NOW(), NOW())
    ON CONFLICT (id) DO UPDATE SET
      name = excluded.name,
      type = excluded.type,
      enabled = excluded.enabled,
      updated_at = NOW();
  `

	_, err := c.conn.Exec(context.Background(), sql,
		data.ID, data.Name, data.Type, data.Enabled)

	return errors.Wrap(err, "IntegrationUpSert sql error")
}
