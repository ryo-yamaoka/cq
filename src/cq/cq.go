package cq

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version     = "-1.-1.-1" // It will be set by -ldflags (look at Makefile)
	versionFlag bool         // --version -v
)

var rootCmd = &cobra.Command{
	Use:   "cq",
	Short: "cq is a tool for simple and fast control cloud environment.",
	Long:  "cq is a tool for simple and fast control cloud environment.\nhttps://github.com/ap-communications/cq",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Show cq version")
	rootCmd.AddCommand(versionCmd)
}

// Execute is root of cq
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show cq version",
	Long:  "Show cq version",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func printVersion() {
	fmt.Printf("Cloud Query (cq) version %s\n", version)
}
