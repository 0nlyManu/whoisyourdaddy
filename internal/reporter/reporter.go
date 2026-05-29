package reporter

import (
	_ "embed"
	"html/template"
	"os"

	"github.com/OnlyManuel/whoisyourdaddy/internal/models"
)

//go:embed templates/report.html
var reportTemplate string

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
	file, err := os.Create(r.OutputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	tmpl, err := template.New("report").Parse(reportTemplate)
	if err != nil {
		return err
	}
	return tmpl.Execute(file, reportData)
}
