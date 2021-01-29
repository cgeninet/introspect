init:
	mkdir -p coverage
	$(rm coverage/cover.*)

coverage: init
	go test -coverprofile coverage/cover.out || true
	go tool cover -html=coverage/cover.out -o coverage/cover.html

all: coverage
