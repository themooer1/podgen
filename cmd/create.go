/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	library "github.com/themooer1/audiobook-library"
	scanner "github.com/themooer1/audiobook-scanner"
	"github.com/themooer1/podgen/podcast"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates an RSS feed directory",
	Long: `Scans a directory and prints creates RSS feeds for each audiobook it contains
	./rss`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		audioRoot := args[0]

		lib := library.AudioBookLibrary{}
		lib.Initialize()

		sorter := scanner.SortByDiscNumber[scanner.RelativeAudioBookChapter]

		errors := scanner.Scan(audioRoot, &lib, sorter)

		// Create RSS shadow directory with RSS feeds and assets e.g. thumbnail photos.
		// Should be restricted by umask
		err := os.MkdirAll("rss", 0777)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create RSS directory: ", err)
			os.Exit(-1)
		}

		lib.ForEach(

			func(name string, book *library.AudioBook) error {

				rssDirectory := filepath.Join("./rss", book.Author, book.Title)
				rssFilePath := filepath.Join(rssDirectory, "rss.xml")

				err := os.MkdirAll(rssDirectory, 0755)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Failed to create directory: ", rssDirectory, "\n", err)
				}

				rssFile, err := os.OpenFile(rssFilePath, os.O_WRONLY|os.O_CREATE, 0644)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Failed to create podcast at ", rssFilePath, "\n", err)
				} else {
					defer rssFile.Close()
				}

				err = podcast.Create(book, "../../../", rssFile)
				if err != nil {
					fmt.Println(err)
				}

				return err
			},
		)

		for e := range errors {
			fmt.Println(e)
		}

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// createCmd.Args(createCmd, ["audio-root", "./", "Folder containing audio files"])
}
