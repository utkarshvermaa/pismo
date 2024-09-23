package cmd

import (
	"context"
	"log"
	"time"

	"github.com/projects/sys-des/txn-routine/pkg/logger"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:           "app",
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger.GetLogger().Info("[%v] App start: %v", cmd.Use, time.Now().Format("2006-January-02 15:04:05"))
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		log.Printf("[%v] App close: %v", cmd.Use, time.Now().Format("2006-January-02 15:04:05"))
	},
}

func init() {
	RootCmd.SetContext(context.TODO())
	RootCmd.AddCommand(startServer)
}
