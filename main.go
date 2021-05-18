package main

import (
	"github.com/quannguyennn/banking/app"
	"github.com/quannguyennn/banking/logger"
)

func main() {
	logger.Info("Starting application...")
	app.Start()
}
