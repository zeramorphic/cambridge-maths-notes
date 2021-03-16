package cmd

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds all LaTeX sources in parallel.",
	Long:  "Builds all LaTeX sources in parallel.",
	Run: func(cmd *cobra.Command, args []string) {
		compileBook, _ := cmd.Flags().GetBool("book")
		build(compileBook)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	buildCmd.Flags().BoolP("book", "b", false, "Compile the book as well as the individual courses")
}

func build(compileBook bool) {
	fmt.Println("Building...")
	var wg sync.WaitGroup

	size := int64(len(Files))
	if compileBook {
		size++
	}
	bar := progressbar.Default(size)
	for _, file := range Files {
		wg.Add(1)
		file := file
		go func() {
			defer wg.Done()

			// Compile the file.
			cmd := exec.Command("latexmk", "--shell-escape", "-pdf", "-cd", "-output-directory=build", "-file-line-error", "-halt-on-error", "-interaction=nonstopmode", file.FilePath+"/main.tex")
			cmd.Start()
			cmd.Wait()
			bar.Add(1)
		}()
	}
	if compileBook {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Compile the file.
			cmd := exec.Command("latexmk", "--shell-escape", "-pdf", "-cd", "-output-directory=build", "-file-line-error", "-halt-on-error", "-interaction=nonstopmode", "book.tex")
			cmd.Start()
			cmd.Wait()
			bar.Add(1)
		}()
	}

	wg.Wait()
	color.HiGreen("Done!")
}
