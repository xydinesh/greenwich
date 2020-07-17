load("@io_bazel_rules_go//go:def.bzl", "go_binary")
go_binary(
    name = "greenwich",
    srcs = ["main.go", "handlers.go", "points.go", "routers.go"],
    importpath = "greenwich/main",
    visibility = ["//visibility:public"],
    deps = [
          "@com_github_gorilla_mux//:go_default_library",
          "@com_github_bradfitz_latlong//:go_default_library",
    ]
)
