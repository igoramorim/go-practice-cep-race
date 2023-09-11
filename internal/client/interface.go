package client

import (
	"context"
	"github.com/igoramorim/go-practice-cep-race/internal/cep"
	"time"
)

type CEPGetter interface {
	Get(ctx context.Context, CEP cep.CEP) (Response, error)
	Name() string
}

type Response struct {
	Resource string
	Data     string
	Duration time.Duration
}
