BIN_NAME="run_adf_x_dr_d_dr_voice_online"

.PHONY: build
build:
	go build -o build/$(BIN_NAME) .


.PHONY: install
install:
	go install

.PHONY: test
test:
	go test -cover -v -race
