# go-practice-cep-routine

### What
It is a race of CEP APIs. The API that returns the response first wins.

There are only two racers initially:

- http://viacep.com.br/ws/{CEP}/json/
- https://cdn.apicep.com/file/apicep/{CEP}.json

### Run

To execute:
```
go run cmd/main.go -cep 04313110 -timeout 300ms
```

Where:

`-cep` is the cep to search for. Required.

`-timeout` is the max timeout for racers. The racer is over if none of them returns before it. Optional. Default is 1s
