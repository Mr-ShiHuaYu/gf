module github.com/Mr-ShiHuaYu/gf/contrib/drivers/mysql/v2

go 1.11

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/Mr-ShiHuaYu/gf/v2 v2.16.10
)

replace (
	github.com/Mr-ShiHuaYu/gf/v2 => ../../../
	go.opentelemetry.io/otel => ../../../third_party/otel
	go.opentelemetry.io/otel/sdk => ../../../third_party/otel
	go.opentelemetry.io/otel/trace => ../../../third_party/otel
	go.opentelemetry.io/otel/semconv => ../../../third_party/otel
	go.opentelemetry.io/otel/metric => ../../../third_party/otel
)
