package easegress

//go:wasm-module easegress
//export host_cluster_get_binary
func host_cluster_get_binary(addr int32) int32
func ClusterGetBinary(key string) []byte {
	var ptr = marshalString(key)
	ptr = host_cluster_get_binary(ptr)
	return unmarshalData(ptr)
}

//go:wasm-module easegress
//export host_cluster_put_binary
func host_cluster_put_binary(keyAddr int32, valAddr int32)
func ClusterPutBinary(key string, val []byte) {
	var ptrKey = marshalString(key)
	var ptrVal = marshalData(val)
	host_cluster_put_binary(ptrKey, ptrVal)
}

//go:wasm-module easegress
//export host_cluster_get_string
func host_cluster_get_string(addr int32) int32
func ClusterGetString(key string) string {
	var ptr = marshalString(key)
	ptr = host_cluster_get_string(ptr)
	return unmarshalString(ptr)
}

//go:wasm-module easegress
//export host_cluster_put_string
func host_cluster_put_string(keyAddr int32, valAddr int32)
func ClusterPutString(key string, val string) {
	var ptrKey = marshalString(key)
	var ptrVal = marshalString(val)
	host_cluster_put_string(ptrKey, ptrVal)
}

//go:wasm-module easegress
//export host_cluster_get_integer
func host_cluster_get_integer(addr int32) int64
func ClusterGetInteger(key string) int64 {
	var ptr = marshalString(key)
	return host_cluster_get_integer(ptr)
}

//go:wasm-module easegress
//export host_cluster_put_integer
func host_cluster_put_integer(addr int32, val int64)
func ClusterPutInteger(key string, val int64) {
	var ptr = marshalString(key)
	host_cluster_put_integer(ptr, val)
}

//go:wasm-module easegress
//export host_cluster_add_integer
func host_cluster_add_integer(addr int32, val int64) int64
func ClusterAddInteger(key string, val int64) int64 {
	var ptr = marshalString(key)
	return host_cluster_add_integer(ptr, val)
}

//go:wasm-module easegress
//export host_cluster_get_float
func host_cluster_get_float(addr int32) float64
func ClusterGetFloat(key string) float64 {
	var ptr = marshalString(key)
	return host_cluster_get_float(ptr)
}

//go:wasm-module easegress
//export host_cluster_put_float
func host_cluster_put_float(addr int32, val float64)
func ClusterPutFloat(key string, val float64) {
	var ptr = marshalString(key)
	host_cluster_put_float(ptr, val)
}

//go:wasm-module easegress
//export host_cluster_add_float
func host_cluster_add_float(addr int32, val float64) float64
func ClusterAddFloat(key string, val float64) float64 {
	var ptr = marshalString(key)
	return host_cluster_add_float(ptr, val)
}

//go:wasm-module easegress
//export host_cluster_count_key
func host_cluster_count_key(addr int32) int32
func ClusterCountKey(prefix string) int32 {
	var ptr = marshalString(prefix)
	return host_cluster_count_key(ptr)
}
