package database

import (
	"context"
	"os"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	conn Connection
}

type Connection interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Close(ctx context.Context) error
}

func New() Client {
	endpoint := os.Getenv("POSTGRES_URL")
	connection, err := pgx.Connect(context.Background(), endpoint)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = connection.Ping(context.Background())
	if err != nil {
		log.Fatalln("ping error:", err)
	}

	return Client{
		conn: connection,
	}
}

func (c *Client) Close() error {
	return c.conn.Close(context.Background())
}
