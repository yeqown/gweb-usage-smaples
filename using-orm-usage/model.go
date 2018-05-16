package main

import (
	"github.com/jinzhu/gorm"
	"github.com/yeqown/gweb"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// DB Init
// connect to db with configs
func DBInit() {
	mysql_conf := &gweb.MysqlConfig{
		Addr:      "yeqown:yeqown@tcp(127.0.0.1:3306)/testdb?",
		Loc:       "Local",
		Charset:   "utf8",
		Pool:      20,
		ParseTime: "true",
	}
	gweb.ConnectMysql(mysql_conf)

	redis_conf := &gweb.RedisConfig{
		Addr:        "127.0.0.1:6379",
		DB:          1,
		Password:    "",
		PoolSize:    20,
		Timeout:     2,
		MaxActive:   500,
		MaxIdle:     50,
		IdleTimeout: 600,
		Wait:        true,
	}
	gweb.ConnectRedis(redis_conf)

	mongo_conf := &gweb.MongoConfig{
		Addrs:     "localhost:27017",
		Timeout:   5,
		PoolLimit: 20,
	}
	gweb.ConnectMongo(mongo_conf)
}

// User
// mysql Model demo
type User struct {
	ID         int64     `gorm:"primary_key;column:id"`
	Name       string    `gorm:"column:name"`
	Age        int       `gorm:"column:age"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

type UserColl struct {
	*gorm.DB
}

func NewUserColl() *UserColl {
	return &UserColl{
		DB: gweb.GetMysqlDB().Model(&User{}),
	}
}

// Recipe
// Mongo Model demo
type Recipe struct {
	Id         bson.ObjectId `bson:"_id,omitempty" json:"id"`        // id
	Name       string        `bson:"name" json:"name"`               // 名字
	Cat        string        `bson:"cat" json:"cat"`                 // 分类
	CreateTime time.Time     `bson:"create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time     `bson:"update_time" json:"update_time"` // 更新时间
}

type RecipeColl struct {
	*mgo.Collection
}

func NewRecipeDetailColl() *RecipeColl {
	return &RecipeColl{
		Collection: gweb.NewMongoColl("recipe", "recipe_detail"),
	}
}
