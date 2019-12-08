.PHONY: server client www test

all: server client www
clean: server_clean www_clean

server:
	go build -v -o daijoubu ./server

client:
	GOOS=js GOARCH=wasm go build -v -o www/wasm/main.wasm ./client

www:
	git submodule update --init --recursive
	mkdir -p www/{css,js,media/{source,thumb,ui/{flags,videos}},wasm}
	npm up --quiet
	npm run --silent gulp -- -LL
	cp -u --no-preserve mode $(shell go env GOROOT)/misc/wasm/wasm_exec.js www/js
	ln -sf $(PWD)/external/flags/svg/*.svg www/media/ui/flags
	brotli -f www/{css/*.css,js/*.js,lang/*/*.json,media/ui/{favicons/*.ico,flags/*.svg},wasm/*.wasm}

# Doesn't work well with https://github.com/gascore/dom
tinygo:
	tinygo build -target wasm -o www/wasm/main.wasm ./client
	cp -u --no-preserve mode $(shell tinygo env TINYGOROOT)/targets/wasm_exec.js www/js

test:
	go test --race ./...

test_no_race:
	go test ./...

server_clean:
	rm -f daijoubu daijoubu.exe

www_clean:
	rm -rf www/{css/*.css*,js/*.js*,lang/*/*.br,media/ui/{favicons/*.br,flags/*.br},wasm/*.wasm*} node_modules
