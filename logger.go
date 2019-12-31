package main

// 声明一个日志写入器
type LogWriter interface {
	Write(data interface{}) error
}

// 声明日志器
type Logger struct {
	writerList []LogWriter
}

// 注册一个日志写入器
func (l *Logger) RegisterWriter(writer LogWriter)  {
	l.writerList = append(l.writerList, writer)
}

// 将一个data类型的数据写入日志
func (l *Logger)  Log(data interface{}){

	// 遍历所有注册写入器
	for _, writer := range l.writerList{

		// 将日志
		writer.Write(data)
	}
}

// 创建日志的实例
func NewLogger()  *Logger{
	return &Logger{}
}
