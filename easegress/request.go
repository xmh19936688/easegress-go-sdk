package easegress

//go:wasm-module easegress
//export host_req_get_real_ip
func host_req_get_real_ip() int32
func ReqGetRealIp() string {
	ptr := host_req_get_real_ip()
	return unmarshalString(ptr)
}

//go:wasm-module easegress
//export host_req_get_scheme
func host_req_get_scheme() int32
func ReqGetScheme() string {
	ptr := host_req_get_scheme()
	return unmarshalString(ptr)
}

//go:wasm-module easegress
//export host_req_get_proto
func host_req_get_proto() int32
func ReqGetProto() string {
	ptr := host_req_get_proto()
	return unmarshalString(ptr)
}

//go:wasm-module easegress
//export host_req_get_method
func host_req_get_method() int32
func ReqGetMethod() string {
	ptr := host_req_get_method()
	return unmarshalString(ptr)
}

//go:wasm-module easegress
//export host_req_set_method
func host_req_set_method(addr int32)
func ReqSetMethod(method string) {
	ptr := marshalString(method)
	host_req_set_method(ptr)
}

//go:wasm-module easegress
//export host_req_get_host
func host_req_get_host() int32
func ReqGetHost() string {
	ptr := host_req_get_host()
	return unmarshalString(ptr)
}

//go:wasm-module easegress
//export host_req_set_host
func host_req_set_host(addr int32)
func ReqSetHost(host string) {
	ptr := marshalString(host)
	host_req_set_host(ptr)
}

//go:wasm-module easegress
//export host_req_get_path
func host_req_get_path() int32
func ReqGetPath() string {
	ptr := host_req_get_path()
	return unmarshalString(ptr)
}

//go:wasm-module easegress
//export host_req_set_path
func host_req_set_path(addr int32)
func ReqSetPath(path string) {
	ptr := marshalString(path)
	host_req_set_path(ptr)
}

//go:wasm-module easegress
//export host_req_get_escaped_path
func host_req_get_escaped_path() int32
func ReqGetEscapedPath() string {
	ptr := host_req_get_escaped_path()
	return unmarshalString(ptr)
}

//go:wasm-module easegress
//export host_req_get_query
func host_req_get_query() int32
func ReqGetQuery() string {
	ptr := host_req_get_query()
	return unmarshalString(ptr)
}

//go:wasm-module easegress
//export host_req_set_query
func host_req_set_query(addr int32)
func ReqSetQuery(query string) {
	ptr := marshalString(query)
	host_req_set_query(ptr)
}

//go:wasm-module easegress
//export host_req_get_fragment
func host_req_get_fragment() int32
func ReqGetFragment() string {
	ptr := host_req_get_fragment()
	return unmarshalString(ptr)
}

//go:wasm-module easegress
//export host_req_get_header
func host_req_get_header(addr int32) int32
func ReqGetHeader(name string) string {
	ptr := marshalString(name)
	ptr = host_req_get_header(ptr)
	return unmarshalString(ptr)
}

//go:wasm-module easegress
//export host_req_get_all_header
func host_req_get_all_header() int32
func ReqGetAllHeader() map[string][]string {
	ptr := host_req_get_all_header()
	return unmarshalAllHeader(ptr)
}

//go:wasm-module easegress
//export host_req_set_header
func host_req_set_header(nameAddr int32, valueAddr int32)
func ReqSetHeader(name string, value string) {
	namePtr := marshalString(name)
	valuePtr := marshalString(value)
	host_req_set_header(namePtr, valuePtr)
}

//go:wasm-module easegress
//export host_req_set_all_header
func host_req_set_all_header(addr int32)
func ReqSetAllHeader(headers map[string][]string) {
	ptr := marshalAllHeader(headers)
	host_req_set_all_header(ptr)
}

//go:wasm-module easegress
//export host_req_add_header
func host_req_add_header(nameAddr int32, valueAddr int32)
func ReqAddHeader(name string, value string) {
	namePtr := marshalString(name)
	valuePtr := marshalString(value)
	host_req_add_header(namePtr, valuePtr)
}

//go:wasm-module easegress
//export host_req_del_header
func host_req_del_header(addr int32)
func ReqDelHeader(name string) {
	ptr := marshalString(name)
	host_req_del_header(ptr)
}

//go:wasm-module easegress
//export host_req_get_cookie
func host_req_get_cookie(addr int32) int32
func ReqGetCookie(name string) *Cookie {
	ptr := marshalString(name)
	ptr = host_req_get_cookie(ptr)
	return unmarshalCookie(ptr)
}

//go:wasm-module easegress
//export host_req_get_all_cookie
func host_req_get_all_cookie() int32
func ReqGetAllCookie() []Cookie {
	var result []Cookie
	ptr := host_req_get_all_cookie()
	strs := unmarshalStringArray(ptr)
	for i := 0; i < len(strs); i++ {
		c := (&Cookie{}).unmarshal(strs[i])
		if c != nil {
			result = append(result, *c)
		}
	}

	return result
}

//go:wasm-module easegress
//export host_req_add_cookie
func host_req_add_cookie(addr int32)
func ReqAddCookie(c Cookie) {
	ptr := marshalCookie(c)
	host_req_add_cookie(ptr)
}

//go:wasm-module easegress
//export host_req_get_body
func host_req_get_body() int32
func ReqGetBody() []byte {
	ptr := host_req_get_body()
	return unmarshalData(ptr)
}

//go:wasm-module easegress
//export host_req_set_body
func host_req_set_body(addr int32)
func ReqSetBody(body []byte) {
	ptr := marshalData(body)
	host_req_set_body(ptr)
}
