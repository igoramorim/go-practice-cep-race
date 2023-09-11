package racer

import (
	"context"
	"fmt"
	"github.com/igoramorim/go-practice-cep-race/internal/cep"
	"github.com/igoramorim/go-practice-cep-race/internal/client"
)

func New(cepGetter client.CEPGetter) *Racer {
	return &Racer{
		CEPGetter: cepGetter,
	}
}

type Racer struct {
	client.CEPGetter
}

func (r *Racer) Run(ctx context.Context, cep cep.CEP, resCh chan<- client.Response, errCh chan<- error) {
	go func() {
		fmt.Printf("Racer %s started!\n", r.Name())
		defer func() {
			fmt.Printf("Racer %s finished!\n", r.Name())
		}()
		res, err := r.Get(ctx, cep)
		if err != nil {
			errCh <- err
			return
		}
		resCh <- res
	}()
}
