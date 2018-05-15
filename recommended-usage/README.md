# recommended-usage

How to organize the file structure to achieve more intuitive use. and more use about controllers Requets method: 
**[Link to controller](/controllers)**

## file-tree
```sh
.
├── LICENSE
├── README.md
├── configs                     // application config files
│   ├── config.db.json
│   └── config.server.json
├── controllers                 // handlers define
│   ├── README.md
│   ├── hello.go
│   ├── rpctr
│   └── user.go
├── logs                        // log file output
│   ├── app.log
│   └── request.log
├── mainC                       // application entry
│   ├── conf.go
│   └── main.go
├── models                      // connect to db and define models with orm
│   ├── db.go
│   ├── mysql_demo.go
│   ├── redis_demo.go
│   └── redis_demo_test.go
├── services                    // handler logic functions
│   └── user.go
├── sh                          // commands include: test, deploy, build (provides by Makefile)
│   ├── Makefile
│   └── create_tables.sql
└── vendor                      // application deps
    ├── github.com
    ├── gopkg.in
    └── manifest
```