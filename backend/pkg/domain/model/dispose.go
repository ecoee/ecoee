package model

import (
	"context"
)

type Dispose struct {
	ID    string
	Name  string
	Count int
}

type DisposeRepository interface {
	Save(ctx context.Context, dispose Dispose) (Dispose, error)
}
