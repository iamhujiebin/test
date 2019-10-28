package config

var LOCAL_CLOUDAC string
var HTTP_TIMEOUT_SECOND int


func init() {
	LOCAL_CLOUDAC = GetGlobalStringValue("cloudac_server", "")
	HTTP_TIMEOUT_SECOND = GetGlobalIntValue(COMMON_HTTP_CLIENT_TIMEOUT_SECONDS, 10)
}
