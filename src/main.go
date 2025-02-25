package main

import (
	"fmt"
	"time"

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
		Use:           "cobweb",
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

	msg := "Same as the page's title"
	rootCmd.Flags().StringVarP(&args.outputEpub, "output", "O", msg, "Name of output epub file")
	rootCmd.Flags().UintVarP(&args.depth, "depth", "d", 1, "Maximum recursion")

	// Copied from monolith => Maybe reorder it a bit
	// Exclusion flags for various elements
	rootCmd.Flags().BoolP("no-js", "j", false, "Remove JavaScript")
	rootCmd.Flags().BoolP("no-css", "c", false, "Remove CSS")
	rootCmd.Flags().BoolP("no-images", "i", false, "Remove images")
	rootCmd.Flags().BoolP("no-fonts", "f", false, "Remove fonts")
	rootCmd.Flags().BoolP("no-audio", "a", false, "Remove audio sources")
	rootCmd.Flags().BoolP("no-video", "v", false, "Remove video sources")
	rootCmd.Flags().BoolP("no-metadata", "m", false, "Exclude timestamp and source information")
	//TODO: Check if this is possible in EPUB
	// rootCmd.Flags().BoolP("no-frames", "F", false, "Remove frames and iframes")

	//TODO: Will need Need to lookup these flags' behaviour
	// rootCmd.Flags().BoolP("ignore-errors", "e", false, "Ignore network errors")
	// rootCmd.Flags().BoolP("isolate", "I", false, "Cut off document from the Internet")
	// rootCmd.Flags().BoolP("insecure", "k", false, "Allow invalid X.509 (TLS) certificates")
	// rootCmd.Flags().BoolP("unwrap-noscript", "n", false, "Replace NOSCRIPT elements with their contents")

	// Remaing flags from monolith
	// rootCmd.Flags().BoolP("base-url" <http://localhost/> "b",   "Set custom base URL")
	// rootCmd.Flags().BoolP("user-agent",           "u",   "Set custom User-Agent string")
	rootCmd.Flags().StringP("cookies", "C", "", "Specify cookie file")
	rootCmd.Flags().StringArrayP("allow-domain", "A", []string{}, "Specify domains to use for white/black-listing")
	rootCmd.Flags().StringArrayP("block-domains", "D", []string{}, "Treat list of specified domains as blacklist")
	rootCmd.Flags().DurationP("timeout", "T", 60*time.Second, "Adjust network request timeout")

	rootCmd.Flags().BoolP("silent", "s", false, "Suppress verbosity")
	rootCmd.Flags().BoolP("version", "V", false, "Print version information")

	rootCmd.Flags().SortFlags = false

	rootCmd.Execute()

	// Filling the remaining arguments
	//TODO: validate output file name

	fmt.Printf("Args: %+v\n", args)
}
