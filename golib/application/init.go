package application

import (
	"flag"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"

	"github.com/golang/glog"
	"github.com/spf13/pflag"
)

var (
	once        sync.Once
	initialized uint32

	initFuncs []func()
)

// Init initializes the application
func Init() {
	if atomic.LoadUint32(&initialized) != 0 {
		return
	}

	atomic.StoreUint32(&initialized, 1)
	once.Do(func() {
		pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
		pflag.Parse()

		for _, f := range initFuncs {
			f()
		}
	})
}

func IsInitialized() bool {
	return atomic.LoadUint32(&initialized) == 1
}

func RegisterInit(f func()) {
	initFuncs = append(initFuncs, f)
}

func ShutdownOnInterrupt(shutdownFunc func()) {
	recoverFromPanic()
	glog.Flush()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	<-stop

	glog.Infoln("Received Close Signal, calling cancel")
	shutdownFunc()
	glog.Infof("Shutting down application")
}

func recoverFromPanic() {
	if r := recover(); r != nil {
		// Same as stdlib http server code. Manually allocate stack trace buffer size
		// to prevent excessively large logs
		const size = 64 << 10
		stacktrace := make([]byte, size)
		stacktrace = stacktrace[:runtime.Stack(stacktrace, false)]

		glog.Fatalf("Recovered from panic %q. Call stack:\n%s", r, stacktrace)
	}
}
