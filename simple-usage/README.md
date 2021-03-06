# simple-usage

This is a small demonstration program that quickly uses gweb to start a web-api server that contains form parameter parsing and validation

### sample code is here:

```golang
package main

import (
	"fmt"
	"github.com/yeqown/gweb"
	"github.com/yeqown/gweb/utils"
	"net/http"
	"sync"
)

func main() {
	// step1
	// define http server config
	http_server_conf := &gweb.ServerConfig{
		Logpath: "./logs",
		Port:    9012,
	}

	// step2
	// add handlers
	gweb.AddRoute(&gweb.Route{
		Path:    "/hello",
		Method:  http.MethodGet,
		Fn:      HelloGet,
		ReqPool: PoolHelloGetForm,
		ResPool: PoolHelloGetResp,
	})

	// step3
	// start htttp server
	gweb.StartHttpServer(http_server_conf)
}

// define a handler
// Get Method Demo
type HelloGetForm struct {
	Name string `schema:"name" valid:"Required" json:"name"`
	Age  int    `schema:"age" valid:"Required;Min(18)" json:"age"`
}

var PoolHelloGetForm = &sync.Pool{
	New: func() interface{} {
		return &HelloGetForm{}
	},
}

type HelloGetResp struct {
	utils.CodeInfo
	Tip string `json:"tip"`
}

var PoolHelloGetResp = &sync.Pool{
	New: func() interface{} {
		return &HelloGetResp{}
	},
}

func HelloGet(req *HelloGetForm) *HelloGetResp {
	resp := PoolHelloGetResp.Get().(*HelloGetResp)
	defer PoolHelloGetResp.Put(resp)

	resp.Tip = fmt.Sprintf(
		"Get Hello, %s! your age[%d] is valid to access",
		req.Name,
		req.Age,
	)

	utils.Response(resp, utils.NewCodeInfo(utils.CodeOk, ""))
	return resp
}

```

### screenshots

![run-and-output](https://raw.githubusercontent.com/yeqown/gweb-usage-smaples/master/simple-usage/screenshot/run.png)
![request-test](https://raw.githubusercontent.com/yeqown/gweb-usage-smaples/master/simple-usage/screenshot/request.png)
