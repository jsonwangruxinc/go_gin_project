package logs

import "io"

// 定义map集合，根据字符串，找一个初始化的方法
var writerAdapter = make(map[string]InitLogWriterFunc, 0)

// 定义func类型,返回LogWriter
type InitLogWriterFunc func() LogWriter

// 定义log日志类型，Flush 关闭网络连接，关闭文件连接
type LogWriter interface {
	Flush()
	io.Writer
}

// 全局注册方法
func RegisterInitWriterFunc(adapterName string, writerFunc InitLogWriterFunc) {
	writerAdapter[adapterName] = writerFunc
}
