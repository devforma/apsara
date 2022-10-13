package core

import "time"

type Bcc struct {
	Endpoint   string
	AppName    string
	AppKey     string
	Account    string
	AccountKey string
}

type Config struct {
	AsapiEndpoint     string // 天基 -> 报表 -> 服务注册变量 -> asapi.public.endpoint
	Protocol          string
	RegionID          string
	AccessKeyID       string // 运营控制台运营账号/组织的ak
	AccessKeySecret   string // 运营控制台运营账号/组织的sk
	AasSecret         string // 天基 -> 报表 -> 服务注册变量 -> 搜索 "baseService-aas" -> 得到 "bid.accesssecret"
	Bcc               Bcc    // 从天基系统报表的"服务注册变量"中搜索认证变量: bcc_endpoint super_app_name super_app_key super_account_account super_account_secret_key
	ConnectionTimeout time.Duration
	ReadTimeout       time.Duration
	MaxIdleConns      int
	UserAgent         string
	EnableLog         bool
}
