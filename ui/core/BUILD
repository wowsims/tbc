load("@npm//@bazel/typescript:index.bzl", "ts_config", "ts_project")

package(default_visibility=["//ui:__subpackages__"])

ts_config(
	name = "tsconfig",
	src = "tsconfig.json",
  deps = [
    "//ui:tsconfig",
  ],
)

ts_project(
	name = "core_ts",
	srcs = glob(["**/*.ts"]),
	deps = [],
	tsconfig = ":tsconfig",
  composite = True,
)

filegroup(
  name = "core_js",
  srcs = [
    "gear_picker.js",
  ],
)
