# using-orm-usage

gweb is using `gorm` and `mgo.v2`. you can ref them usage from following link:

* [gorm](http://gorm.io/)
* [mgo.v2](https://gopkg.in/mgo.v2)

also gweb will include `redis` cache.

this sample is talking *How to use gweb inner orm to connect to DB and op them with model*. first of all, you must connect to your db, like this:
```golang
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
```

then, you have to define your model, like you use with `gorm` and `mgo.v2`. there is a demo:
```golang
// User
// mysql Model demo
type User struct {
	ID         int64     `gorm:"primary_key;column:id"`
	Name       string    `gorm:"column:name"`
	Age        int       `gorm:"column:age"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
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
```

finnally, just use it following the orm doc.
```golang
// NewUser to create a record in Mysql
func NewUser(ui *User) error {
	uic := NewUserColl()
	if err := uic.Create(ui).Error; err != nil {
		return err
	}
	return nil
}

// NewRecipe to create a record in Mongo
func NewRecipe(r *Recipe) error {
	rc := NewRecipeDetailColl()
	if err := rc.Insert(r); err != nil {
		return err
	}
	return nil
}
```

# screenshots

![run](https://raw.githubusercontent.com/yeqown/gweb-usage-smaples/master/using-orm-usage/screenshots/run.png)
![mysql-result](https://raw.githubusercontent.com/yeqown/gweb-usage-smaples/master/using-orm-usage/screenshots/mysql-result.png)
![mgo-result](https://raw.githubusercontent.com/yeqown/gweb-usage-smaples/master/using-orm-usage/screenshots/mgo-result.png)

all code are in current-repo/using-orm-usage.
til now, gweb support-orm are so simple, more functions are coming.
