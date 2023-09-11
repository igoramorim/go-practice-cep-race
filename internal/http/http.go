package http

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(ctx context.Context, url string) (string, time.Duration, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}

	start := time.Now().UTC()
	httpRes, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer httpRes.Body.Close()
	end := time.Now().UTC().Sub(start)

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return "", 0, err
	}

	if httpRes.StatusCode > 200 {
		return "", 0, errors.New(fmt.Sprintf("api response code: %d %s", httpRes.StatusCode, string(body)))
	}

	return string(body), end, nil
}
