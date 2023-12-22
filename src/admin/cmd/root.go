package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "admin serve",
	Short: "serves admin panel of chaos maker",
	Run:   nil,
}

func init() {
	cobra.OnInitialize()
}
