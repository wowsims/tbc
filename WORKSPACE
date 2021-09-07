workspace(
  name = "tbc",
  # Map the @npm bazel workspace to the node_modules directory.
  # This lets Bazel use the same node_modules as other local tooling.
  managed_directories = {"@npm": ["node_modules"]},
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Loads all the NPM packages dependencies we have
http_archive(
  name = "build_bazel_rules_nodejs",
  url = "https://github.com/bazelbuild/rules_nodejs/releases/download/4.0.0/rules_nodejs-4.0.0.tar.gz",
  sha256 = "8a7c981217239085f78acc9898a1f7ba99af887c1996ceb3b4504655383a2c3c",
)
load("@build_bazel_rules_nodejs//:index.bzl", "npm_install")
npm_install(
  # Name this npm so that Bazel Label references look like @npm//package
  name = "npm",
  package_json = "//:package.json",
  package_lock_json = "//:package-lock.json",
)

# Sass rules for building CSS files
http_archive(
  name = "io_bazel_rules_sass",
  url = "https://github.com/bazelbuild/rules_sass/archive/1.39.0.zip",
  strip_prefix = "rules_sass-1.39.0",
  sha256 = "334b2ad87c13109486a8bfdc8d80d90e4ce6a4528bc6fb090b021ec87c2c3080"
)
load("@io_bazel_rules_sass//:package.bzl", "rules_sass_dependencies")
rules_sass_dependencies()
load("@io_bazel_rules_sass//:defs.bzl", "sass_repositories")
sass_repositories()
