package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_test",
    "go_library",
)

go_test(
    name = "go_default_test",
    srcs = [
        "dao_test.go",
        "mysql_test.go",
        "redis_test.go",
    ],
    embed = [":go_default_library"],
    rundir = ".",
    tags = ["automanaged"],
    deps = [
        "//app/job/main/usersuit/conf:go_default_library",
        "//vendor/github.com/go-sql-driver/mysql:go_default_library",
        "//vendor/github.com/smartystreets/goconvey/convey:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    srcs = [
        "dao.go",
        "mysql.go",
        "redis.go",
    ],
    importpath = "go-common/app/job/main/usersuit/dao/pendant",
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
    deps = [
        "//app/job/main/usersuit/conf:go_default_library",
        "//app/job/main/usersuit/model:go_default_library",
        "//library/cache/redis:go_default_library",
        "//library/database/sql:go_default_library",
        "//library/log:go_default_library",
        "//library/net/http/blademaster:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
