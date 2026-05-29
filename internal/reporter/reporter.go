package reporter

import (
	"os"

	"github.com/OnlyManuel/whoisyourdaddy/internal/models"
)

type ReportData struct {
	Target      string
	Date        string
	TotalAssets int
	HighRisk    int
	MediumRisk  int
	LowRisk     int
	Assets      []models.Asset
}

type Reporter struct {
	OutputPath string
}

func (r Reporter) Generate(reportData ReportData) error {
	os.Create(r.OutputPath)
	return nil
}
