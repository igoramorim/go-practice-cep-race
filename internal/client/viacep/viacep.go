package viacep

import (
	"context"
	"fmt"
	"github.com/igoramorim/go-practice-cep-race/internal/cep"
	"github.com/igoramorim/go-practice-cep-race/internal/client"
	"github.com/igoramorim/go-practice-cep-race/internal/http"
	"github.com/pkg/errors"
)

func New() *ViaCEP {
	return &ViaCEP{}
}

type ViaCEP struct{}

const urlFmt = "http://viacep.com.br/ws/%s/json/"

func (api *ViaCEP) Get(ctx context.Context, cep cep.CEP) (client.Response, error) {
	url := fmt.Sprintf(urlFmt, api.formatCEP(cep))
	data, duration, err := http.Get(ctx, url)
	if err != nil {
		return client.Response{}, errors.WithMessage(err, url)
	}

	return client.Response{
		Resource: url,
		Data:     data,
		Duration: duration,
	}, nil
}

func (api *ViaCEP) Name() string {
	return "VIA CEP"
}

func (api *ViaCEP) formatCEP(cep cep.CEP) string {
	// CEP must return in format xxxxxxx
	return fmt.Sprintf("%s%s", cep.Prefix(), cep.Suffix())
}
