package main

import (
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"github.com/MehdiMstv/ChaosMaker/internal/forms/flagadmin"
	"github.com/MehdiMstv/ChaosMaker/internal/forms/serviceadmin"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server",
	Run:   serve,
}

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

	serviceGenerator := &serviceadmin.ServiceGenerator{}
	eng.AddGenerators(serviceGenerator.GetGenerator())
	flagGenerator := &flagadmin.FlagsGenerator{}
	eng.AddGenerators(flagGenerator.GetGenerator())

	_ = r.Run(":9033")
}

func provideDatabaseConfig() config.DatabaseList {
	return config.DatabaseList{
		"default": config.Database{
			Name:   "chaos",
			Driver: "sqlite",
			File:   "./chaos.db",
		}}
}
