package correlator

import (
	"strings"

	"github.com/OnlyManuel/whoisyourdaddy/internal/models"
)

type Correlator struct{}

func (c Correlator) Merge(results []models.Result) []models.Asset {
	seen := map[string]models.Asset{}

	for _, res := range results {
		for _, asset := range res.Assets {
			if existing, ok := seen[asset.Value]; ok {
				existing.Metadata["sources"] += "," + asset.Source
				seen[asset.Value] = existing
			} else {
				if asset.Metadata == nil {
					asset.Metadata = map[string]string{}
				}
				asset.Metadata["sources"] = asset.Source
				seen[asset.Value] = asset
			}
		}
	}

	var assets []models.Asset
	for _, asset := range seen {
		asset.RiskScore = score(asset)
		assets = append(assets, asset)
	}

	return assets
}

func score(asset models.Asset) int {
	s := 0

	interesting := []string{"admin", "vpn", "gitlab", "internal", "dev", "staging", "api"}
	for _, word := range interesting {
		if strings.Contains(asset.Value, word) {
			s += 3
			break
		}
	}

	if strings.Contains(asset.Metadata["sources"], ",") {
		s += 2
	}

	if asset.Metadata["ip"] != "" {
		s += 1
	}

	if s > 10 {
		s = 10
	}

	return s
}
