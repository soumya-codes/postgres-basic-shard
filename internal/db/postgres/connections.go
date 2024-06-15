package postgres

import (
	"fmt"
	"strconv"
	"strings"

	postgresclient "github.com/soumya-codes/postgres-static-shard/internal/client/postgres"
)

var (
	connections = make(map[int]*postgresclient.Client)
)

func init() {
	_, err := initConnections()
	if err != nil {
		fmt.Printf("error initializing connections: %v\n", err)
		panic(err)
	}
}

// InitConnections returns a map that contains shard to connection mapping.
func initConnections() (map[int]*postgresclient.Client, error) {
	// Create and store connections for each shard
	connShard1, err := postgresclient.NewClient(postgresclient.Config{
		Host:     "localhost",
		Port:     5432,
		Username: "test_user",
		Password: "test_password",
		Database: "shard0",
	})
	if err != nil {
		return nil, err
	}

	connections[0] = connShard1

	connShard2, err := postgresclient.NewClient(postgresclient.Config{
		Host:     "localhost",
		Port:     5433,
		Username: "test_user",
		Password: "test_password",
		Database: "shard1",
	})
	if err != nil {
		return nil, err
	}

	connections[1] = connShard2

	return connections, nil
}

func GetShardConnection(machineId string) (*postgresclient.Client, error) {
	parts := strings.Split(machineId, "_")

	if len(parts) < 2 {
		return nil, fmt.Errorf("error parsing machine-id %s", machineId)
	}

	numStr := parts[1]
	machine_num, err := strconv.Atoi(numStr)
	if err != nil {
		return nil, err
	}

	shardKey := machine_num % 2

	return connections[shardKey], nil
}
