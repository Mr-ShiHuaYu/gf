module github.com/Mr-ShiHuaYu/gf/contrib/rpc/grpcx/v2

go 1.23.0

require (
	github.com/Mr-ShiHuaYu/gf/contrib/registry/file/v2 v2.9.4
	github.com/Mr-ShiHuaYu/gf/v2 v2.9.4
	github.com/Mr-ShiHuaYu/otel-go111 v0.20.0
	github.com/gogf/gf/contrib/registry/file/v2 v2.10.2
	github.com/gogf/gf/contrib/rpc/grpcx/v2 v2.10.2
	github.com/golang/mock v1.6.0 // indirect
	go.opentelemetry.io/otel v1.38.0
	go.opentelemetry.io/otel/trace v1.38.0
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	google.golang.org/grpc v1.64.1
	google.golang.org/protobuf v1.34.2
)

replace (
	github.com/Mr-ShiHuaYu/gf/contrib/registry/file/v2 => ../../registry/file/
	github.com/Mr-ShiHuaYu/gf/v2 => ../../../
)
