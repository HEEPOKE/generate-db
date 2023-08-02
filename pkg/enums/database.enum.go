package enums

type Database string

const (
	MYSQL      Database = "mysql"
	POSTGRESQL Database = "postgresql"
	SQLSERVER  Database = "sqlserver"
	MONGO      Database = "mongodb"
)
