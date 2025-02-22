package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

type arguments struct {
	inputUrl, outputEpub string
	depth                uint
}

func main() {
	var args arguments
	// This should accept more than 1 urls
	rootCmd := &cobra.Command{
		Use:           "cobweb [URL]",
		Short:         "Download Webpages as ePUB",
		Long:          "Mimimal, friendly tool to download webpage as ePUB",
		SilenceErrors: false,
		SilenceUsage:  false,
		Args:          cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, arg []string) error {
			if len(arg) == 1 {
				args.inputUrl = arg[0]
				return nil
			} else {
				return cmd.Help()
			}
		},
	}
	rootCmd.Flags().StringVarP(&args.outputEpub, "output", "o", "", "Name of output epub file")
	rootCmd.Flags().UintVarP(&args.depth, "depth", "d", 1, "Maximum recursion")
	rootCmd.Execute()

	// Filling the remaining arguments
	//TODO: validate output file name

	fmt.Printf("Args: %+v\n", args)
}
