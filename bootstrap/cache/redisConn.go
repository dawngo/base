package cache

import (
	"fmt"
	"github.com/Brave-man/base/config"
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"time"
)

func NewRedisConn(host, password string, port, db, maxRetries, dialTimeout, poolSize int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         host + ":" + strconv.Itoa(port),          // 连接地址 host:port
		Password:     password,                                 // 密码
		DB:           db,                                       // 待连接的db
		MaxRetries:   maxRetries,                               // 最大尝试次数
		DialTimeout:  time.Duration(dialTimeout) * time.Second, // 连接超时时间
		ReadTimeout:  5 * time.Second,                          // 读取超时时间
		WriteTimeout: 5 * time.Second,                          // 写入超时时间
		PoolSize:     poolSize,                                 // 池大小
	})
	return rdb
}

// RDB redis连接池
type RDB map[string]*redis.Client

// GetSqlDBMap 获取redis连接池
func GetRDBMap() RDB {
	RDBMap := RDB{}
	for name, rdbCfg := range config.GetRdbConfigs() {
		rdb := NewRedisConn(rdbCfg.Host, rdbCfg.Password, rdbCfg.Port, rdbCfg.DB, rdbCfg.MaxRetries,
			rdbCfg.DialTimeout, rdbCfg.MaxRetries)

		pong, err := rdb.Ping().Result()
		logStr := fmt.Sprintf("Connection (Redis: %s): host:%s; port:%d; db:%d; password:%s ==== ",
			name, rdbCfg.Host, rdbCfg.Port, rdbCfg.DB, rdbCfg.Password)
		if err != nil {
			log.Println(err)
			log.Fatal(logStr + "fail")
		}
		log.Println(logStr+"success; ping-->", pong)
		RDBMap[name] = rdb
	}
	return RDBMap
}

// Close 关闭redis连接池
func (d RDB) Close() {
	for _, rdb := range d {
		err := rdb.Close()
		log.Println("cache redis close error: ", err)
	}
}
