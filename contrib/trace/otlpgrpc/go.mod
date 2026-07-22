module github.com/Mr-ShiHuaYu/gf/contrib/trace/otlpgrpc/v2

go 1.23.0

require (
	github.com/Mr-ShiHuaYu/gf/v2 v2.9.4
	go.opentelemetry.io/otel v1.38.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.38.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.38.0
	go.opentelemetry.io/otel/sdk v1.38.0
	google.golang.org/grpc v1.75.0
)

require (
	cloud.google.com/go/compute v1.25.1 // indirect
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/Mr-ShiHuaYu/otel-go111 v0.20.0
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/magiconair/properties v1.8.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/olekukonko/tablewriter v1.1.0 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
)

replace github.com/Mr-ShiHuaYu/gf/v2 => ../../../
