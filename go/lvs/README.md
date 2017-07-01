Go 解析 LVS 配置文件
===========================

> 本代码是为了解答 [golang 解析CFG文件，有没有类似解析json的包](https://segmentfault.com/q/1010000009553642) 问题而编写。

## 安装

```
go get github.com/chekun/code-snippets/go/lvs
```

## 使用

```go
package main

import (
	"fmt"

	"encoding/json"

	"github.com/chekun/code-snippets/go/lvs"
)

func main() {

	config := `
		virtual_server {
			label JuXG-HTTPS
			ip 123.103.91.122
			port 443
			lb_algo rr
			lb_kind tun
			protocol TCP

			real_server {
				label RealServer1
				ip 123.103.91.115
				port 443
				weight 100
				HTTP_GET {
					check_port 80
					path 'health'
					http_recv 'Welcome to nginx'
					connect_timeout 3
					}
			}
			real_server {
				label RealServer2
				ip 123.103.91.116
				port 443
				weight 100
				HTTP_GET {
					check_port 80
					path 'health'
					http_recv 'Welcome to nginx'
					connect_timeout 3
					}
			}
		}`

	r, _ := lvs.Unmarshal(config)
	rJSON, _ := json.Marshal(r)
	fmt.Println(string(rJSON))
}

```

> 未实现 `Marshal` 方法