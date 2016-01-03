package fungolog

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args...interface{})
	Debugln(args ...interface{})

	Info(args ...interface{})
	Infof(format string, args...interface{})
	Infoln(args ...interface{})

	Warning(args ...interface{})
	Warningf(format string, args...interface{})
	Warningln(args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args...interface{})
	Errorln(args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args...interface{})
	Panicln(args ...interface{})

	Write(Level, ...interface{})
}
