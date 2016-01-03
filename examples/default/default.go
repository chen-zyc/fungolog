package main

import (
	"github.com/zhangyuchen0411/fungolog"
	"github.com/zhangyuchen0411/fungolog/logger"
)

func main() {
	l := logger.NewLogger(fungolog.Info)
	l.Debugln("Debug Message")
	l.Infoln("Info Message")
	l.Warningln("Warning Message")
	l.Errorln("Error Message")
	l.Panicln("Panic Message")
	l.Write(fungolog.Info, "Message from Write()\n")
}
