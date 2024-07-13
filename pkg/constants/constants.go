package constants

import "time"

const (
	MYSQL_ROOT_PASSWORD = "GraphQL-demo"
	MYSQL_DATABASE      = "GraphQL-demo"
	MYSQL_USER          = "GraphQL-demo"
	MYSQL_PASSWORD      = "GraphQL-demo"

	DbAddr = "127.0.0.1"

	MaxIdleConns    = 10
	MaxConnections  = 100
	ConnMaxLifetime = time.Hour

	JWTValue = "GraphQL-demo"
)
