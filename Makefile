APP=capital-gains

build:
	go build -o $(APP) ./cmd/cli

run:
	go run ./cmd/cli

test:
	go test ./... -cover

lint:
	golangci-lint run || true

clean:
	rm -f $(APP)

e2e:
	echo '[{"operation":"buy","unit-cost":10,"quantity":100}]' | go run ./cmd/cli
