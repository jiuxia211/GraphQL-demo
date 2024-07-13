package utils

import (
	"strings"

	"github.com/jiuxia211/GraphQL-demo/pkg/constants"
)

func GetMysqlDsn() string {

	dsn := strings.Join([]string{constants.MYSQL_USER, ":", constants.MYSQL_PASSWORD, "@tcp(", constants.DbAddr, ")/", constants.MYSQL_DATABASE, "?charset=utf8mb4" + "&parseTime=true"}, "")

	return dsn
}
