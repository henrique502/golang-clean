package database

import (
	"context"

	"github.com/henrique502/golang-clean/domain/team"
	"github.com/pkg/errors"
)

func (c Client) TeamUpSert(data team.Team) error {
	sql := `
    INSERT INTO ` + team.TableName + `
      (id, name, description, created_at, updated_at)
    VALUES
      ($1, $2, $3, NOW(), NOW())
    ON CONFLICT (id) DO UPDATE SET
      name = excluded.name,
      description = excluded.description,
      updated_at = NOW();
  `

	_, err := c.conn.Exec(context.Background(), sql,
		data.ID, data.Name, data.Description)

	return errors.Wrap(err, "TeamUpSert sql error")
}
