package config

// mdbBaseConfig mysql基础配置
type mdbBaseConfig struct {
	Host     string // 主机
	Port     int    // 端口
	Username string // 用户名
	Password string // 密码
	Database string // 数据库名称
}

var (
	mdbTest1 = &mdbBaseConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "zxczxc111",
		Database: "test_db",
	}

	mdbMeidai = &mdbBaseConfig{
		Host:     "10.10.10.233",
		Port:     3306,
		Username: "root",
		Password: "meidai",
		Database: "meidai",
	}
)

// MdbConfigMap 连接信息map
type MdbConfigMap map[string]*mdbBaseConfig

// MdbConfigs 需初始化的多个连接配置
var mdbConfigs = map[string]*mdbBaseConfig{
	//"test": mdbTest1,
	//"meidai": mdbMeidai,
}

// GetMdbConfigs 获取mysql配置信息
func GetMdbConfigs() MdbConfigMap {
	return mdbConfigs
}
