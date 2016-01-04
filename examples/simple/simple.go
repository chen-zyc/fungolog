package main

import (
	"fmt"
	"github.com/zhangyuchen0411/fungolog"
	"github.com/zhangyuchen0411/fungolog/formatter"
	"github.com/zhangyuchen0411/fungolog/logger"
)

func main() {
	fmt.Println("++++++++++ SimpleFormatFullTime ++++++++++")
	l := logger.NewLogger(fungolog.Debug)
	l.Format = formatter.SimpleFormatFullTime
	l.Debugln("Debug Message")
	l.Infoln("Info Message")
	l.Warningln("Warning Message")
	l.Errorln("Error Message")
	l.Panicln("Panic Message")
	l.Write(fungolog.Info, "Message from Write()\n")

	fmt.Println("\n++++++++++ SimpleFormatColor ++++++++++")
	l.Format = formatter.SimpleFormatColor
	l.Debugln("Debug Message")
	l.Infoln("Info Message")
	l.Warningln("Warning Message")
	l.Errorln("Error Message")
	l.Panicln("Panic Message")
	l.Write(fungolog.Info, "Message from Write()\n")

	fmt.Println("\n++++++++++ SimpleFormatCallerInfo ++++++++++")
	l.Format = formatter.SimpleFormatCallerInfo
	l.Debugln("Debug Message")
	l.Infoln("Info Message")
	l.Warningln("Warning Message")
	l.Errorln("Error Message")
	l.Panicln("Panic Message")
	l.Write(fungolog.Info, "Message from Write()\n")

	fmt.Println("\n++++++++++ SimpleFormatFullCaller ++++++++++")
	l.Format = formatter.SimpleFormatFullCaller
	l.Debugln("Debug Message")
	l.Infoln("Info Message")
	l.Warningln("Warning Message")
	l.Errorln("Error Message")
	l.Panicln("Panic Message")
	l.Write(fungolog.Info, "Message from Write()\n")

	fmt.Println("\n++++++++++ SimpleFormatStackInfo ++++++++++")
	l.Format = formatter.SimpleFormatStackInfo
	l.Debugln("Debug Message")
	l.Infoln("Info Message")
	l.Warningln("Warning Message")
	l.Errorln("Error Message")
	l.Panicln("Panic Message")
	l.Write(fungolog.Info, "Message from Write()\n")

	fmt.Println("\n++++++++++ SimpleFormatCallerColor ++++++++++")
	l.Format = formatter.SimpleFormatCallerColor
	l.Debugln("Debug Message")
	l.Infoln("Info Message")
	l.Warningln("Warning Message")
	l.Errorln("Error Message")
	l.Panicln("Panic Message")
	l.Write(fungolog.Info, "Message from Write()\n")
}
