package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"stock/pkg/service"
	"stock/pkg/transport"
	"syscall"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func main() {
	// Flag
	var (
		httpAddr          = flag.String("http", ":9292", "http listen address")
		logOutput         = flag.String("log.output", "console", "Output log")
		profile           = flag.String("profile", "dev", "Profile app environment")
		logOutputFilePath = flag.String("log.output.file.path", ".", "Output log file path")
	)
	flag.Parse()

	// Log
	logfile, err := os.OpenFile(*logOutputFilePath+"/stock.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logfile.Close()
	var logger log.Logger
	{
		if *logOutput == "file" {
			w := log.NewSyncWriter(logfile)
			logger = log.NewLogfmtLogger(w)
		} else {
			logger = log.NewLogfmtLogger(os.Stderr)
		}

		if *profile == "prod" {
			logger = level.NewFilter(logger, level.AllowInfo())
		} else {
			logger = level.NewFilter(logger, level.AllowAll())
		}

		logger = level.NewInjector(logger, level.InfoValue())
		logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	}

	// Database
	db, err := sql.Open("mysql", "root:@/stock_management")
	if err != nil {
		level.Error(logger).Log("err", err.Error())
	}
	level.Info(logger).Log("db status", "connected")

	var s service.Service
	{
		s = service.NewServiceImplV1(logger, db)
	}

	var h http.Handler
	{
		h = transport.MakeHTTPHandler(s, log.With(logger, "component", "HTTP"))
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	level.Info(logger).Log("exit", <-errs)
}
