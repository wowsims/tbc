load("@npm//@bazel/typescript:index.bzl", "ts_config", "ts_project")

package(default_visibility=["//ui:__subpackages__"])

ts_config(
	name = "tsconfig",
	src = "tsconfig.json",
  deps = [
    "//ui:tsconfig",
  ],
)
