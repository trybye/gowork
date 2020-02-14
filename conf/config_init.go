package conf

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/streadway/amqp"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	Mysql    Mysql
	Redis    Redis
	MqRabbit string `yaml:"mq_rabbit"`
}

type Mysql struct {
	Usr  string `yaml:"usr"`
	Pwd  string `yaml:"pwd"`
	Host string `yaml:"host"`
	Db   string `yaml:"db"`
}

type Redis struct {
	Host string `yaml:"host"`
	Pwd  string `yaml:"pwd"`
	Db   int    `yaml:"db"`
}

//enviroment
var EnviromentPc string

//mq client
var MqCh *amqp.Channel


// DB 数据库链接单例
var DB *sqlx.DB
var RedisClient *redis.Client

//user:password@tcp(127.0.0.1)/dbname?charset=utf8&parseTime=True&loc=Local
var mysqlConn string
var GlobalConfig Config

func ConfigInit() {
	//fmt.Println(filepath.Abs(filepath.Dir(os.Args[0])))
	var by []byte
	var err error
	if strings.Compare("dev", EnviromentPc) == 0 {
		by, err = ioutil.ReadFile(`conf/conf_dev.yaml`)
		fmt.Println(string(by))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("app-dev")

	} else if strings.Compare("prod", EnviromentPc) == 0 {
		by, err = ioutil.ReadFile(`conf/conf-prod.yaml`)
		fmt.Println(string(by))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("app-prod")

	} else {

		by, err = ioutil.ReadFile(`conf/conf-test.yaml`)
		fmt.Println(string(by))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("app-test")

	}

	//C := AppConfig{}
	err = yaml.Unmarshal(by, &GlobalConfig)
	if err != nil {
		log.Println(err)

	}
	fmt.Printf("--- t:\n%v\n\n", GlobalConfig)
	mysqlConn = fmt.Sprintf("%s:%s@tcp(%s)/%s", GlobalConfig.Mysql.Usr, GlobalConfig.Mysql.Pwd, GlobalConfig.Mysql.Host, GlobalConfig.Mysql.Db)
	fmt.Println(GlobalConfig.Mysql.Usr)
	///mysql
	db := sqlx.MustConnect("mysql", mysqlConn)
	// Error
	if err != nil {
		panic(err)
	}
	//设置连接池
	//空闲
	db.SetMaxIdleConns(50)

	//打开
	db.SetMaxOpenConns(1000)
	//超时
	db.SetConnMaxLifetime(time.Second * 30)

	DB = db
	///redis
	client := redis.NewClient(&redis.Options{
		Addr:     GlobalConfig.Redis.Host,
		Password: GlobalConfig.Redis.Pwd,
		DB:       GlobalConfig.Redis.Db,
	})

	result, err := client.Ping().Result()
	fmt.Println(result)
	if err != nil {
		panic(err)
	}

	RedisClient = client

	//mq
	var conn *amqp.Connection
	if strings.Compare("dev", EnviromentPc) == 0 {
		conn, err = amqp.Dial(GlobalConfig.MqRabbit) //开发环境用这个，测试用下面的
		fmt.Println("mq dev")

	} else {
		conn, err = amqp.DialConfig(GlobalConfig.MqRabbit, amqp.Config{
			Vhost: "qmrd_mq",
		})
		fmt.Println("mq test or prod")

	}

	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	MqCh = ch

	//生成doc目录，存放地址nonce
	os.Mkdir("doc", 0777)

}

func FlagUseg() {
	fmt.Println(`
	Usage:
		./pc dev|test|prod  set enviroment,dev开发，test测试,prod 生产
	`)

}

//cmd 解析,目前之加入不同的开发环境，local,dev,test,prod,暂时没部署
func InitFlag() {

	if len(os.Args) != 2 {
		FlagUseg()
		panic("")
	}
	// flag.StringVar() //没用到，感觉直接加参数也挺好的
	// flag.Parse()
	switch os.Args[1] {
	case "local", "dev":
		fmt.Println("local enviroment")
		EnviromentPc = "dev"

	case "test":
		fmt.Println("test")
		EnviromentPc = "test"
	case "prod":
		fmt.Println("prod")
		EnviromentPc = "prod"

	default:
		panic(errors.New("no this choose"))

	}
}

func Setup() {
	ConfigInit()
	r := RedisClient.Set("DealTxRuningSyncChain", "0", 0)
	x := RedisClient.Set("SyncToFrontendRuning", "0", 0)
	c := RedisClient.Set("CancelTxTimeOutRuning", "0", 0)
	u := RedisClient.Set("SyncToFrontendRuningOutput", "0", 0)

	if u.Err() != nil {
		panic(u.Err())
	}
	if r.Err() != nil {
		panic(r.Err())
	}
	if x.Err() != nil {
		panic(x.Err())
	}
	if c.Err() != nil {
		panic(x.Err())
	}
	//InsertFee()
}
