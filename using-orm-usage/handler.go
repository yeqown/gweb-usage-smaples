package main

import (
	"github.com/yeqown/gweb/utils"
	"sync"
	"time"
)

// define a handler
// Post Method to op Mysql
type UserRegisterForm struct {
	Name string `schema:"name" valid:"Required" json:"name"`
	Age  int    `schema:"age" valid:"Required;Min(18)" json:"age"`
}

var PoolUserRegisterForm = &sync.Pool{
	New: func() interface{} {
		return &UserRegisterForm{}
	},
}

type UserRegisterResp struct {
	utils.CodeInfo
}

var PoolUserRegisterResp = &sync.Pool{
	New: func() interface{} {
		return &UserRegisterResp{}
	},
}

func UserRegister(req *UserRegisterForm) *UserRegisterResp {
	resp := PoolUserRegisterResp.Get().(*UserRegisterResp)
	defer PoolUserRegisterResp.Put(resp)

	if err := NewUser(&User{
		Name:       req.Name,
		Age:        req.Age,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}); err != nil {
		utils.Response(resp, utils.NewCodeInfo(utils.CodeSystemErr, err.Error()))
		return resp
	}

	utils.Response(resp, utils.NewCodeInfo(utils.CodeOk, ""))
	return resp
}

// define a handler
// Post Method to op Mongo
type AppendRecipeForm struct {
	Name string `schema:"name" valid:"Required" json:"name"`
	Cat  string `schema:"cat" valid:"Required;" json:"age"`
}

var PoolAppendRecipeForm = &sync.Pool{
	New: func() interface{} {
		return &AppendRecipeForm{}
	},
}

type AppendRecipeResp struct {
	utils.CodeInfo
}

var PoolAppendRecipeResp = &sync.Pool{
	New: func() interface{} {
		return &AppendRecipeResp{}
	},
}

func AppendRecipe(req *AppendRecipeForm) *AppendRecipeResp {
	resp := PoolAppendRecipeResp.Get().(*AppendRecipeResp)
	defer PoolAppendRecipeResp.Put(resp)

	if err := NewRecipe(&Recipe{
		Name:       req.Name,
		Cat:        req.Cat,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}); err != nil {
		utils.Response(resp, utils.NewCodeInfo(utils.CodeSystemErr, err.Error()))
		return resp
	}

	utils.Response(resp, utils.NewCodeInfo(utils.CodeOk, ""))
	return resp
}
