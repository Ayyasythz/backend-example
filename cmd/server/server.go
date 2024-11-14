package server

import (
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"sagara_backend_test/config"
	"sagara_backend_test/internal/handler/api"
	"sagara_backend_test/lib/database/sql"
	"sagara_backend_test/lib/log"
	"syscall"
)

var (
	serveHTTPCmd = &cobra.Command{
		Use:   "serve-http",
		Short: "Wardrobe System",
		Long:  "asd",
		RunE:  run,
	}
)

func ServeHTTPCmd() *cobra.Command {
	serveHTTPCmd.Flags().StringP("config", "c", "", "Config Path, both relative or absolute. i.e: /usr/local/bin/config/files")
	return serveHTTPCmd
}

func run(cmd *cobra.Command, args []string) error {
	configLocation, _ := cmd.Flags().GetString("config")

	cfg := &config.MainConfig{}
	config.ReadConfig(cfg, configLocation)

	database := sql.New(sql.DBConfig{
		SlaveDSN:        cfg.Database.SlaveDSN,
		MasterDSN:       cfg.Database.MasterDSN,
		RetryInterval:   cfg.Database.RetryInterval,
		MaxIdleConn:     cfg.Database.MaxIdleConn,
		MaxConn:         cfg.Database.MaxConn,
		ConnMaxLifetime: cfg.Database.ConnMaxLifetime,
	}, sql.DriverPostgres)

	appContainer := newContainer(&options{
		Cfg: cfg,
		DB:  database,
	})

	server := api.New(&api.Options{
		Cfg:        appContainer.Cfg,
		WardrobeUc: appContainer.WardrobeUc,
	})

	go server.Run()

	term := make(chan os.Signal)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		log.Info("Exiting gracefully...")
	case err := <-server.ListenError():
		log.Error("Error starting web server, exiting gracefully:", err)
	}

	return nil
}
