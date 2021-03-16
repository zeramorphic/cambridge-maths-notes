package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// fmtCmd represents the fmt command
var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Uses latexindent to automatically format LaTeX files",
	Long:  "Uses latexindent to automatically format LaTeX files",
	Run: func(cmd *cobra.Command, args []string) {
		format()
	},
}

// https://stackoverflow.com/a/55300382/9420290
func walkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func init() {
	rootCmd.AddCommand(fmtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fmtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fmtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func format() {
	fmt.Println("Searching for LaTeX sources...")
	files, _ := walkMatch(".", "*.tex")
	fileChan := make(chan string)
	go func() {
		for _, v := range files {
			fileChan <- v
		}
	}()
	fmt.Println("Formatting...")
	workers := 6

	var wg sync.WaitGroup

	bar := progressbar.Default(int64(len(files)))
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for {
				shouldBreak := false
				// Check if there's a new item on the channel.
				select {
				case file := <-fileChan:
					// Format the file.
					input := file
					output := file + ".tmp.tex"
					cmd := exec.Command("latexindent", "-s", input, "-o", output)
					cmd.Start()
					cmd.Wait()
					os.Remove(input)
					os.Rename(output, input)
					bar.Add(1)

				default:
					// Check if the progress bar is full.
					// If so, we're done.
					if bar.IsFinished() {
						shouldBreak = true
					}
				}

				if shouldBreak {
					break
				}
			}

			wg.Done()
		}()
	}

	wg.Wait()
	color.HiGreen("Done!")
}
