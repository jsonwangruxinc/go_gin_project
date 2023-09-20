package logs

import "os"

// 定义一个std struct类型
type stdWriter struct {
	//	内嵌一个文件
	*os.File
}

// 实现flush
func (s *stdWriter) Flush() {
	s.Sync()
}

func newStdWriter() LogWriter {
	return &stdWriter{os.Stderr}
}

// 在init方法里面注册Func
func init() {
	RegisterInitWriterFunc("std", newStdWriter)
}
