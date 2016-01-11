package writers

import (
	"fmt"
	"github.com/zhangyuchen0411/fungolog"
	"testing"
	"time"
)

func TestMultiFileWriterFunc(t *testing.T) {
	nameGenerate := func(suffix string) FileNameGenerator {
		return FileNameGenerator(func() string {
			return "test_" + suffix + ".log"
		})
	}
	w := NewMultiFileWriter(map[fungolog.Level]MultiFileWriterSetting{
		fungolog.Panic: MultiFileWriterSetting{
			NameGenerator: nameGenerate(fungolog.Panic.String()),
			Async:         false,
		},
		fungolog.Error: MultiFileWriterSetting{
			NameGenerator: nameGenerate(fungolog.Error.String()),
			Async:         false,
		},
		fungolog.Info: MultiFileWriterSetting{
			NameGenerator: nameGenerate(fungolog.Info.String()),
			Async:         true,
		},
	}, MultiFileWriterSetting{
		NameGenerator: nameGenerate("other"),
		Async:         true,
	}, MultiFileWriterSetting{
		NameGenerator: nameGenerate("all"),
		Async:         true,
	})

	f := NewMultiFileWriteFunc(w)

	for _, level := range []fungolog.Level{fungolog.Debug, fungolog.Info, fungolog.Warning, fungolog.Error, fungolog.Panic} {
		f([]byte(fmt.Sprintf("message, level is %s\n", level.String())), level)
	}

	time.Sleep(time.Second)
}
