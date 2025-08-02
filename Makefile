run:
	go run cmd/cli-pomodoro/main.go

build:
	go build -o bin/cli-pomodoro cmd/cli-pomodoro/main.go

.PHONY: run build