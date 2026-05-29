package sources

import (
	"context"

	"github.com/OnlyManuel/whoisyourdaddy/internal/models"
)

type Source interface {
	Name() string
	Run(ctx context.Context, target string) models.Result
}
