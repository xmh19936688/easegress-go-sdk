package easegress

import (
	"unsafe"

	"github.com/xmh19936688/easegress-go-sdk/easegress/test"
)

// 给host调用申请内存
//export wasm_alloc
func wasm_alloc(size int32) int32 {
	buf := make([]byte, size)
	return int32(uintptr(unsafe.Pointer(&buf[0])))
}

//export wasm_free
func wasm_free(ptr int32) {
	// TODO
}

var program Program

//export wasm_init
func wasm_init(ptr int32) {
	params := make(map[string]string)
	strs := unmarshalStringArray(ptr)
	for i := 0; i+1 < len(strs); i += 2 {
		params[strs[i]] = strs[i+1]
	}
	program = createProgram(params)

	if err := test.RunTest(); err != nil {
		Log(Error, err.Error())
	}
}

//export wasm_run
func wasm_run() int32 {
	return program.Run()
}
