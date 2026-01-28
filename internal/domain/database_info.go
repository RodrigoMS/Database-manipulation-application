package domain

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/RodrigoMS/app/cmd/internal/database"
)

type DatabaseInfo struct {
	ServerVersion     string
	MaxConnections    string
	OpenedConnections string
}

func GetDBInfo() (*DatabaseInfo, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var info DatabaseInfo

	db := database.GetDB()

	err := db.QueryRowContext(
		ctx,
		`SELECT
		  split_part(current_setting('server_version'), ' ', 1) AS ServerVersion,
			current_setting('max_connections')AS MaxConnections,
			(SELECT COUNT(*) FROM pg_stat_activity WHERE datname = $1) AS OpenedConnections;`,
		os.Getenv("PGDATABASE"),
	).Scan(&info.ServerVersion, &info.MaxConnections, &info.OpenedConnections)

	if err != nil {
		return nil, fmt.Errorf("failed to get database info: %v", err)
	}

	return &info, nil
}