BIN_NAME="run_adf_x_dr_d_dr_voice_online"

.PHONY: build
build:
	env GOOS=darwin GOARCH=amd64 go build -o bin/$(BIN_NAME)_darwin
	env GOOS=linux GOARCH=amd64 go build -o bin/$(BIN_NAME)_linux
	env GOOS=windows GOARCH=amd64 go build -o bin/$(BIN_NAME)_windows


.PHONY: install
install:
	go install

.PHONY: test
test:
	go test -cover -v -race
