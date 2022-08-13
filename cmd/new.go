/*
Copyright © 2022 David Hay davidhaydev@protonmail.com
*/
package cmd

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

//go:embed templates/*
var templates embed.FS

// handleExisting a .gitignore already exists delete it
func handleExisting() {
	_, err := os.Stat(".gitignore")

	if os.IsNotExist(err) {
		return
	}

	var ans string
	fmt.Println(".gitignore detected. Would you like to create a fresh file? [y/n]")
	fmt.Scanln(&ans)

	if ans == "y" {
		fmt.Println("Deleting...")
		os.Remove(".gitignore")
	} else {
		fmt.Println("Appending...")
	}
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new .gitignore",
	Long: `
create new .gitignore followed by the types of languages, os, etc. that are relevant.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please enter the types of files your project will contain")
			return
		}

		handleExisting() // check if user has file

		ignoreFile, err := os.OpenFile(".gitignore",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		defer func() {
			if err := ignoreFile.Close(); err != nil {
				panic(err)
			} else {
				fmt.Println("All done, go build something awesome!")
			}
		}()

		for _, v := range args {
			data, err := templates.ReadFile("templates/" + strings.ToLower(v) + ".txt")

			if err == nil {
				if _, err = ignoreFile.Write(data); err != nil {
					panic(err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
