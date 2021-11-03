package cmd

import (
	"bytes"
	"errors"
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

func exists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
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
			for i := 1; exists(file.FilePath + "/" + fmt.Sprintf("%02d", i) + ".tex"); i++ {
				iStr := fmt.Sprintf("%02d", i)
				rawFileContents, _ := os.ReadFile(file.FilePath + "/" + iStr + ".tex")
				fileString := string(rawFileContents)

				// We can modify the contents of the file that pandoc sees here.
				fileString = "\\input{../../util.tex}\\n" + fileString

				cmd := exec.Command(
					"pandoc",
					"--from", "latex",
					"--standalone",
					"--mathjax",
					"--include-in-header", "../../html/header.html",
					"--css=../../../normalize.css",
					"--css=../../../main.css",
					"--output", "../../html/build/"+file.FilePath+"/"+iStr+".html",
				)
				cmd.Dir = file.FilePath
				var out bytes.Buffer
				var err bytes.Buffer
				cmd.Stdout = &out
				cmd.Stderr = &err
				stdin, _ := cmd.StdinPipe()
				cmd.Start()
				stdin.Write([]byte(fileString))
				stdin.Close()
				if cmd.Wait() != nil {
					errorChan <- file.FilePath + "#" + iStr + ":\n" + out.String()
					bar.Set(bar.GetMax())
				}
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
