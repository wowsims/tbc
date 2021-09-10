# Recursive wildcard function
rwildcard=$(foreach d,$(wildcard $(1:=/*)),$(call rwildcard,$d,$2) $(filter $(subst *,%,$2),$d))

# Make everything. Keep this first so it's the default rule.
dist: elemental_shaman

clean:
	rm -f ui/core/api/newapi.ts
	rm -rf dist

# Host a local server, for dev testing
host: dist
	npx http-server dist

ui/core/api/newapi.ts: api/newapi.proto
	mkdir -p ui/core/api
	npx protoc --ts_out ui/core/api --proto_path api api/newapi.proto

dist/core/core.js: $(call rwildcard,ui/core,*.ts) ui/core/api/newapi.ts
	npx tsc -p ui/core

# Generic rule for building index.js for any class directory
dist/%/index.js: ui/%/index.ts dist/core/core.js
	npx tsc -p $(<D) 

# Generic rule for building index.css for any class directory
dist/%/index.css: ui/%/index.scss
	mkdir -p $(@D)
	npx sass $< $@

# Generic rule for building index.html for any class directory
dist/%/index.html: ui/%/index.html
	mkdir -p $(@D)
	cp $< $@

elemental_shaman: dist/elemental_shaman/index.js dist/elemental_shaman/index.css dist/elemental_shaman/index.html
