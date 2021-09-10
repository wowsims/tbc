# Make everything
dist: elemental_shaman

clean_dist:
	rm -r dist

# Host a local server, for dev testing
host:
	npx http-server dist

proto_ts: api/newapi.proto
	rm -r ui/core/api
	mkdir -p ui/core/api
	npx protoc --ts_out ui/core/api --proto_path api api/newapi.proto

# tsc -b already knows how to build dependencies, so this rule is only for
# building ui/core in isolation for development
core_js: proto_ts
	npx tsc -p ui/core

elemental_shaman_js:
	npx tsc -b ui/elemental_shaman

elemental_shaman_css:
	npx sass ui/elemental_shaman/index.scss dist/elemental_shaman/index.css

elemental_shaman_html:
	cp ui/elemental_shaman/index.html dist/elemental_shaman

elemental_shaman: elemental_shaman_js elemental_shaman_css elemental_shaman_html
