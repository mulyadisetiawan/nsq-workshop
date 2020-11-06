build:
	@go build -o nsq main.go

run:
	make build
	./nsq