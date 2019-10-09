package database

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Brave-man/base/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// NewMongoConn 连接mongo
func NewMongoConn(c *config.MongoBaseConfig) (*mongo.Client, error) {
	cs := &options.ClientOptions{}

	// 设置认证
	// Supported values include "SCRAM-SHA-256", "SCRAM-SHA-1", "MONGODB-CR", "PLAIN", "GSSAPI", and "MONGODB-X509".
	if c.Auth {
		cs.SetAuth(options.Credential{
			AuthMechanism: c.Mechanism,
			Username:      c.Username,
			Password:      c.Password,
			PasswordSet:   c.PasswordSet,
		})
	}
	// 设置地址
	cs.SetHosts([]string{c.Host + ":" + strconv.Itoa(c.Port)})
	// 设置连接超时时间
	cs.SetConnectTimeout(time.Duration(c.Timeout) * time.Second)
	// 设置最大连接数
	cs.SetMaxPoolSize(uint64(c.MaxPoolSize))
	// 设置最小连接数
	cs.SetMinPoolSize(uint64(c.MinPoolSize))

	client, err := mongo.NewClient(cs)
	if err != nil {
		return client, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return client, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return client, err
	}
	return client, err
}

// MDB mongo连接池
type MDB map[string]*mongo.Client

// GetMDBMap 获取mongo连接池
func GetMDBMap() MDB {
	MDBMap := MDB{}
	for name, mgoCfg := range config.GetMgoConfigs() {
		client, err := NewMongoConn(mgoCfg)

		logStr := fmt.Sprintf("Connection (Mongo: %s): host:%s; port:%d; user:%s; password:%s ==== ",
			name, mgoCfg.Host, mgoCfg.Port, mgoCfg.Username, mgoCfg.Password)

		if err != nil {
			log.Println(err)
			log.Fatal(logStr + "fail")
		}

		log.Println(logStr + "success")
		MDBMap[name] = client
	}
	return MDBMap
}

// Close 关闭连接池
func (d MDB) Close() {
	for _, mdb := range d {
		ctx := context.Background()
		err := mdb.Disconnect(ctx)
		log.Println("db mongo close error: ", err)
	}
}
