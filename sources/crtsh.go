package sources

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/OnlyManuel/whoisyourdaddy/internal/models"
)

type CrtSh struct {
}

func (c CrtSh) Name() string {
	return "crt.sh"
}

func (c CrtSh) Run(ctx context.Context, target string) models.Result {
	result := models.Result{Source: c.Name()}
	type crtshEntry struct {
		NameValue string `json:"name_value"`
	}

	url := fmt.Sprintf("https://crt.sh/?q=%%.%s&output=json", target)

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		result.Errors = append(result.Errors, err)
		return result
	}
	req.Header.Set("User-Agent", "whoisyourdaddy/1.0")
	res, err := client.Do(req)
	if err != nil {
		result.Errors = append(result.Errors, err)
		return result
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		result.Errors = append(result.Errors, fmt.Errorf("crtsh returned %d", res.StatusCode))
		return result
	}

	var entries []crtshEntry
	if err := json.NewDecoder(res.Body).Decode(&entries); err != nil {
		result.Errors = append(result.Errors, err)
		return result
	}
	seen := map[string]bool{}
	for _, v := range entries {
		str := strings.Split(v.NameValue, "\n")
		for _, v1 := range str {
			if strings.HasPrefix(v1, "*") {
				continue
			}
			if _, ok := seen[v1]; ok {
				continue
			}
			seen[v1] = true
			asset := models.Asset{
				Type:     "subdomain",
				Value:    v1,
				Metadata: map[string]string{},
				Source:   c.Name(),
			}
			result.Assets = append(result.Assets, asset)
		}
	}

	return result
}
