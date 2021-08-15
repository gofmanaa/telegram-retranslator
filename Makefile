run:
	go run cmd/telegram-bot/main.go

build:
	go build -race -v cmd/telegram-bot/main.go

test:
	go test -v -race -timeout 30s ./...

env:
	echo GOPATH=$(GOPATH)



up:
	 docker-compose up -d & go build -v cmd/telegram-bot/main.go & ./main &
