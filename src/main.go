package main

import (
	"fmt"
	"time"

	"github.com/cloakwiss/cobweb/app"
	"github.com/spf13/cobra"
)

func main() {
	var args app.Options
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
				if err := args.AddUrl(arg[0]); err != nil {
					// this is not the best thing to do
					panic(err)
				}
				return nil
			} else {
				return cmd.Help()
			}
		},
	}

	rootCmd.Flags().StringVarP(&args.Output, "output", "O", "", "Name of output epub file (Same as the page's title)")

	rootCmd.Flags().Uint8VarP(&args.Depth, "depth", "d", 0, "Maximum recursion")

	// Copied from monolith => Maybe reorder it a bit
	// Exclusion flags for various elements
	rootCmd.Flags().BoolVarP(&args.Js, "no-js", "j", false, "Remove JavaScript")
	rootCmd.Flags().BoolVarP(&args.Css, "no-css", "c", false, "Remove CSS")
	rootCmd.Flags().BoolVarP(&args.Images, "no-images", "i", false, "Remove images")
	rootCmd.Flags().BoolVarP(&args.Fonts, "no-fonts", "f", false, "Remove fonts")
	rootCmd.Flags().BoolVarP(&args.Audio, "no-audio", "a", false, "Remove audio sources")
	rootCmd.Flags().BoolVarP(&args.Video, "no-video", "V", false, "Remove video sources")
	rootCmd.Flags().BoolVarP(&args.Metadata, "no-metadata", "m", false, "Exclude timestamp and source information")
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
	var allow, block []string
	allow = *rootCmd.Flags().StringArrayP("allow-domain", "A", []string{}, "Specify domains to use for white/black-listing")
	args.AddAllowDomain(allow)
	block = *rootCmd.Flags().StringArrayP("block-domains", "D", []string{}, "Treat list of specified domains as blacklist")
	args.AddBlockDomain(block)

	rootCmd.Flags().StringVarP(&args.Cookie, "cookies", "C", "", "Specify cookie file")
	rootCmd.Flags().DurationVarP(&args.Timeout, "timeout", "T", 60*time.Second, "Adjust network request timeout")

	rootCmd.Flags().BoolP("silent", "s", false, "Suppress verbosity")
	rootCmd.Flags().BoolP("version", "v", false, "Print version information")

	rootCmd.Flags().SortFlags = false

	rootCmd.Execute()

	// Filling the remaining arguments
	//TODO: validate output file name

	fmt.Printf("Args: %+v\n", args)
}
