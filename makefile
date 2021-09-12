# Recursive wildcard function
rwildcard=$(foreach d,$(wildcard $(1:=/*)),$(call rwildcard,$d,$2) $(filter $(subst *,%,$2),$d))

# Make everything. Keep this first so it's the default rule.
dist: elemental_shaman

elemental_shaman: dist/elemental_shaman/index.js dist/elemental_shaman/index.css dist/elemental_shaman/index.html

clean:
	rm -f ui/core/api/newapi.ts
	rm -rf dist

# Host a local server, for dev testing
host: dist
	npx http-server dist

ui/core/api/newapi.ts: api/newapi.proto
	npx r.js -convert node_modules/@protobuf-ts/runtime/build/commonjs/ dist/@protobuf-ts
	mkdir -p ui/core/api
	npx protoc --ts_opt generate_dependencies --ts_out ui/core/api --proto_path api api/newapi.proto

dist/core/tsconfig.tsbuildinfo: $(call rwildcard,ui/core,*.ts) ui/core/api/newapi.ts
	npx tsc -p ui/core

# Generic rule for building index.js for any class directory
dist/%/index.js: ui/%/index.ts dist/core/tsconfig.tsbuildinfo
	npx tsc -p $(<D) 

# Generic rule for building index.css for any class directory
dist/%/index.css: ui/%/index.scss $(call rwildcard,ui/core,*.scss)
	mkdir -p $(@D)
	npx sass $< $@

# Generic rule for building index.html for any class directory
dist/%/index.html: index_template.html
	$(eval title := $(shell echo $(shell basename $(@D)) | sed -r 's/(^|_)([a-z])/\U \2/g' | cut -c 2-))
	echo $(title)
	mkdir -p $(@D)
	cat index_template.html | sed 's/@@TITLE@@/$(title)/g' > $@

wasm:
	GOOS=js GOARCH=wasm go build -o ./dist/lib.wasm ./cmd/simwasm/

web: proto_go
	go build -o simweb ./cmd/simweb/web.go

proto_go:
	protoc -I=./api/ --go_out=. ./api/newapi.proto