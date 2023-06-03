package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Cleans up temporary build and output files",
	Long:  "Cleans up temporary build and output files",
	Run: func(cmd *cobra.Command, args []string) {
		os.RemoveAll("build")
		os.RemoveAll("tikz")

		for _, file := range Files {
			os.RemoveAll(file.FilePath + "/build")
			os.RemoveAll(file.FilePath + "/tikz")
		}
		for _, file := range BookFiles {
			os.RemoveAll(file.FilePath + "/build")
			os.RemoveAll(file.FilePath + "/tikz")
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
