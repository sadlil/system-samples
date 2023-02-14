package persistent

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang/glog"
	"gorm.io/gorm/logger"
)

func newAppLogger() logger.Interface {
	return logger.New(&appLogWriter{}, logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
	})
}

type appLogWriter struct{}

func (*appLogWriter) Printf(s string, v ...interface{}) {
	s = strings.Replace(s, "\n", " ", -1)
	sourceFile := v[0].(string)
	v[0] = sourceFile[strings.LastIndex(sourceFile, "/")+1:]
	glog.InfoDepth(1, fmt.Sprintf(s, v...))
}
