package main

import (
	"fmt"
	"os"

	"github.com/projects/sys-des/txn-routine/cmd"
	"github.com/projects/sys-des/txn-routine/pkg/logger"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		logger.GetLogger().ErrorContext(cmd.RootCmd.Context(), fmt.Sprintf("Error: %v", err))
		os.Exit(-1)
	}
}
