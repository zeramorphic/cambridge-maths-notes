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

// optimiseCmd represents the optimise command
var optimiseCmd = &cobra.Command{
	Use:   "optimise",
	Short: "Optimises output PDF files",
	Long:  "Optimises output PDF files using GhostScript",
	Run: func(cmd *cobra.Command, args []string) {
		optimise()
	},
}

func init() {
	rootCmd.AddCommand(optimiseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// optimiseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// optimiseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func optimise() {
	fmt.Println("Building...")
	var wg sync.WaitGroup

	bar := progressbar.Default(int64(len(FilesWithBook)))
	for _, file := range FilesWithBook {
		wg.Add(1)
		file := file
		go func() {
			defer wg.Done()

			for {
				// Compile the file.
				output := file.FilePath + "/build/main_opt.pdf"
				input := file.FilePath + "/build/main.pdf"
				if file.FilePath == "book" {
					output = "build/book_opt.pdf"
					input = "build/book.pdf"
				}
				cmd := exec.Command("gs", "-sDEVICE=pdfwrite", "-dCompatibilityLevel=1.4", "-dNOPAUSE", "-dQUIET", "-dBATCH", "-sOutputFile="+output, input)
				cmd.Start()
				cmd.Wait()

				ostat, _ := os.Stat(output)
				istat, _ := os.Stat(input)

				if ostat.Size() >= istat.Size() {
					os.Remove(output)
					break
				}
				os.Rename(output, input)
			}

			bar.Add(1)
		}()
	}

	wg.Wait()
	color.HiGreen("Done!")
}
