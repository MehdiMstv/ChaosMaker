package main

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // web framework adapter
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // sql driver
	"github.com/GoAdminGroup/go-admin/modules/language"
	_ "github.com/GoAdminGroup/themes/adminlte" // ui theme
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"github.com/MehdiMstv/ChaosMaker/src/admin/internal/forms/flagadmin"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server",
	Run:   serve,
}

const (
	serviceName = "chaos_maker_panel"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, _ []string) {
	r := gin.Default()
	eng := engine.Default()

	cfg := &config.Config{
		Databases: provideDatabaseConfig(),
		UrlPrefix: "chaos",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language: language.EN,
	}

	_ = eng.AddConfig(cfg).
		Use(r)

	r.GET("api/flags", func(context *gin.Context) {
		flagadmin.GetFlagsByService(context, eng.SqliteConnection())
	})

	_ = r.Run(":9033")
}

func provideDatabaseConfig() config.DatabaseList {
	return config.DatabaseList{
		"default": config.Database{
			Name:   "chaos",
			Driver: "sqlite",
			File:   "./src/admin/chaos.db",
		}}
}
