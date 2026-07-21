package semconv

import "github.com/Mr-ShiHuaYu/otel-go111/attribute"

const (
	ServiceNameKey           = attribute.Key("service.name")
	ServiceVersionKey        = attribute.Key("service.version")
	ServiceInstanceIDKey     = attribute.Key("service.instance.id")
	ServiceNamespaceKey      = attribute.Key("service.namespace")

	HostNameKey              = attribute.Key("host.name")
	HostIDKey                = attribute.Key("host.id")
	HostArchKey              = attribute.Key("host.arch")
	HostImageNameKey         = attribute.Key("host.image.name")
	HostImageIDKey           = attribute.Key("host.image.id")
	HostImageVersionKey      = attribute.Key("host.image.version")
	HostTypeKey              = attribute.Key("host.type")
	HostRuntimeNameKey       = attribute.Key("host.runtime.name")
	HostRuntimeVersionKey    = attribute.Key("host.runtime.version")
	HostRuntimeDescriptionKey = attribute.Key("host.runtime.description")

	HTTPMethodKey            = attribute.Key("http.method")
	HTTPURLKey               = attribute.Key("http.url")
	HTTPTargetKey            = attribute.Key("http.target")
	HTTPHostKey              = attribute.Key("http.host")
	HTTPSchemeKey            = attribute.Key("http.scheme")
	HTTPStatusCodeKey        = attribute.Key("http.status_code")
	HTTPStatusTextKey        = attribute.Key("http.status_text")
	HTTPFlavorKey            = attribute.Key("http.flavor")
	HTTPUserAgentKey         = attribute.Key("http.user_agent")
	HTTPRouteKey             = attribute.Key("http.route")
	HTTPServerNameKey        = attribute.Key("http.server_name")
	HTTPClientIPKey          = attribute.Key("http.client_ip")

	DBSystemKey              = attribute.Key("db.system")
	DBConnectionStringKey    = attribute.Key("db.connection_string")
	DBUserKey                = attribute.Key("db.user")
	DBNameKey                = attribute.Key("db.name")
	DBStatementKey           = attribute.Key("db.statement")
	DBOperationKey           = attribute.Key("db.operation")
	DBTableKey               = attribute.Key("db.sql.table")

	RPCSystemKey             = attribute.Key("rpc.system")
	RPCServiceKey            = attribute.Key("rpc.service")
	RPCMethodKey             = attribute.Key("rpc.method")
	RPCGRPCStatusCodeKey     = attribute.Key("rpc.grpc.status_code")

	ExceptionEventName       = "exception"
	ExceptionTypeKey         = attribute.Key("exception.type")
	ExceptionMessageKey      = attribute.Key("exception.message")
	ExceptionStacktraceKey   = attribute.Key("exception.stacktrace")
	ExceptionEscapedKey      = attribute.Key("exception.escaped")
)

var (
	DBStatement = attribute.Key("db.statement")
	DBSystem    = attribute.Key("db.system")
	DBName      = attribute.Key("db.name")
	HostName    = attribute.Key("host.name")
	HTTPMethod  = attribute.Key("http.method")
	HTTPUrl     = attribute.Key("http.url")
	HTTPTarget  = attribute.Key("http.target")
	HTTPHost    = attribute.Key("http.host")
	HTTPScheme  = attribute.Key("http.scheme")
	HTTPStatusCode = attribute.Key("http.status_code")
	HTTPFlavor  = attribute.Key("http.flavor")
	HTTPUserAgent = attribute.Key("http.user_agent")
	HTTPRoute   = attribute.Key("http.route")
)
