package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Brave-man/base/config"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// NewConn 创建一个数据库连接
func NewMysqlConn(host, username, password, database string, port int) (*gorm.DB, error) {
	cfg := mysql.Config{
		User:   username,
		Passwd: password,
		Net: "tcp",
		Addr:   host + ":" + strconv.Itoa(port),
		DBName: database,
		AllowNativePasswords:true,
	}
	dsn := cfg.FormatDSN()

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	//SetMaxOpenConns用于设置最大打开的连接数
	//SetMaxIdleConns用于设置闲置的连接数
	db.DB().SetMaxIdleConns(30)
	db.DB().SetMaxOpenConns(30)

	// 启用Logger，显示详细日志
	db.LogMode(true)

	return db, nil
}

// DBSql mysql连接池
type DBSql map[string]*gorm.DB

// GetSqlDBMap 获取数据库连接池
func GetSqlDBMap() DBSql {
	sqlDBMap := DBSql{}
	for name, sqlCfg := range config.GetMdbConfigs() {
		db, err := NewMysqlConn(sqlCfg.Host, sqlCfg.Username, sqlCfg.Password, sqlCfg.Database, sqlCfg.Port)
		logStr := fmt.Sprintf("Connection (Mysql/Maria: %s): host:%s; port:%d; username:%s; password:%s ==== ",
			sqlCfg.Database, sqlCfg.Host, sqlCfg.Port, sqlCfg.Username, sqlCfg.Password)
		if err != nil {
			log.Println(err)
			log.Fatal(logStr + "fail")
		}
		log.Println(logStr + "success")
		sqlDBMap[name] = db
	}
	return sqlDBMap
}

// Close 关闭连接池
func (d DBSql) Close() {
	for _, db := range d {
		err := db.Close()
		log.Println("db mysql close error: ", err)
	}
}
