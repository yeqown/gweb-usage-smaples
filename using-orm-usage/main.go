package main

import (
	"github.com/yeqown/gweb"
	"net/http"
)

func main() {
	http_server_conf := &gweb.ServerConfig{
		Logpath: "./logs",
		Port:    9012,
	}

	DBInit()

	// add handler
	gweb.AddRoute(&gweb.Route{
		Path:    "/user/register",
		Method:  http.MethodPost,
		Fn:      UserRegister,
		ReqPool: PoolUserRegisterForm,
		ResPool: PoolUserRegisterResp,
	})

	gweb.AddRoute(&gweb.Route{
		Path:    "/recipe/new",
		Method:  http.MethodPost,
		Fn:      AppendRecipe,
		ReqPool: PoolAppendRecipeForm,
		ResPool: PoolAppendRecipeResp,
	})

	gweb.StartHttpServer(http_server_conf)
}
