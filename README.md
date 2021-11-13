# easegress-go-sdk

## Usage

```shell
# compile demo
tinygo build -o bk.wasm -target wasi ./examples/flashsale/

# launch easegress-server
/bin/easegress-server --debug

# create object
cat <<EOF | /bin/egctl object create
---
kind: HTTPServer
name: http-server
port: 10080
keepAlive: true
https: false
rules:
- paths:
  - pathPrefix: /flashsale
    backend: flash-sale-pipeline

---
name: flash-sale-pipeline
kind: HTTPPipeline
flow:
- filter: wasm
filters:
- name: wasm
  kind: WasmHost
  maxConcurrency: 2
  code: bk.wasm
  timeout: 100ms
  parameters:
    startTime: "2021-11-08 21:50:21"
    blockRatio: 0.5
    maxPermission: 3
EOF

# curl
curl http://localhost:10080/flashsale -H Authorization:xmh
```