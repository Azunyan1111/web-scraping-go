setup:
	go get ./...

start:
	go run main/localserver.go &
	go run main.go
	killall localserver