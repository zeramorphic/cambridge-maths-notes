package cmd

import (
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extracts definitions and theorems from files",
	Long:  "Extracts definitions and theorems from files. Can be used in conjunction with Anki to make flashcards semi-automatically.",
	Run: func(cmd *cobra.Command, args []string) {
		course, _ := cmd.Flags().GetString("course")
		extract(course)
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	extractCmd.Flags().String("course", "c", "The course to extract definitions from, e.g. ib/antop")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// optimiseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// optimiseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func extract(course string) {
	files, _ := walkMatch(course, "*.tex")
	result := ""
	for _, v := range files {
		contentsBytes, _ := os.ReadFile(v)
		contents := string(contentsBytes)
		for {
			i := strings.Index(contents, `\begin{definition}`)
			if i == -1 {
				break
			}
			j := strings.Index(contents, `\end{definition}`)
			def := contents[i+len(`\begin{definition}`) : j]
			contents = contents[j+len(`\end{definition}`):]

			def = strings.ReplaceAll(def, "\t", "")
			def = strings.ReplaceAll(def, "\n", " ")
			def = strings.TrimSpace(def)

			i1 := strings.Index(def, `\textit{`)
			guessedDefName := "[unknown definition]"
			if i1 != -1 {
				j1 := strings.Index(def[i1:], `}`)
				guessedDefName = def[i1+len(`\textit{`) : i1+j1]
			}

			for {
				// Remove all textit
				i2 := strings.Index(def, `\textit{`)
				if i2 == -1 {
					break
				}
				j2 := strings.Index(def[i2:], `}`)
				def = def[:i2] + "<i>" + def[i2+len(`\textit{`):i2+j2] + "</i>" + def[i2+j2+1:]
			}

			result += guessedDefName + "\t" + def + "\n"
		}
	}
	os.WriteFile("definitions.txt", []byte(result), 0644)
	color.HiGreen("Done!")
}
