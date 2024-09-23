package cmd

import (
	"os"
	"strconv"

	"github.com/projects/sys-des/txn-routine/internal/accounts"
	"github.com/projects/sys-des/txn-routine/internal/transactions"
	"github.com/projects/sys-des/txn-routine/migrations"
	"github.com/projects/sys-des/txn-routine/pkg/env"
	"github.com/projects/sys-des/txn-routine/pkg/logger"
	"github.com/projects/sys-des/txn-routine/server"
	"github.com/spf13/cobra"
)

var startServer = &cobra.Command{
	Use:   "start",
	Short: "Start the http server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		logger.GetLogger().InfoContext(ctx, "Starting server...")
		e := env.New(ctx)
		ctx = e.Setup(ctx)
		migrations.Run(ctx)

		p := os.Getenv("APP_PORT")
		port, err := strconv.ParseUint(p, 10, 64)
		if err != nil {
			logger.GetLogger().ErrorContext(ctx, "Failed to get port number from env")
			os.Exit(-1)
		}

		acc := accounts.New(ctx)
		txn := transactions.New(ctx)

		h := server.Handler(acc, txn)
		server.Start(ctx, h, port)
	},
}
