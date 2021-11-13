package easegress

import (
	"encoding/binary"
	"strings"
	"unsafe"
)

var sizeI32 = int32(unsafe.Sizeof(int32(1)))

// read length from the first 4 bytes
func readSize(ptr int32) int32 {
	var buf []byte
	for i := int32(0); i < sizeI32; i++ {
		p := (*byte)(unsafe.Pointer(uintptr(ptr + i)))
		buf = append(buf, *p)
	}

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
	// make a byte array, the first 4 bytes for length of string, string data then, empty byte at last.
	s := make([]byte, len(str)+int(sizeI32)+1)
	binary.LittleEndian.PutUint32(s[0:], uint32(len(str)+1))
	copy(s[sizeI32:], str)
	s[int(sizeI32)+len(str)] = 0
	return int32(uintptr(unsafe.Pointer(&s[0])))
}

func unmarshalString(ptr int32) string {
	size := readSize(ptr)

	// read string data
	buf := make([]byte, size-1) // `size-1` to ignore the empty byte at last
	for i := int32(0); i < size-1; i++ {
		buf[i] = *(*byte)(unsafe.Pointer(uintptr(ptr + int32(sizeI32) + i)))
	}

	wasm_free(ptr)
	return string(buf)
}

func unmarshalStringArray(ptr int32) []string {
	size := readSize(ptr)
	result := make([]string, size)
	offset := int32(0)
	for i := int32(0); i < size; i++ {
		result[i] = unmarshalString(ptr + sizeI32 + offset)
		offset += int32(len(result[i])) + sizeI32 + 1
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
