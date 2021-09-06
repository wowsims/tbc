load("@npm//@bazel/typescript:index.bzl", "ts_project")
load("@npm//http-server:index.bzl", "http_server")

ts_project(
	name = "elemental_shaman",
	srcs = ["elemental_shaman.ts"],
	args = [
		"--module",
		"commonjs",
		"--moduleResolution",
		"node",
	],
	deps = [],
	tsconfig = "//ui:tsconfig",
)

http_server(
  name = "server",
  data = [
    "index.html",
    "elemental_shaman.js",
  ],
  args = [
    ".",
  ],
)
