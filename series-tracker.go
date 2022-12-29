package main

import (
	"fmt"
	"github.com/erayaydin/series-tracker/cmd/server"
	"github.com/erayaydin/series-tracker/pkg/di"
	"github.com/erayaydin/series-tracker/pkg/logger"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run() error {

	g := gin.Default()
	d, err := di.Build()
	if err != nil {
		return err
	}

	var l logger.LogInfoFormat
	err = di.Invoke(func(log logger.LogInfoFormat) {
		l = log
	})
	if err != nil {
		return err
	}

	svr := server.NewServer(g, d, l)
	svr.RegisterRoutes()
	return svr.Start()
}
