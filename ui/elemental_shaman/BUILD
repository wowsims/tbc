load("@npm//@bazel/typescript:index.bzl", "ts_config", "ts_project")
load("@npm//http-server:index.bzl", "http_server")

ts_config(
	name = "tsconfig",
	src = "tsconfig.json",
  deps = [
    "//ui:tsconfig",
  ],
)

ts_project(
	name = "elemental_shaman",
	srcs = [
    "elemental_shaman.ts",
  ],
	deps = [
    "//ui/core:core_ts",
  ],
	tsconfig = ":tsconfig",
)

http_server(
  name = "server",
  data = [
    "//ui/core:core_js",
    ":elemental_shaman",
    "index.html",
  ],
)
