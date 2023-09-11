package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/igoramorim/go-practice-cep-race/internal/cep"
	"github.com/igoramorim/go-practice-cep-race/internal/client"
	"github.com/igoramorim/go-practice-cep-race/internal/client/apicep"
	"github.com/igoramorim/go-practice-cep-race/internal/client/viacep"
	"github.com/igoramorim/go-practice-cep-race/internal/racer"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Cleaning the track...")
	cepIn := flag.String("cep", "", "cep to search")
	timeout := flag.Duration("timeout", time.Second, "max timeout for racers")
	flag.Parse()

	cep, err := cep.New(*cepIn)
	if err != nil {
		fmt.Printf("invalid input: %s\n", err.Error())
		os.Exit(1)
	}

	if err = run(cep, *timeout); err != nil {
		fmt.Printf("%s\n", err.Error())
		os.Exit(1)
	}
}

func run(cep cep.CEP, timeout time.Duration) error {
	ctx := context.Background()
	resCh := make(chan client.Response)
	errCh := make(chan error)

	fmt.Println("Positioning runners...")
	racers := loadRacers()
	for _, racer := range racers {
		racer.Run(ctx, cep, resCh, errCh)
	}

	select {
	case res := <-resCh:
		fmt.Println(fmtOutput(res))
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-time.After(timeout):
		return errors.New("timeout! the racers are too slow :(")
	}

	return nil
}

func loadRacers() []*racer.Racer {
	return []*racer.Racer{
		racer.New(viacep.New()),
		racer.New(apicep.New()),
	}
}

func fmtOutput(res client.Response) string {
	var sb strings.Builder
	sb.WriteString("========== WINNER ==========\n")
	sb.WriteString(fmt.Sprintf("API: %s\n", res.Resource))
	sb.WriteString(fmt.Sprintf("Response: %s\n", res.Data))
	sb.WriteString(fmt.Sprintf("Request time: %s\n", res.Duration))
	sb.WriteString("============================")
	return sb.String()
}
