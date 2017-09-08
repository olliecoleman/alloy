package main

import (
	"runtime"

	"github.com/olliecoleman/alloy/app/services"
	"github.com/olliecoleman/alloy/app/views"
	"github.com/olliecoleman/alloy/cmd"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	// Connect to database
	db := services.InitDB()
	defer db.Close()

	// Setup templates
	views.LoadTemplates()

	// Get started
	cmd.Execute()
}
