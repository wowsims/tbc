load("@npm//@bazel/typescript:index.bzl", "ts_project")
load("@io_bazel_rules_sass//:defs.bzl", "sass_library")

package(default_visibility=["//ui:__subpackages__"])

ts_project(
	name = "default",
	srcs = [
    "default.ts",
    "theme.ts",
  ],
	deps = [
    "//ui/core/components",
  ],
	tsconfig = "//ui/core:tsconfig",
  composite = True,
)

sass_library(
  name = "default_scss",
  srcs = ["_default.scss"],
  deps = [
    "//ui/core/components:character_stats_scss",
  ],
)
