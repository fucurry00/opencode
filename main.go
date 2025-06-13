package main

import (
	"github.com/fucurry00/opencode/cmd"
	"github.com/fucurry00/opencode/internal/logging"
)

func main() {
	defer logging.RecoverPanic("main", func() {
		logging.ErrorPersist("Application terminated due to unhandled panic")
	})

	cmd.Execute()
}
