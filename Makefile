# TODO: Remove this export when Go 1.14 lands.
export GO111MODULE=auto

.PHONY: server client test

all: server client flags
clean: server_clean client_clean flags_clean

server:
	go build -v -o daijoubu server/...

templates:
	go generate server/...

client:
	GOOS=js GOARCH=wasm go build -v -o www/wasm/main.wasm client/...
	brotli -jf www/wasm/main.wasm

flags:
	git submodule update --init --recursive
	cp -rf external/flags/svg/*.svg www/media/ui/flags
	brotli -jf www/media/ui/flags/*.svg

test:
	go test --race ./...

test_no_race:
	go test ./...

server_clean:
	rm -f daijoubu daijoubu.exe

client_clean:
	rm -rf www/wasm/*.{wasm,wasm.br} www/css/*.{css,css.br} www/css/maps

flags_clean:
	rm -f www/media/ui/flags/*.{svg,svg.br}
