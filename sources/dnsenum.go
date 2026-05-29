package sources

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/OnlyManuel/whoisyourdaddy/internal/models"
)

var defaultWordlist = []string{"www", "mail", "admin", "dev", "staging", "gitlab", "vpn", "api", "test", "internal"}

type DNSEnum struct {
	Workers  int
	Wordlist string
}

func loadWordlist(path string) []string {
	if path == "" {
		return defaultWordlist
	}
	file, err := os.Open(path)

	if err != nil {
		fmt.Printf("Error while opening external wordlist\n")
		fmt.Printf("Loading default wordlist...\n")
		return defaultWordlist
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	wordlist := []string{}
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}
	return wordlist
}

func (d DNSEnum) Name() string {
	return "dns-enum"
}

func (d DNSEnum) Run(ctx context.Context, target string) models.Result {
	result := models.Result{Source: d.Name()}

	if d.Workers == 0 {
		d.Workers = 50
	}

	jobs := make(chan string)
	results := make(chan models.Asset)
	var wg sync.WaitGroup
	seen := map[string]bool{}
	for i := 0; i < d.Workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				addrs, err := net.DefaultResolver.LookupHost(ctx, job)
				if err != nil {
					continue
				}

				for _, addr := range addrs {
					results <- models.Asset{
						Type:   "subdomain",
						Value:  job,
						Source: d.Name(),
						Metadata: map[string]string{
							"ip": addr,
						},
					}
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	wordlist := loadWordlist(d.Wordlist)
	for _, word := range wordlist {
		jobs <- fmt.Sprintf("%s.%s", word, target)
	}
	close(jobs)

	for asset := range results {
		if _, ok := seen[asset.Value]; ok {
			continue
		}
		seen[asset.Value] = true
		result.Assets = append(result.Assets, asset)
	}

	return result
}
