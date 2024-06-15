package postgresclient

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// Config represents the configuration for connecting to the Postgres database.
type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

// Client represents the Postgres client.
type Client struct {
	Conn *pgx.Conn
}

// Create a Postgres client
func NewClient(config Config) (*Client, error) {
	connConfig, err := pgx.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.Username, config.Password, config.Host, config.Port, config.Database))
	if err != nil {
		return nil, err
	}

	conn, err := pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		return nil, err
	}

	return &Client{
		Conn: conn,
	}, nil
}

// Close closes the connection to the Postgres database.
func (c *Client) Close() error {
	return c.Conn.Close(context.Background())
}
