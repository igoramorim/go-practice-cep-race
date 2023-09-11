package apicep

import (
	"context"
	"fmt"
	"github.com/igoramorim/go-practice-cep-race/internal/cep"
	"github.com/igoramorim/go-practice-cep-race/internal/client"
	"github.com/igoramorim/go-practice-cep-race/internal/http"
	"github.com/pkg/errors"
)

func New() *APICEP {
	return &APICEP{}
}

type APICEP struct{}

const urlFmt = "https://cdn.apicep.com/file/apicep/%s.json"

func (api *APICEP) Get(ctx context.Context, cep cep.CEP) (client.Response, error) {
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

func (api *APICEP) Name() string {
	return "API CEP"
}

func (api *APICEP) formatCEP(cep cep.CEP) string {
	// CEP must return in format xxxxx-xx
	return fmt.Sprintf("%s-%s", cep.Prefix(), cep.Suffix())
}
