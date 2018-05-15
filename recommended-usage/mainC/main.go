package main

import (
	"flag"

	"github.com/yeqown/gweb"
	. "github.com/yeqown/gweb/logger"
	"net/http"
	"net/rpc"

	ctr "recommended-usage/controllers"
	rpctr "recommended-usage/controllers/rpctr"
	. "recommended-usage/models"
)

var (
	db_conf     = flag.String("db_conf", "./configs/config.db.json", "-db_conf filename")
	server_conf = flag.String("server_conf", "./configs/config.server.json", "-server_conf filename")
)

func main() {
	flag.Parse()

	if err := loadDBConf(*db_conf); err != nil {
		AppL.Fatal(err.Error())
	}

	if err := loadServerConf(*server_conf); err != nil {
		AppL.Fatal(err.Error())
	}

	// connect to dbs
	ConnectMysql(db_ins.MysqlC)
	ConnectMongo(db_ins.MgoC)
	ConnectRedis(db_ins.RedisC)

	// to start Servers enum "rpc & http"
	rpc_s := gweb.GetRpcServer()
	registerRpcService(rpc_s)
	go gweb.StartRpcServer(server_ins.RpcC)

	registerRouter()
	gweb.StartHttpServer(server_ins.HttpC)
}

// registerRouter
func registerRouter() {
	gweb.AddRoute(
		&gweb.Route{"/hello", http.MethodGet, ctr.HelloGet,
			ctr.PoolHelloGetForm,
			ctr.PoolHelloGetResp,
		})
	gweb.AddRoute(
		&gweb.Route{"/hello", http.MethodPost, ctr.HelloPost,
			ctr.PoolHelloPostForm,
			ctr.PoolHelloPostResp,
		})
	gweb.AddRoute(
		&gweb.Route{"/hello", http.MethodPut, ctr.HelloPut,
			ctr.PoolHelloPutForm,
			ctr.PoolHelloPutResp,
		})
	gweb.AddRoute(
		&gweb.Route{"/hello/json", http.MethodPost, ctr.HelloJsonBody,
			ctr.PoolHelloJsonBodyForm,
			ctr.PoolHelloJsonBodyResp,
		})
	gweb.AddRoute(
		&gweb.Route{"/hello/file", http.MethodPost, ctr.HelloFile,
			ctr.PoolHelloFileForm,
			ctr.PoolHelloFileResp,
		})
	gweb.AddRoute(
		&gweb.Route{"/user/register", http.MethodPost, ctr.RegisterUserPost,
			ctr.PoolRegUserForm,
			ctr.PoolRegUserResp,
		})
	gweb.AddRoute(
		&gweb.Route{"/user/login", http.MethodPost, ctr.LoginUserPost,
			ctr.PoolLogUserForm,
			ctr.PoolLogUserResp,
		})
}

// registerRpcService register All service for rpc server
func registerRpcService(s *rpc.Server) {
	AppL.Info("registerRpcService doing...")
	calculator := new(rpctr.Calculator)
	s.Register(calculator)
}
