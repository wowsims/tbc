load("@io_bazel_rules_sass//:defs.bzl", "sass_binary")
load("@npm//http-server:index.bzl", "http_server")
load("@npm//@bazel/typescript:index.bzl", "ts_config", "ts_project")

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
    "//ui/core/themes:default",
  ],
	tsconfig = ":tsconfig",
)

sass_binary(
  name = "elemental_shaman_css",
  src = "elemental_shaman.scss",
  deps = [
    "//ui/core/themes:default_scss",
  ],
)

http_server(
  name = "server",
  data = [
    "//ui/core/components",
    "//ui/core/themes:default",
    ":elemental_shaman",
    ":elemental_shaman_css",
    "index.html",
  ],
)
