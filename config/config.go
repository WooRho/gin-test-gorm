package config

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"os"
	"time"
)

type DBBASE interface {
	GetLogMode() string
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}

type Environment string

var file = "config.yaml" // 配置文件名

var Config config

type Autocode struct {
	TransferRestart bool   `mapstructure:"transfer-restart" json:"transfer-restart" yaml:"transfer-restart"`
	Root            string `mapstructure:"root" json:"root" yaml:"root"`
	Path            string `mapstructure:"path" json:"path" yaml:"path"`
	Server          string `mapstructure:"server" json:"server" yaml:"server"`
	SPlug           string `mapstructure:"server-plug" json:"server-plug" yaml:"server-plug"`
	SInitialize     string `mapstructure:"server-initialize" json:"server-initialize" yaml:"server-initialize"`
	SModel          string `mapstructure:"server-model" json:"server-model" yaml:"server-model"`
	SRequest        string `mapstructure:"server-request" json:"server-request"  yaml:"server-request"`
	SRouter         string `mapstructure:"server-router" json:"server-router" yaml:"server-router"`
	SService        string `mapstructure:"server-service" json:"server-service" yaml:"server-service"`
	SDao            string `mapstructure:"server-dao" json:"server-dao" yaml:"server-dao"`
	SServer         string `mapstructure:"server-server" json:"server-server" yaml:"server-server"`
	SAggs           string `mapstructure:"server-aggs" json:"server-aggs" yaml:"server-aggs"`
	SPb             string `mapstructure:"server-pb" json:"server-pb" yaml:"server-pb"`
	SPdiface        string `mapstructure:"server-pdiface" json:"server-pdiface" yaml:"server-pdiface"`
	SCommon         string `mapstructure:"server-common" json:"server-common" yaml:"server-common"`
}

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 地址
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	PassWord string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	PoolSize int    `mapstructure:"poolsize" json:"poolsize" yaml:"poolsize"` // 连接池大小
}

type Mysql struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                         // 地址
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`                         // 端口
	UserName     string `mapstructure:"username" json:"username" yaml:"username"`             // 用户
	PassWord     string `mapstructure:"password" json:"password" yaml:"password"`             // 密码
	DbName       string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`                   // 数据库名
	Debug        bool   `mapstructure:"debug" json:"debug" yaml:"debug"`                      // 是否开启调试，true表示开启
	Migrate      bool   `mapstructure:"migrate" json:"migrate" yaml:"migrate"`                // 是否开启自动建表，true表示开启
	WasherSwitch bool   `mapstructure:"washerswitch" json:"washerswitch" yaml:"washerswitch"` // 是否开启清洗数据开关，true表示开启
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`             // 是否开启全局禁用复数，true表示开启
	Charset      string `mapstructure:"charset" json:"charset" yaml:"charset"`
	Collation    string `mapstructure:"collation" json:"collation" yaml:"collation"`
	MaxIdleConns int    `mapstructure:"maxidleconns" json:"maxidleconns" yaml:"maxidleconns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"maxopenconns" json:"maxopenconns" yaml:"maxopenconns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"logmode" json:"logmode" yaml:"logmode"`                // 是否开启Gorm全局日志
}

type TenantManagement struct {
	TenantPhone string `mapstructure:"tenantphone" json:"tenantphone" yaml:"tenantphone"`
}

type MysqlSlave struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 地址
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	UserName string `mapstructure:"username" json:"username" yaml:"username"` // 用户
	PassWord string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	DbName   string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`       // 数据库名
}

type Server struct {
	ServiceName string `mapstructure:"servicename" json:"servicename" yaml:"servicename"` // 服务名
	Debug       bool   `mapstructure:"debug" json:"debug" yaml:"debug"`                   // gin是否开启debug模式
	Environment string `mapstructure:"environment" json:"environment" yaml:"environment"` // 服务部署环境  可以是 local, dev, test, release （默认release）
	AuthOn      bool   `mapstructure:"authon" json:"authon" yaml:"authon"`
	Version     string `mapstructure:"version" json:"version" yaml:"version"`
	WebhookURL  string `mapstructure:"webhookurL" json:"webhookurL" yaml:"webhookurL"`
}

type Web struct {
	Port string `mapstructure:"port" json:"port" yaml:"port"` // 服务使用端口
}

type Html2Image struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"` // 地址
	Port int    `mapstructure:"port" json:"port" yaml:"port"` // 端口
}

type RabbitMQ struct {
	User       string
	Password   string
	Host       string
	Port       string
	Dead       string
	RoutingKey string
}

type WebSocket struct {
	PingWait string
}

type CORS struct {
	Mode      string          `mapstructure:"mode" json:"mode" yaml:"mode"`
	Whitelist []CORSWhitelist `mapstructure:"whitelist" json:"whitelist" yaml:"whitelist"`
}

type CORSWhitelist struct {
	AllowOrigin      string `mapstructure:"allow-origin" json:"allow-origin" yaml:"allow-origin"`
	AllowMethods     string `mapstructure:"allow-methods" json:"allow-methods" yaml:"allow-methods"`
	AllowHeaders     string `mapstructure:"allow-headers" json:"allow-headers" yaml:"allow-headers"`
	ExposeHeaders    string `mapstructure:"expose-headers" json:"expose-headers" yaml:"expose-headers"`
	AllowCredentials bool   `mapstructure:"allow-credentials" json:"allow-credentials" yaml:"allow-credentials"`
}

type Alibaba struct {
	AccessKey    string `mapstructure:"accesskey" json:"accesskey" yaml:"accesskey"`
	AccessSecret string `mapstructure:"accesssecret" json:"accesssecret" yaml:"accesssecret"`
	SignName     string `mapstructure:"signname" json:"signname" yaml:"signname"`
	TemplateCode string `mapstructure:"templatecode" json:"templatecode" yaml:"templatecode"`
}

type config struct {
	Redis      Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	MysqlSlave MysqlSlave `mapstructure:"mysqlslave" json:"mysqlslave" yaml:"mysqlslave"`
	Server     Server     `mapstructure:"server" json:"server" yaml:"server"`
	Web        Web        `mapstructure:"web" json:"web" yaml:"web"`
	Html2Image Html2Image `mapstructure:"html2image" json:"html2image" yaml:"html2image"`
	RabbitMQ   RabbitMQ   `mapstructure:"rabbitmq" json:"rabbitmq" yaml:"rabbitmq"`
	WebSocket  WebSocket  `mapstructure:"websocket" json:"websocket" yaml:"websocket"`
	// auto
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// 跨域配置
	Cors             CORS             `mapstructure:"cors" json:"cors" yaml:"cors"`
	TenantManagement TenantManagement `mapstructure:"tenantmanagement" json:"tenantmanagement" yaml:"tenantmanagement"`
	Alibaba          Alibaba          `mapstructure:"alibaba" json:"alibaba" yaml:"alibaba"`
}

// 加载配置文件
func loadConfig() {

	args := os.Args
	if len(args) > 1 {
		file = args[1]
	}

	Viper(file)
	ToJson := func(v interface{}) string {
		bs, err := jsoniter.MarshalIndent(v, "", "   ")
		if err != nil {
			return ""
		}
		return string(bs)
	}

	fmt.Printf("[%s] "+time.Now().Format("2006-01-02 15:04:05")+"\r\n"+"%s"+"\r\n",
		"Config",
		fmt.Sprintf("%s\n%s\n%v\n%s\n",
			"Init Load Config Success",
			"-----------------------------------",
			ToJson(Config),
			"-----------------------------------",
		),
	)
}

func Init() {
	loadConfig()
}

func Stop() {

}
