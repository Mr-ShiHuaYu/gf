module main

go 1.11

require (
	github.com/Mr-ShiHuaYu/gf/v2 v2.8.3
	github.com/go-redis/redis/v7 v7.4.1
)

replace (
	github.com/BurntSushi/toml => github.com/BurntSushi/toml v0.3.1
	github.com/Mr-ShiHuaYu/gf/v2 => /workspace
	github.com/clbanning/mxj/v2 => github.com/clbanning/mxj/v2 v2.5.0
	github.com/emirpasic/gods => github.com/emirpasic/gods v1.18.1
	github.com/fsnotify/fsnotify => github.com/fsnotify/fsnotify v1.4.7
	github.com/gorilla/websocket => github.com/gorilla/websocket v1.4.2
	github.com/grokify/html-strip-tags-go => github.com/grokify/html-strip-tags-go v0.0.1
	github.com/magiconair/properties => github.com/magiconair/properties v1.8.0
	github.com/mattn/go-colorable => github.com/mattn/go-colorable v0.1.1
	github.com/mattn/go-isatty => github.com/mattn/go-isatty v0.0.4
	github.com/mattn/go-runewidth => github.com/mattn/go-runewidth v0.0.4
	github.com/olekukonko/tablewriter => github.com/olekukonko/tablewriter v0.0.5
	golang.org/x/net => golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/sys => golang.org/x/sys v0.0.0-20200202164722-d101bd2416d5
	golang.org/x/text => golang.org/x/text v0.3.2
)
