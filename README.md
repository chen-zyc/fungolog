# fungolog

## 简单使用

```go
	l := logger.NewLogger(fungolog.Info)
	l.Debugln("Debug Message")
	l.Infoln("Info Message")
	l.Warningln("Warning Message")
	l.Errorln("Error Message")
	l.Panicln("Panic Message")
	l.Write(fungolog.Info, "Message from Write()\n")
```

上面这个例子在 `examples/default/default.go` 中。
 

## 自定义

Logger中各字段都是公开的，这意味着你可以根据自己的需要更改Logger的默认实现。

### 自定义Formatter

```go
	l := logger.NewLogger(fungolog.Debug)
	l.Format = formatter.SimpleFormatFullTime
	l.Debugln("Debug Message")
	l.Infoln("Info Message")
	l.Warningln("Warning Message")
	l.Errorln("Error Message")
	l.Panicln("Panic Message")
	l.Write(fungolog.Info, "Message from Write()\n")
```

上述代码见 `examples/simple.simple.go`

`formatter.SimpleFormatFullTime` 是 formatter 包中内置的格式化函数。

formatter 包中内置了几种简单的格式化函数，也提供了一些格式化时可以使用的工具函数，比如调用者信息，栈信息，时间格式化等。

如果需要实现自己特定的格式化函数，需要返回一个如下的函数原型：

```go
func(buf *bytes.Buffer, level fungolog.Level, args ...interface{})
```

level 是当前消息等级，所以等级可参考 level.go 文件。

args 是传入的参数。

buf 格式化后的内容需要写入到这个buf中。

### 自定义WriteFunc

WriteFunc的原型是：

```go
func([]byte, fungolog.Level)
```

第一个参数是格式化后的内容，第二个参数是等级。

默认情况下， 日志时写入到 `os.Stderr` 的。

如果希望将日志写入到文件中， `examples/write_file/write_file.go` 中有一个粗糙的例子可以参考。

另外，writers 包中还提供了以下几种文件写入模式：

#### FileWriter 

该结构体需要传入一个FileNameGenerator类型的函数，每次调用FileWriter.Write()方法时都会调用该函数获取写入文件的名称。

如果两次调用FileNameGenerator时返回的文件名不同，FileWriter会自动关闭上一个文件并打开新的文件写入（如果目录不存在则自动创建）。

#### AsyncFileWriter

AsyncFileWriter 封装了 FileWriter， 但是AsyncFileWriter的写入是异步的。

#### MultiFileWriter

MultiFileWriter 的目的是可以为每个level制定不同的文件，并可以指定每个文件的写入是否是异步的。

如果没有为某个level制定文件，那么MultiFileWriter会尝试写入到other指定的文件中（如果有的话）。

如果指定了all，则所有的日志都会写入到all指定的文件中。
