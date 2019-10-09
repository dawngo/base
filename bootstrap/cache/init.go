package cache

// GlobalRDB 全局redis连接池
var GlobalRDB RDB

func init() {
	// 初始化全局redis连接池
	GlobalRDB = GetRDBMap()
}
