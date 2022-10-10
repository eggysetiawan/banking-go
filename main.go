package main

import (
	"github.com/eggysetiawan/banking-go/app"
	"github.com/eggysetiawan/banking-go/errs/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
