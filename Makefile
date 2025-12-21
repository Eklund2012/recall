build:
	go build -o recall cmd/root.go

run:
	go run cmd/root.go

test:
	go test ./... -v

clean:
	rm -f recall
