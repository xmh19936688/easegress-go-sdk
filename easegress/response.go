package easegress

//go:wasm-module easegress
//export host_resp_get_status_code
func host_resp_get_status_code() int32
func RespGetStatusCode() int32 {
	return host_resp_get_status_code()
}

//go:wasm-module easegress
//export host_resp_set_status_code
func host_resp_set_status_code(code int32)
func RespSetStatusCode(code int32) {
	host_resp_set_status_code(code)
}

//go:wasm-module easegress
//export host_resp_get_header
func host_resp_get_header(addr int32) int32
func RespGetHeader(name string) string {
	ptr := marshalString(name)
	ptr = host_resp_get_header(ptr)
	return unmarshalString(ptr)
}

//go:wasm-module easegress
//export host_resp_get_all_header
func host_resp_get_all_header() int32
func RespGetAllHeader() map[string][]string {
	ptr := host_resp_get_all_header()
	return unmarshalAllHeader(ptr)
}

//go:wasm-module easegress
//export host_resp_set_header
func host_resp_set_header(nameAddr int32, valueAddr int32)
func RespSetHeader(name string, value string) {
	namePtr := marshalString(name)
	valuePtr := marshalString(value)
	host_resp_set_header(namePtr, valuePtr)
}

//go:wasm-module easegress
//export host_resp_set_all_header
func host_resp_set_all_header(addr int32)
func RespSetAllHeader(headers map[string][]string) {
	ptr := marshalAllHeader(headers)
	host_resp_set_all_header(ptr)
}

//go:wasm-module easegress
//export host_resp_add_header
func host_resp_add_header(nameAddr int32, vlaueAddr int32)
func RespAddHeader(name string, value string) {
	namePtr := marshalString(name)
	valuePtr := marshalString(value)
	host_resp_add_header(namePtr, valuePtr)
}

//go:wasm-module easegress
//export host_resp_del_header
func host_resp_del_header(addr int32)
func RespDelHeader(name string) {
	ptr := marshalString(name)
	host_resp_del_header(ptr)
}

//go:wasm-module easegress
//export host_resp_set_cookie
func host_resp_set_cookie(addr int32)
func RespSetCookie(c Cookie) {
	ptr := marshalCookie(c)
	host_resp_set_cookie(ptr)
}

//go:wasm-module easegress
//export host_resp_get_body
func host_resp_get_body() int32
func RespGetBody() []byte {
	ptr := host_resp_get_body()
	return unmarshalData(ptr)
}

//go:wasm-module easegress
//export host_resp_set_body
func host_resp_set_body(addr int32)
func RespSetBody(body []byte) {
	addr := marshalData(body)
	host_resp_set_body(addr)
}
