.PHONY: server client www test

all: server client www
clean: server_clean www_clean

server:
	go build -v -o daijoubu ./server

client:
	tinygo build -target wasm -o www/wasm/main.wasm ./client

www:
	git submodule update --init --recursive
	mkdir -p www/{css,media/{source,thumb,ui/{flags,videos}},wasm}
	ln -sf $(PWD)/external/flags/svg/*.svg www/media/ui/flags
	brotli -f www/{css/*.css,lang/*/*.json,media/ui/{favicons/*.ico,flags/*.svg},wasm/*.wasm}

test:
	go test --race ./...

test_no_race:
	go test ./...

server_clean:
	rm -f daijoubu daijoubu.exe

www_clean:
	rm -rf www/{css/*.css*,lang/*/*.br,media/ui/{favicons/*.br,flags/*.br},wasm/*.wasm*}
