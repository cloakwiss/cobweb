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

	// Copied from monolith => Maybe reorder it a bit
	rootCmd.Flags().BoolP("no-audio", "a", false, "Remove audio sources")
	rootCmd.Flags().BoolP("no-css", "c", false, "Remove CSS")
	rootCmd.Flags().BoolP("no-frames", "f", false, "Remove frames and iframes")
	rootCmd.Flags().BoolP("no-fonts", "F", false, "Remove fonts")
	rootCmd.Flags().BoolP("no-js", "j", false, "Remove JavaScript")
	rootCmd.Flags().BoolP("no-images", "i", false, "Remove images")
	rootCmd.Flags().BoolP("no-video", "v", false, "Remove video sources")
	rootCmd.Flags().BoolP("no-metadata", "M", false, "Exclude timestamp and source information")
	rootCmd.Flags().BoolP("ignore-errors", "e", false, "Ignore network errors")
	rootCmd.Flags().BoolP("isolate", "I", false, "Cut off document from the Internet")
	rootCmd.Flags().BoolP("insecure", "k", false, "Allow invalid X.509 (TLS) certificates")
	rootCmd.Flags().BoolP("version", "V", false, "Print version information")
	rootCmd.Flags().BoolP("unwrap-noscript", "n", false, "Replace NOSCRIPT elements with their contents")
	rootCmd.Flags().BoolP("silent", "s", false, "Suppress verbosity")

	// Remaing flags from monolith
	// rootCmd.Flags().BoolP("base-url" <http://localhost/> "b",   "Set custom base URL")
	// rootCmd.Flags().BoolP("blacklist-domains"            "B",   "Treat list of specified domains as blacklist")
	// rootCmd.Flags().BoolP("cookies" <cookies.txt>        "C",   "Specify cookie file")
	// rootCmd.Flags().BoolP("domain" <example.com>         "d",   "Specify domains to use for white/black-listing")
	// rootCmd.Flags().BoolP("encoding" <UTF-8>             "E",   "Enforce custom charset")
	// rootCmd.Flags().BoolP("output" <document.html>       "o",   "Write output to <file>, use - for STDOUT")
	// rootCmd.Flags().BoolP("timeout" <60>                 "t",   "Adjust network request timeout")
	// rootCmd.Flags().BoolP("user-agent" <Firefox>         "u",   "Set custom User-Agent string")

	rootCmd.Execute()

	// Filling the remaining arguments
	//TODO: validate output file name

	fmt.Printf("Args: %+v\n", args)
}
