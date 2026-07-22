module main

go 1.11

require (
	github.com/Mr-ShiHuaYu/gf/v2 v2.8.3
	golang.org/x/time v0.8.0
)



replace (
	github.com/Mr-ShiHuaYu/gf/v2 => /workspace
	github.com/Mr-ShiHuaYu/gf/contrib/nosql/redis/v2 => /workspace/contrib/nosql/redis
	github.com/Mr-ShiHuaYu/gf/contrib/registry/file/v2 => /workspace/contrib/registry/file
	github.com/grokify/html-strip-tags-go => github.com/grokify/html-strip-tags-go v0.0.1
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	golang.org/x/net => golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	golang.org/x/sys => golang.org/x/sys v0.0.0-20200202164722-d101bd2416d5
	golang.org/x/text => golang.org/x/text v0.3.2
)
