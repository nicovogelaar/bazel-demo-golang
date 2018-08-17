load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.15.0/rules_go-0.15.0.tar.gz"],
    sha256 = "56d946edecb9879aed8dff411eb7a901f687e242da4fa95c81ca08938dd23bb4",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.14.0/bazel-gazelle-0.14.0.tar.gz"],
    sha256 = "c0a5739d12c6d05b6c1ad56f2200cb0b57c5a70e03ebd2f7b87ce88cabf09c7b",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "6dede2c65ce86289969b907f343a1382d33c14fbce5e30dd17bb59bb55bb6593",
    strip_prefix = "rules_docker-0.4.0",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.4.0.tar.gz"],
)

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

go_repository(
    name = "com_github_golang_protobuf",
    commit = "b4deda0973fb4c70b50d226b1af49f3da59f5265",
    importpath = "github.com/golang/protobuf",
)

go_repository(
    name = "org_golang_google_genproto",
    commit = "c66870c02cf823ceb633bcd05be3c7cda29976f4",
    importpath = "google.golang.org/genproto",
)

go_repository(
    name = "org_golang_x_net",
    commit = "aaf60122140d3fcf75376d319f0554393160eb50",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "org_golang_x_sys",
    commit = "1c9583448a9c3aa0f9a6a5241bf73c0bd8aafded",
    importpath = "golang.org/x/sys",
)

go_repository(
    name = "org_golang_x_text",
    commit = "f21a4dfb5e38f5895301dc265a8def02365cc3d0",
    importpath = "golang.org/x/text",
)

go_repository(
    name = "com_github_sirupsen_logrus",
    commit = "3e01752db0189b9157070a0e1668a620f9a85da2",
    importpath = "github.com/sirupsen/logrus",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "aabede6cba87e37f413b3e60ebfc214f8eeca1b0",
    importpath = "golang.org/x/crypto",
)

go_repository(
    name = "org_golang_google_grpc",
    commit = "32fb0ac620c32ba40a4626ddf94d90d12cce3455",
    importpath = "google.golang.org/grpc",
)
