/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	library "github.com/themooer1/audiobook-library"
	scanner "github.com/themooer1/audiobook-scanner"
	"github.com/themooer1/podgen/podcast"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Prints a new RSS feed to stdout",
	Long:  `Scans a directory and prints RSS feeds for each audiobook it contains.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		audioRoot := args[0]

		lib := library.AudioBookLibrary{}
		lib.Initialize()

		sorter := scanner.SortByDiscNumber[scanner.RelativeAudioBookChapter]

		errors := scanner.Scan(audioRoot, &lib, sorter)

		lib.ForEach(

			func(name string, book *library.AudioBook) error {
				err := podcast.Create(book, "", os.Stdout)
				fmt.Println("")

				return err
			},
		)

		for e := range errors {
			fmt.Println(e)
		}

	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// printCmd.Args(printCmd, ["audio-root", "./", "Folder containing audio files"])
}
