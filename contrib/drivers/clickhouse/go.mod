module github.com/Mr-ShiHuaYu/gf/contrib/drivers/clickhouse/v2

go 1.23.0

require (
	github.com/ClickHouse/clickhouse-go/v2 v2.0.15
	github.com/Mr-ShiHuaYu/gf/v2 v2.9.4
	github.com/google/uuid v1.6.0
	github.com/shopspring/decimal v1.3.1
)

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/magiconair/properties v1.8.10 // indirect
	github.com/olekukonko/tablewriter v1.1.0 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	go.opentelemetry.io/otel/sdk v1.38.0 // indirect
	golang.org/x/net v0.43.0 // indirect
)

replace (
	github.com/ClickHouse/clickhouse-go/v2 => github.com/gogf/clickhouse-go/v2 v2.0.15-compatible
	github.com/Mr-ShiHuaYu/gf/v2 => ../../../
)
