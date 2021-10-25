package easegress

import (
	"encoding/binary"
	"strings"
	"unsafe"
)

var sizeI32 = int32(unsafe.Sizeof(int32(1)))

// 从前4个字节中读取长度
func readSize(ptr int32) int32 {
	var buf []byte
	for i := int32(0); i < sizeI32; i++ {
		p := (*byte)(unsafe.Pointer(uintptr(ptr + i)))
		buf = append(buf, *p)
	}

	//ptr0 := (*byte)(unsafe.Pointer(uintptr(ptr)))
	//ptr1 := (*byte)(unsafe.Pointer(uintptr(ptr + 1)))
	//ptr2 := (*byte)(unsafe.Pointer(uintptr(ptr + 2)))
	//ptr3 := (*byte)(unsafe.Pointer(uintptr(ptr + 3)))
	//buf := []byte{*ptr0, *ptr1, *ptr2, *ptr3}

	return int32(binary.LittleEndian.Uint32(buf[:]))
}

func marshalData(data []byte) int32 {
	buf := make([]byte, len(data)+int(sizeI32))
	binary.LittleEndian.PutUint32(buf[0:], uint32(len(data)))
	copy(buf[sizeI32:], data)
	return int32(uintptr(unsafe.Pointer(&buf[0])))
}

func unmarshalData(ptr int32) []byte {
	size := readSize(ptr)
	buf := make([]byte, size)
	for i := int32(0); i < size; i++ {
		buf[i] = *(*byte)(unsafe.Pointer(uintptr(ptr + sizeI32 + i)))
	}

	wasm_free(ptr)
	return buf
}

func marshalString(str string) int32 {
	// 创建一个byte数组，前4个字节放字符串长度，后面放字符串，最后是空白字符
	s := make([]byte, len(str)+int(sizeI32)+1)
	binary.LittleEndian.PutUint32(s[0:], uint32(len(str)+1))
	copy(s[sizeI32:], str)
	s[4+len(str)] = 0
	return int32(uintptr(unsafe.Pointer(&s[0])))
}

func unmarshalString(ptr int32) string {
	size := readSize(ptr)

	// 读取字符串数据
	buf := make([]byte, size-1) // size-1是因为最后一个为空白字符
	for i := int32(0); i < size-1; i++ {
		buf[i] = *(*byte)(unsafe.Pointer(uintptr(ptr + int32(sizeI32) + i)))
	}

	wasm_free(ptr)
	return string(buf)
}

func unmarshalStringArray(ptr int32) []string {
	size := readSize(ptr)
	result := make([]string, size)
	for i := int32(0); i < size; i++ {
		result[i] = unmarshalString(ptr + int32(sizeI32))
	}

	wasm_free(ptr)
	return result
}

func marshalAllHeader(headers map[string][]string) int32 {
	str := ""
	for k, v := range headers {
		for i := 0; i < len(v); i++ {
			str += k + ": " + v[i] + "\r\n"
		}
	}
	return marshalString(str)
}

func unmarshalAllHeader(ptr int32) map[string][]string {
	str := unmarshalString(ptr)
	headers := strings.Split(str, "\r\n")
	result := make(map[string][]string)

	for i := 0; i < len(headers); i++ {
		kv := strings.Split(headers[i], ": ")
		if len(kv) != 2 {
			continue
		}

		result[kv[0]] = append(result[kv[0]], kv[1])
	}

	return result
}

func marshalCookie(c Cookie) int32 {
	str := c.marshal()
	return marshalString(str)
}

func unmarshalCookie(ptr int32) *Cookie {
	str := unmarshalString(ptr)
	if str == "" {
		return nil
	}
	return (&Cookie{}).unmarshal(str)
}
