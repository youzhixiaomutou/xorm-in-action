package main

import (
	"fmt"
	"time"

	"git.nd.com.cn/go-common/library/log"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"xorm.io/xorm"

	"git.nd.com.cn/go-common/library/database/sql/mysql"
)

var (
	driver = "tq_mysql"
	dsn    = "root:lin123@tcp(192.168.52.164:3306)/test?charset=utf8mb4&interpolateParams=true&parseTime=True&loc=Local&allowOldPasswords=true"
)

var engine *xorm.Engine

func main() {
	fmt.Println("xorm-in-action run.")

	// 在 8080 端口启动 /metrics 端点用来测试监控指标
	e := gin.New()
	prometheus.MustRegister(mysql.MetricReqErr, mysql.MetricReqDur, mysql.MetricConnCurrent)
	e.GET("/metrics", func(context *gin.Context) {
		promhttp.Handler().ServeHTTP(context.Writer, context.Request)
	})
	go func() {
		_ = e.Run(":8080")
	}()

	logger, _, _ := log.NewLogger()
	// 注册 tq_mysql 驱动
	_, _, _ = mysql.NewDB(
		mysql.Logger(logger),

		mysql.DSN(dsn),

		// 给这个 MySQL 实例取个名字，默认为空（监控时可以通过 name 来区分不同实例）
		mysql.Name("test"),
		// 是否错误日志记录（默认为 true）
		mysql.LogErr(true),
		// logSlow 慢查询日志记录时间（默认 250 毫秒）
		mysql.LogSlow(250*time.Millisecond),
		// metrics 是否提供指标信息（默认为 true）
		mysql.Metrics(true),
	)

	var err error
	engine, err = xorm.NewEngine(driver, dsn)
	if err != nil {
		panic(err)
	}

	err = engine.Ping()
	if err != nil {
		panic(err)
	}

	// xorm 支持获取表结构信息，通过调用 engine.DBMetas() 可以获取到数据库中所有的表，字段，索引的信息
	metas, err := engine.DBMetas()
	if err != nil {
		panic(err)
	}
	for _, meta := range metas {
		fmt.Printf("meta.Name=%s\n", meta.Name)
	}

	// 根据传入的结构体指针及其对应的 Tag，提取出模型对应的表结构信息。这里不是数据库当前的表结构信息，而是我们通过 struct 建模时希望数据库的表的结构信息
	info, err := engine.TableInfo(xUser{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", info)

	// 表结构操作
	//table()
	//index()
	//sync()

	for {
		// 查询和统计数据
		query()
		time.Sleep(3 * time.Second)
	}
}
