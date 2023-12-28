package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chaos serve",
	Short: "serves chaos admin panel",
	Run:   nil,
}

func init() {
	cobra.OnInitialize()
}
