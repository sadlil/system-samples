// Service is a Demo service that dumps periodical log line with the timestamp
// and reports a promethues metrics.
package main

import (
	"context"
	"time"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"sadlil.com/samples/golib/application"
	"sadlil.com/samples/golib/server/statserver"
)

var (
	counterMetric = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "samples",
		Subsystem: "periscope",
		Name:      "service_counter",
	})
)

func init() {
	prometheus.MustRegister(counterMetric)
}

func main() {
	application.Init()

	ctx, cancel := context.WithCancel(context.Background())

	// We are only going to run a metric server for this exercise
	srv := statserver.New(
		statserver.WithMonitoringAddr(":6443"),
		statserver.WithProfiling(true),
		statserver.WithPromMetric(true),
	)
	srv.Start(ctx)

	go func() {
		for {
			select {
			case <-time.After(5 * time.Second):
				glog.Infof("Logging from Periscope: timestamp %v", time.Now().String())
				counterMetric.Add(1)
			case <-time.After(15 * time.Second):
				glog.Infof("Alertable log: timestamp %v", time.Now().String())
				counterMetric.Add(1)
			case <-ctx.Done():
				return
			}
		}
	}()

	// Wait until Application recives a Terminate signals.
	// Once the SIGKILL/SIGTERM recived we should
	// - stop servers,
	// - stop database,
	// - cancel context just to be sure anyone depending on the
	// context is notified.
	application.ShutdownOnInterrupt(func() {
		srv.Stop(ctx)
		cancel()
	})

}
