# Recursive wildcard function
rwildcard=$(foreach d,$(wildcard $(1:=/*)),$(call rwildcard,$d,$2) $(filter $(subst *,%,$2),$d))

OUT_DIR=dist/tbc

# Make everything. Keep this first so it's the default rule.
$(OUT_DIR): balance_druid elemental_shaman enhancement_shaman shadow_priest raid

# Add new sim rules here! Don't forget to add it as a dependency to the default rule above.
balance_druid: $(OUT_DIR)/balance_druid/index.js $(OUT_DIR)/balance_druid/index.css $(OUT_DIR)/balance_druid/index.html ui_shared
elemental_shaman: $(OUT_DIR)/elemental_shaman/index.js $(OUT_DIR)/elemental_shaman/index.css $(OUT_DIR)/elemental_shaman/index.html ui_shared
enhancement_shaman: $(OUT_DIR)/enhancement_shaman/index.js $(OUT_DIR)/enhancement_shaman/index.css $(OUT_DIR)/enhancement_shaman/index.html ui_shared
shadow_priest: $(OUT_DIR)/shadow_priest/index.js $(OUT_DIR)/shadow_priest/index.css $(OUT_DIR)/shadow_priest/index.html ui_shared

raid: $(OUT_DIR)/raid/index.js $(OUT_DIR)/raid/index.css $(OUT_DIR)/raid/index.html

ui_shared: $(OUT_DIR)/lib.wasm $(OUT_DIR)/sim_worker.js $(OUT_DIR)/net_worker.js detailed_results
detailed_results: $(OUT_DIR)/detailed_results/index.js $(OUT_DIR)/detailed_results/index.css $(OUT_DIR)/detailed_results/index.html

clean:
	rm -f ui/core/proto/*.ts
	rm -f sim/core/proto/*.pb.go
	rm -f wowsimtbc
	rm -f wowsimtbc-windows.exe
	rm -f wowsimtbc-amd64-darwin
	rm -f wowsimtbc-amd64-linux
	rm -rf dist
	rm -rf binary_dist
	find . -name "*.testresults.tmp" -type f -delete

# Host a local server, for dev testing
host: $(OUT_DIR)
	# Intentionally serve one level up, so the local site has 'tbc' as the first
	# directory just like github pages.
	npx http-server $(OUT_DIR)/..

ui/core/proto/proto.ts: proto/*.proto node_modules
	mkdir -p $(OUT_DIR)/protobuf-ts
	cp -r node_modules/@protobuf-ts/runtime/build/es2015/* $(OUT_DIR)/protobuf-ts
	sed -i -E "s/from '(.*)';/from '\1\.js';/g" $(OUT_DIR)/protobuf-ts/*
	sed -i -E "s/from \"(.*)\";/from '\1\.js';/g" $(OUT_DIR)/protobuf-ts/*
	npx protoc --ts_opt generate_dependencies --ts_out ui/core/proto --proto_path proto proto/api.proto
	npx protoc --ts_out ui/core/proto --proto_path proto proto/test.proto
	npx protoc --ts_out ui/core/proto --proto_path proto proto/ui.proto

node_modules: package-lock.json
	npm install

$(OUT_DIR)/core/tsconfig.tsbuildinfo: $(call rwildcard,ui/core,*.ts) ui/core/proto/proto.ts
	npx tsc -p ui/core
	sed -i 's/@protobuf-ts\/runtime/\/tbc\/protobuf-ts\/index/g' $(OUT_DIR)/core/proto/*.js
	sed -i -E "s/from \"(.*?)(\.js)?\";/from '\1\.js';/g" $(OUT_DIR)/core/proto/*.js

# Generic rule for building index.js for any class directory
$(OUT_DIR)/%/index.js: ui/%/index.ts ui/%/*.ts $(OUT_DIR)/core/tsconfig.tsbuildinfo
	npx tsc -p $(<D) 

# Generic rule for building index.css for any class directory
$(OUT_DIR)/%/index.css: ui/%/index.scss ui/%/*.scss $(call rwildcard,ui/core,*.scss)
	mkdir -p $(@D)
	npx sass $< $@

# Generic rule for building index.html for any class directory
$(OUT_DIR)/%/index.html: ui/index_template.html
	$(eval title := $(shell echo $(shell basename $(@D)) | sed -r 's/(^|_)([a-z])/\U \2/g' | cut -c 2-))
	echo $(title)
	mkdir -p $(@D)
	cat ui/index_template.html | sed 's/@@TITLE@@/TBC $(title) Simulator/g' > $@

.PHONY: wasm
wasm: $(OUT_DIR)/lib.wasm

# Builds the generic .wasm, with all items included.
$(OUT_DIR)/lib.wasm: sim/wasm/* sim/core/proto/api.pb.go $(filter-out sim/core/items/all_items.go, $(call rwildcard,sim,*.go))
	@echo "Starting webassembly compile now..."
	@if GOOS=js GOARCH=wasm go build -o ./$(OUT_DIR)/lib.wasm ./sim/wasm/; then \
		echo "\033[1;32mWASM compile successful.\033[0m"; \
	else \
		echo "\033[1;31mWASM COMPILE FAILED\033[0m"; \
		exit 1; \
	fi
	

# Generic sim_worker that uses the generic lib.wasm
$(OUT_DIR)/sim_worker.js: ui/worker/sim_worker.js
	cp ui/worker/sim_worker.js $(OUT_DIR)

$(OUT_DIR)/net_worker.js: ui/worker/net_worker.js
	cp ui/worker/net_worker.js $(OUT_DIR)

binary_dist/dist.go: sim/web/dist.go.tmpl
	mkdir -p binary_dist/tbc
	touch binary_dist/tbc/embedded
	cp sim/web/dist.go.tmpl binary_dist/dist.go

binary_dist: $(OUT_DIR)
	rm -rf binary_dist
	mkdir -p binary_dist
	cp -r $(OUT_DIR) binary_dist/
	rm binary_dist/tbc/lib.wasm

# Builds the web server with the compiled client.
wowsimtbc: sim/web/main.go  binary_dist binary_dist/dist.go devserver

devserver:
	@echo "Starting server compile now..."
	@if go build -o wowsimtbc ./sim/web/main.go; then \
		echo "\033[1;32mBuild Completed Succeessfully\033[0m"; \
	else \
		echo "\033[1;31mBUILD FAILED\033[0m"; \
		exit 1; \
	fi

release: wowsimtbc
	GOOS=windows GOARCH=amd64 go build -o wowsimtbc-windows.exe ./sim/web/main.go
	GOOS=darwin GOARCH=amd64 go build -o wowsimtbc-amd64-darwin ./sim/web/main.go
	GOOS=linux GOARCH=amd64 go build -o wowsimtbc-amd64-linux ./sim/web/main.go

sim/core/proto/api.pb.go: proto/*.proto
	protoc -I=./proto --go_out=./sim/core ./proto/*.proto

.PHONY: items
items: sim/core/items/all_items.go sim/core/proto/api.pb.go

sim/core/items/all_items.go: generate_items/*.go $(call rwildcard,sim/core/proto,*.go)
	go run generate_items/*.go -outDir=sim/core/items
	gofmt -w ./sim/core/items

test: $(OUT_DIR)/lib.wasm binary_dist/dist.go
	go test ./...

update-tests:
	find . -name "*.testresults" -type f -delete
	find . -name "*.testresults.tmp" -exec bash -c 'cp "$$1" "$${1%.testresults.tmp}".testresults' - '{}' +

fmt:
	gofmt -w ./sim

# one time setup to install pre-commit hook for gofmt and npm install needed packages
setup:
	cp pre-commit .git/hooks
	chmod +x .git/hooks/pre-commit
