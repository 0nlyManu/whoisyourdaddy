package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/OnlyManuel/whoisyourdaddy/internal/correlator"
	"github.com/OnlyManuel/whoisyourdaddy/internal/models"
	"github.com/OnlyManuel/whoisyourdaddy/internal/ui"
	"github.com/OnlyManuel/whoisyourdaddy/sources"
)

func runSource(src sources.Source, ctx context.Context, target string) models.Result {
	done := make(chan bool)
	go ui.Spinner(src.Name(), done)
	result := src.Run(ctx, target)
	done <- true
	return result
}

func main() {
	flag.Usage = ui.PrintHelp
	target := flag.String("target", "", "target domain")
	wordlist := flag.String("wordlist", "", "external wordlists to use for dns enumeration")
	flag.Parse()

	ui.PrintBanner()

	if *target == "" {
		ui.PrintHelp()
		os.Exit(1)
	}

	ctx := context.Background()
	srcs := []sources.Source{sources.CrtSh{}, sources.DNSEnum{Wordlist: *wordlist}}

	var results []models.Result

	for _, src := range srcs {
		result := runSource(src, ctx, *target)
		results = append(results, result)

		for _, err := range result.Errors {
			fmt.Fprintf(os.Stderr, "%s[-] %s error: %v%s\n", ui.Red, src.Name(), err, ui.Reset)
		}

		for _, asset := range result.Assets {
			fmt.Printf("  %s[subdomain]%s %s\n", ui.Cyan, ui.Reset, asset.Value)
		}
		fmt.Println()
	}

	c := correlator.Correlator{}
	assets := c.Merge(results)

	sort.Slice(assets, func(i, j int) bool {
		return assets[i].RiskScore > assets[j].RiskScore
	})

	fmt.Printf("\n%s━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━%s\n", ui.Cyan, ui.Reset)
	fmt.Printf("%s  RESULTS — %d unique assets%s\n", ui.Bold, len(assets), ui.Reset)
	fmt.Printf("%s━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━%s\n\n", ui.Cyan, ui.Reset)

	for _, asset := range assets {
		var color string
		switch {
		case asset.RiskScore >= 7:
			color = ui.Red
		case asset.RiskScore >= 4:
			color = ui.Yellow
		default:
			color = ui.Green
		}

		ip := asset.Metadata["ip"]
		if ip == "" {
			ip = "n/a"
		}

		fmt.Printf("%s[%2d]%s %-35s ip: %-15s sources: %s\n",
			color, asset.RiskScore, ui.Reset,
			asset.Value,
			ip,
			asset.Metadata["sources"],
		)
	}
	fmt.Println()
}
