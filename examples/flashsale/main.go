package main

import (
	"time"

	"github.com/xmh19936688/easegress-go-sdk/easegress"
	"github.com/xmh19936688/easegress-go-sdk/easegress/util"
)

func init() {
	easegress.RegisterProgramFactory(func(params map[string]string) easegress.Program {
		program := &FlashSale{}
		program.Init(params)
		return program
	})
}

// tinygo build -o bk.wasm -target wasi ./examples/flashsale/
// with test: tinygo build -tags "test" -o bk.wasm -target wasi ./examples/flashsale/
// curl http://localhost:10080/flashsale -H Authorization:xmh
func main() {
	// keep empty
}

type FlashSale struct {
	easegress.Program
	startTime     time.Time
	blockRatio    float64
	maxPermission int32
}

func (program *FlashSale) Init(params map[string]string) {
	if v, ok := params["startTime"]; ok {
		if t, err := util.ParseTime("2006-01-02 15:04:05", v); err == nil {
			program.startTime = t
		}
	}

	if v, ok := params["blockRatio"]; ok {
		if f, err := util.ParseFloat(v, 64); err == nil {
			program.blockRatio = f
		}
	}

	if v, ok := params["maxPermission"]; ok {
		if i, err := util.ParseInt(v, 10, 32); err == nil {
			program.maxPermission = int32(i)
		}
	}
}

func (program *FlashSale) Run() int32 {
	easegress.RespSetHeader("Content-Type", "application/json")

	if time.Unix(easegress.GetUnixTimeInMs()/1000, 0).Before(program.startTime) {
		easegress.RespSetBody([]byte("not start yet."))
		return 1
	}

	var id = easegress.ReqGetHeader("Authorization")
	easegress.Log(easegress.Error, "id: "+id)
	if easegress.ClusterGetString("id/"+id) == "true" {
		easegress.RespSetBody([]byte("true."))
		return 0
	}

	count := easegress.ClusterCountKey("id/")
	if count < program.maxPermission {
		if easegress.Rand() > program.blockRatio {
			easegress.ClusterPutString("id/"+id, "true")
			easegress.RespSetBody([]byte("lucky from wasm."))
			return 0
		}
	}

	easegress.RespSetBody([]byte("sold out."))
	return 0
}
