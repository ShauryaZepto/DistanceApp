run:
	@echo Running DistanceApp . . . . . . 
	go run cmd/app/main.go

build:
	@echo Building DistanceApp . . . . . .
	go build cmd/app/main.go
	@echo Executing the build
	./main

test:
	@echo Running all the tests
	go test ./...
	