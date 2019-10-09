package database

var (
	// GlobalDBSql 全局mysql连接池
	GlobalDBSql DBSql
	// GlobalMDB 全局mongodb连接池
	GlobalMDB MDB
)

func init() {
	// 初始化mysql连接池
	GlobalDBSql = GetSqlDBMap()
	// 初始化mongodb连接池
	GlobalMDB = GetMDBMap()
}
