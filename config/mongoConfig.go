package config

// mongoBaseConfig mongo基础配置
type MongoBaseConfig struct {
	Host        string // 主机
	Port        int    // 端口
	Auth        bool   // 是否开启认证
	Mechanism   string // 认证方式
	Username    string // 用户名
	Password    string // 密码
	PasswordSet bool   // 是否设置密码("" 被认为是密码)
	Timeout     int    // 超时时间，单位秒
	MaxPoolSize int    // 最大连接数
	MinPoolSize int    // 最小连接数
}

var (
	mgoTest1 = &MongoBaseConfig{
		Host:        "10.10.10.233",
		Port:        27017,
		Auth:false,
		Username:    "",
		Password:    "",
		PasswordSet: false,
		Timeout:     10,
		MaxPoolSize: 200,
		MinPoolSize: 10,
	}

	mgoTest2 = &MongoBaseConfig{
		Host:        "10.10.10.233",
		Port:        27017,
		Auth:false,
		Username:    "",
		Password:    "",
		PasswordSet: false,
		Timeout:     10,
		MaxPoolSize: 200,
		MinPoolSize: 10,
	}
)

// MgoConfigMap 连接信息map
type MgoConfigMap map[string]*MongoBaseConfig

// MdbConfigs 需初始化的多个连接配置
var mgoConfigs = MgoConfigMap{
	//"test1": mgoTest1,
	//"test2": mgoTest2,
}

// GetMgoConfigs 获取mongo配置信息
func GetMgoConfigs() MgoConfigMap {
	return mgoConfigs
}
