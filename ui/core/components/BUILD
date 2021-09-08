load("@npm//@bazel/typescript:index.bzl", "ts_project")
load("@io_bazel_rules_sass//:defs.bzl", "sass_library")

package(default_visibility=["//ui:__subpackages__"])

ts_project(
	name = "components",
	srcs = glob(["*.ts"]),
	deps = [],
	tsconfig = "//ui/core:tsconfig",
  composite = True,
)

sass_library(
  name = "character_stats_scss",
  srcs = ["_character_stats.scss"],
)
