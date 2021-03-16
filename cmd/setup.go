package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up build output folders",
	Long: `Sets up build output folders with the correct permissions
so that latexmk has the correct privileges to write to the folders.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, file := range Files {
			os.Mkdir(file.FilePath+"/build", 0777)
			os.Mkdir(file.FilePath+"/build/tikz", 0777)
			os.Mkdir(file.FilePath+"/tikz", 0777)
		}

		os.Mkdir("build", 0777)
		os.Mkdir("build/tikz", 0777)
		os.Mkdir("tikz", 0777)
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
