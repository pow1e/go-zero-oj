package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql   Mysql
	Redis   Redis
	Auth    Auth
	Captcha Captcha
}

type Mysql struct {
	UserName string
	Password string
	Path     string
	Port     string
	DataBase string
	Config   string
}
type Redis struct {
	Host string
	Pass string
	Type string
}

type Auth struct {
	AccessSecret  string
	AccessExpire  int64
	RefreshExpire int64
}

type Captcha struct {
	KeyLong        int
	Width          int
	Height         int
	Timeout        int
	MaxTime        int // 一天请求次数
	RefreshMaxTime int // 单个ip请求验证码的
}
