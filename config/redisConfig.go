package config

// rdbBaseConfig redis基础配置信息
type rdbBaseConfig struct {
	Host        string // 主机
	Port        int    // 端口
	Password    string // 密码
	DB          int    // 数据库名称
	MaxRetries  int    // 最大尝试次数
	DialTimeout int    // 连接超时时间
	PoolSize    int    // 池大小
}

var (
	rdbTest1 = &rdbBaseConfig{
		Host:        "10.10.10.233",
		Port:        6379,
		Password:    "meidai",
		DB:          6,
		MaxRetries:  3,
		DialTimeout: 10,
		PoolSize:    20,
	}

	rdbTest2 = &rdbBaseConfig{
		Host:        "10.10.10.233",
		Port:        6379,
		Password:    "meidai",
		DB:          7,
		MaxRetries:  3,
		DialTimeout: 10,
		PoolSize:    20,
	}
)

// MdbConfigMap 连接信息map
type RdbConfigMap map[string]*rdbBaseConfig

// MdbConfigs 需初始化的多个连接配置
var rdbConfigs = map[string]*rdbBaseConfig{
	//"test": rdbTest1,
	//"test2": rdbTest2,
}

// GetMdbConfigs 获取redis配置信息
func GetRdbConfigs() RdbConfigMap {
	return rdbConfigs
}
