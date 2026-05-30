package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/ossf/pvtr-github-repo-scanner/badgeurl"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	var (
		filePath              string
		badge                 string
		includeJustifications bool
	)

	flags := flag.NewFlagSet("badge-url", flag.ContinueOnError)
	flags.SetOutput(stderr)
	flags.StringVar(&filePath, "f", "", "path to Privateer YAML results file")
	flags.StringVar(&badge, "badge", badgeurl.DefaultBadge, "badge target: choose, baseline-1, baseline-2, or baseline-3")
	flags.BoolVar(&includeJustifications, "justifications", true, "include justification text in generated Best Practices Badge links")
	if err := flags.Parse(args); err != nil {
		return 2
	}

	if filePath == "" {
		_, _ = fmt.Fprintln(stderr, "badge-url requires -f <results.yaml>")
		return 2
	}

	urls, err := badgeurl.GenerateFromFile(filePath, badgeurl.Options{
		Badge:                 badge,
		IncludeJustifications: &includeJustifications,
	})
	if err != nil {
		_, _ = fmt.Fprintln(stderr, err)
		return 1
	}

	if len(urls) == 1 {
		_, _ = fmt.Fprintln(stderr, "Open this link in your browser, review the proposed answers, and save.")
	} else {
		_, _ = fmt.Fprintln(stderr, "Open the links in order. After each link, review the proposed answers and save before opening the next one.")
	}

	for _, generatedURL := range urls {
		if _, err := fmt.Fprintln(stdout, generatedURL); err != nil {
			_, _ = fmt.Fprintln(stderr, err)
			return 1
		}
	}

	return 0
}
