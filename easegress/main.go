package easegress

type Program interface {
	Init(params map[string]string)
	Run() int32
}

var programFactory func(params map[string]string) Program

func createProgram(params map[string]string) Program {
	return programFactory(params)
}

func RegisterProgramFactory(factory func(params map[string]string) Program) {
	programFactory = factory
}

//go:wasm-module easegress
//export host_add_tag
func host_add_tag(addr int32)
func AddTag(tag string) {
	ptr := marshalString(tag)
	host_add_tag(ptr)
}

type LogLevel int32

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

//go:wasm-module easegress
//export host_log
func host_log(level int32, addr int32)
func Log(level LogLevel, msg string) {
	ptr := marshalString(msg)
	host_log(int32(level), ptr)
}

//go:wasm-module easegress
//export host_get_unix_time_in_ms
func host_get_unix_time_in_ms() int64
func getUnixTimeInMs() int64 {
	return host_get_unix_time_in_ms()
}

//go:wasm-module easegress
//export host_rand
func host_rand() float64
func Rand() float64 {
	return host_rand()
}
