package main

// http Gut Post Put 监听端口开发配置
//
//	func main() {
//		LoggerMiddleware(InitDB)("tcp://127.0.0.1")
//		r := gin.New()
//		r.Use(gin.Logger(), gin.Recovery())
//		routers.InitRouter(r)
//		r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//	}
//
//	func CheckAuth(param string) gin.HandlerFunc {
//		return func(c *gin.Context) {
//			fmt.Println("call check Auth func", param)
//			c.Next()
//		}
//	}
//
//	func InitDB(connstr string) {
//		fmt.Println("初始化数据库", connstr)
//	}
//
//	func LoggerMiddleware(in func(connstr string)) func(connstr string) {
//		return func(connstr string) {
//			fmt.Println("call LoggerMiddleware start")
//			in(connstr)
//			fmt.Println("call LoggerMiddleware end")
//		}
//
// }
// gorm数据库使用方法
//
//	func main() {
//		grom.CreateRecord()
//		grom.SelcetCreate()
//		grom.OmitCreate()
//		grom.BatchCreate()
//		grom.GetOnce()
//		grom.GetByPrimarKey()
//		grom.GetByConnt()
//	}
//log日志方法 集合

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go_gin_project/logrus/logs"
	"net/http"
	"time"
)

var log *logs.Log

func init() {
	conf := logs.LogConf{
		Level:       logrus.InfoLevel,
		AdapterName: "fileRotate",
	}
	log = logs.InitLog(conf)
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//log.SetOutput(w())

	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)
}

func main() {
	r := gin.New()
	r.Use(MyLoger)
	//r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	r.Run()
}

func MyLoger(ctx *gin.Context) {
	start := time.Now()
	path := ctx.Request.URL.Path
	raw := ctx.Request.URL.RawQuery
	// 执行下一个中间件
	ctx.Next()
	mp := make(map[string]interface{})
	mp["path"] = path
	mp["Client_ip"] = ctx.ClientIP()
	mp["method"] = ctx.Request.Method
	mp["Status_code"] = ctx.Writer.Status()
	mp["size"] = ctx.Writer.Size()
	if raw != "" {
		mp["path"] = path + "?" + raw
	}
	mp["latency"] = time.Since(start).Milliseconds()
	log.WithFields(mp).Info()
}

//func main() {
//	defer func() {
//		log.Flush()
//	}()
//	log.Info()
//	log.WithFields(logrus.Fields{
//		"animal": "walrus",
//		"size":   10,
//	}).Info("A group of walrus emerges from the ocean")
//
//	log.WithFields(logrus.Fields{
//		"omg":    true,
//		"number": 122,
//	}).Warn("The group's number increased tremendously!")
//
//	log.WithFields(logrus.Fields{
//		"omg":    true,
//		"number": 100,
//	}).Fatal("The ice breaks!")
//
//	// A common pattern is to re-use fields between logging statements by re-using
//	// the logrus.Entry returned from WithFields()
//	contextLogger := log.WithFields(logrus.Fields{
//		"common": "this is a common field",
//		"other":  "I also should be logged always",
//	})
//
//	contextLogger.Info("I'll be logged with common and other field")
//	contextLogger.Info("Me too")
//}
