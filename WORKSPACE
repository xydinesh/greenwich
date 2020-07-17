workspace(name = "greenwich")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/v0.19.5/rules_go-v0.19.5.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.19.5/rules_go-v0.19.5.tar.gz",
    ],
    sha256 = "513c12397db1bc9aa46dd62f02dd94b49a9b5d17444d49b5a04c5a89f3053c1c",
)

http_archive(
    name = "bazel_gazelle",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/0.18.2/bazel-gazelle-0.18.2.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.18.2/bazel-gazelle-0.18.2.tar.gz",
    ],
    sha256 = "7fc87f4170011201b1690326e8c16c5d802836e3a0d617d8f75c3af2b23180c4",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_bradfitz_latlong",
    importpath = "github.com/bradfitz/latlong",
    sum = "h1:wsnz4B2CSHJ09pwtMReU/GRqWDsI7XSasq7Nphem3Xk=",
    version = "v0.0.0-20170410180902-f3db6d0dff40",
)

go_repository(
    name = "com_github_golang_freetype",
    importpath = "github.com/golang/freetype",
    sum = "h1:DACJavvAHhabrF08vX0COfcOBJRhZ8lUbR+ZWIs0Y5g=",
    version = "v0.0.0-20170609003504-e2365dfdc4a0",
)

go_repository(
    name = "com_github_gorilla_context",
    importpath = "github.com/gorilla/context",
    sum = "h1:9oNbS1z4rVpbnkHBdPZU4jo9bSmrLpII768arSyMFgk=",
    version = "v0.0.0-20160226214623-1ea25387ff6f",
)

go_repository(
    name = "com_github_gorilla_mux",
    importpath = "github.com/gorilla/mux",
    sum = "h1:KOwqsTYZdeuMacU7CxjMNYEKeBvLbxW+psodrbcEa3A=",
    version = "v1.6.1",
)

go_repository(
    name = "com_github_jonas_p_go_shp",
    importpath = "github.com/jonas-p/go-shp",
    sum = "h1:LY81nN67DBCz6VNFn2kS64CjmnDo9IP8rmSkTvhO9jE=",
    version = "v0.1.1",
)

go_repository(
    name = "org_golang_x_image",
    importpath = "golang.org/x/image",
    sum = "h1:1e2ufUJNM3lCHEY5jIgac/7UTjd6cgJNdatjPdFWf34=",
    version = "v0.0.0-20200618115811-c13761719519",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:g61tztE5qeGQ89tm6NTjjM9VPIm088od1l6aSorWRWg=",
    version = "v0.3.0",
)
