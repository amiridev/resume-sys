package main

import (
	"resume-sys/core/cmd"
)

// @securityDefinitions.apikey Bearer
// @in 						   header
// @name 					   Authorization
func main() {
	cmd.Execute()
}
