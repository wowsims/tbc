# Recursive wildcard function
rwildcard=$(foreach d,$(wildcard $(1:=/*)),$(call rwildcard,$d,$2) $(filter $(subst *,%,$2),$d))

# Make everything. Keep this first so it's the default rule.
dist: elemental_shaman dist/lib.wasm dist/sim_worker.js

elemental_shaman: dist/elemental_shaman/index.js dist/elemental_shaman/index.css dist/elemental_shaman/index.html detailed_results

detailed_results: dist/detailed_results/index.js dist/detailed_results/index.css dist/detailed_results/index.html

clean:
	rm -f ui/core/proto/*.ts
	rm -f sim/core/proto/*.pb.go
	rm -rf dist

# Host a local server, for dev testing
host: dist
	npx http-server dist

ui/core/proto/proto.ts: proto/*.proto
	mkdir -p dist/protobuf-ts
	cp -r node_modules/@protobuf-ts/runtime/build/es2015/* dist/protobuf-ts
	sed -i -E "s/from '(.*)';/from '\1\.js';/g" dist/protobuf-ts/*
	sed -i -E "s/from \"(.*)\";/from '\1\.js';/g" dist/protobuf-ts/*
	# This is needed for local hosting, since github pages serves under the 'tbc' directory.
	mkdir -p dist/tbc/protobuf-ts
	cp -r dist/protobuf-ts dist/tbc
	npx protoc --ts_opt generate_dependencies --ts_out ui/core/proto --proto_path proto proto/api.proto

dist/core/tsconfig.tsbuildinfo: $(call rwildcard,ui/core,*.ts) ui/core/proto/proto.ts
	npx tsc -p ui/core
	sed -i 's/@protobuf-ts\/runtime/\/tbc\/protobuf-ts\/index/g' dist/core/proto/*.js
	sed -i -E "s/from \"(.*?)(\.js)?\";/from '\1\.js';/g" dist/core/proto/*.js

# Generic rule for building index.js for any class directory
dist/%/index.js: ui/%/index.ts ui/%/*.ts dist/core/tsconfig.tsbuildinfo
	npx tsc -p $(<D) 

# Generic rule for building index.css for any class directory
dist/%/index.css: ui/%/index.scss ui/%/*.scss $(call rwildcard,ui/core,*.scss)
	mkdir -p $(@D)
	npx sass $< $@

# Generic rule for building index.html for any class directory
dist/%/index.html: ui/index_template.html
	$(eval title := $(shell echo $(shell basename $(@D)) | sed -r 's/(^|_)([a-z])/\U \2/g' | cut -c 2-))
	echo $(title)
	mkdir -p $(@D)
	cat index_template.html | sed 's/@@TITLE@@/TBC $(title) Simulator/g' > $@

.PHONY: wasm
wasm: dist/lib.wasm

dist/sim_worker.js: ui/worker/sim_worker.js
	cp ui/worker/sim_worker.js dist

# TODO: make different wasm generators per spec
# TODO: how to make this understand 
dist/lib.wasm: sim/wasm/* sim/core/proto/api.pb.go $(filter-out $(wildcard sim/core/items/*), $(call rwildcard,sim,*.go))
	GOOS=js GOARCH=wasm go build --tags=elemental_shaman -o ./dist/lib.wasm ./sim/wasm/

# Just builds the server binary
elesimweb: sim/core/proto/api.pb.go $(filter-out $(wildcard sim/core/items/*), $(call rwildcard,sim,*.go))
	go build --tags=elemental_shaman -o simweb ./sim/web/main.go

# Starts up a webserver hosting the dist/ and API endpoints.
elerunweb: sim/core/proto/api.pb.go
	go run --tags=elemental_shaman ./sim/web/main.go

sim/core/proto/api.pb.go: proto/*.proto
	protoc -I=./proto --go_out=./sim/core ./proto/*.proto

.PHONY: items
items: sim/core/items/all.go

sim/core/items/all.go: generate_items/*.go $(call rwildcard,sim/core/proto,*.go)
	go run generate_items/*.go -outDir=sim/core/items

test: dist/lib.wasm
	go test ./...
