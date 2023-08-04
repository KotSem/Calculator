build:
	go build -o calculator main.go readLine.go errorsCode.go calculator.go romanDictionary.go
	./calculator

all: build