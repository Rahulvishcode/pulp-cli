package cmd

import (
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

var API string = "https://p.aadi.lol/api/"
var Client *resty.Request = resty.New().R()

var rootCmd = &cobra.Command{
	Use:   "pulp",
	Short: "Pulp - share code seamlessly and effectively!",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
