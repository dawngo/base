package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/Brave-man/base/bootstrap/cache"
	"github.com/Brave-man/base/bootstrap/database"
	_ "github.com/Brave-man/base/bootstrap/logger"
	"github.com/Brave-man/base/config"
	"github.com/Brave-man/base/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 设置goroutine 运行时cpu最大核数
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 解析命令行mode
	var mode string
	flag.StringVar(&mode, "mode", "debug", "mode must in (debug, test, release)")
	flag.Parse()

	// 设置模式
	gin.SetMode(mode)
	config.Conf.AppConfig.Set("mode", mode)

	// 初始化路由组
	router := routers.InitRouter()

	// 退出程序时，关闭各个数据库和缓存连接池
	defer func() {
		database.GlobalDBSql.Close()
		database.GlobalMDB.Close()
		cache.GlobalRDB.Close()
	}()

	// 获取地址
	addr := config.Conf.AppConfig.GetString(mode+".host") + ":" + config.Conf.AppConfig.GetString(mode+".port")
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	log.Printf("The current running mode is %s, address is %s.", mode, addr)

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
