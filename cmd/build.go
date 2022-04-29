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
		singleThreaded, _ := cmd.Flags().GetBool("single-threaded")
		hideProofs, _ := cmd.Flags().GetBool("hide-proofs")
		goMode, _ := cmd.Flags().GetBool("go")
		build(compileBook, singleThreaded, hideProofs, goMode)
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
	buildCmd.Flags().BoolP("single-threaded", "s", false, "Compile only one file at once")
	buildCmd.Flags().BoolP("hide-proofs", "p", false, "Hide all instances of the proof environment")
	buildCmd.Flags().BoolP("go", "g", false, "Force latexmk to process document, even if it can't detect changes (useful when enabling or disabling -p)")
}

func build(compileBook bool, singleThreaded bool, hideProofs bool, goMode bool) {
	fmt.Println("Building...")
	var wg sync.WaitGroup

	size := int64(len(Files))
	errorChan := make(chan string, size)
	if compileBook {
		size++
	}
	bar := progressbar.Default(size)

	var fs []TexFile
	if compileBook {
		fs = FilesWithBook
	} else {
		fs = Files
	}

	for _, file := range fs {
		wg.Add(1)
		file := file
		f := func() {
			defer wg.Done()

			// Compile the file.
			cmd := exec.Command("latexmk", "--shell-escape", "-pdflatex=xelatex", "-pdf", "-cd", "-output-directory=build", "-file-line-error", "-halt-on-error", "-interaction=nonstopmode")
			if hideProofs {
				cmd.Args = append(cmd.Args, `-usepretex="\\def\\hideproofs{}"`)
			}
			if goMode {
				cmd.Args = append(cmd.Args, "-g")
			}
			cmd.Args = append(cmd.Args, file.FilePath+"/main.tex")
			out, err := cmd.CombinedOutput()
			if err != nil {
				errorChan <- string(out)
				bar.Set(bar.GetMax())
			}
			bar.Add(1)
		}
		if singleThreaded {
			f()
		} else {
			go f()
		}
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
