# Make everything
dist: elemental_shaman

clean_dist:
	rm -r dist

# Host a local server, for dev testing
host:
	npx http-server dist

# tsc -b already knows how to build dependencies, so this rule is only for
# building ui/core in isolation for development
core_js:
	npx tsc -p ui/core

elemental_shaman_js:
	npx tsc -b ui/elemental_shaman

elemental_shaman_css:
	npx sass ui/elemental_shaman/index.scss dist/elemental_shaman/index.css

elemental_shaman_html:
	cp ui/elemental_shaman/index.html dist/elemental_shaman

elemental_shaman: elemental_shaman_js elemental_shaman_css elemental_shaman_html
