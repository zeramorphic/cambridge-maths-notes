package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// htmlCmd represents the optimise command
var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "Creates output HTML files",
	Long:  "Creates output HTML files",
	Run: func(cmd *cobra.Command, args []string) {
		html()
	},
}

func init() {
	rootCmd.AddCommand(htmlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// optimiseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// optimiseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func html() {
	fmt.Println("Building...")
	var wg sync.WaitGroup

	size := int64(len(Files))
	errorChan := make(chan string, size)
	bar := progressbar.Default(size)
	for _, file := range Files {
		wg.Add(1)
		file := file
		f := func() {
			defer wg.Done()

			// Compile the file.
			os.MkdirAll("html/build/"+file.FilePath, 0777)
			cmd := exec.Command(
				"pandoc",
				"main.tex",
				"--standalone",
				"--mathjax",
				"--include-in-header", "../../html/header.html",
				"--output", "../../html/build/"+file.FilePath+"/index.html",
			)
			cmd.Dir = file.FilePath
			out, err := cmd.CombinedOutput()
			if err != nil {
				errorChan <- file.FilePath + ":\n" + string(out)
				bar.Set(bar.GetMax())
			}
			bar.Add(1)
		}
		go f()
	}

	wg.Wait()
	close(errorChan)

	hasError := false
	for v := range errorChan {
		hasError = true
		color.HiRed("Error raised while compiling:")
		fmt.Println(v)
	}

	if !hasError {
		color.HiGreen("Done!")
	}
}
