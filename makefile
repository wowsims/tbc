# Recursive wildcard function
rwildcard=$(foreach d,$(wildcard $(1:=/*)),$(call rwildcard,$d,$2) $(filter $(subst *,%,$2),$d))

# Make everything. Keep this first so it's the default rule.
dist: elemental_shaman

elemental_shaman: dist/elemental_shaman/index.js dist/elemental_shaman/index.css dist/elemental_shaman/index.html

clean:
	rm -f ui/core/api/api.ts
	rm -f ui/core/api/common.ts
	rm -f ui/core/api/druid.ts
	rm -f ui/core/api/shaman.ts
	rm -rf dist

# Host a local server, for dev testing
host: dist
	npx http-server dist

ui/core/api/api.ts: api/api.proto
	npx r.js -convert node_modules/@protobuf-ts/runtime/build/commonjs/ dist/@protobuf-ts
	mkdir -p ui/core/api
	npx protoc --ts_opt generate_dependencies --ts_out ui/core/api --proto_path api api/api.proto

dist/core/tsconfig.tsbuildinfo: $(call rwildcard,ui/core,*.ts) ui/core/api/api.ts
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

generate_items/api/api.go:
	protoc -I api/ --go_out=generate_items/ api/*.proto
