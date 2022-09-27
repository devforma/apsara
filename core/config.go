package core

import "time"

type Config struct {
	AsapiEndpoint   string // 天基 -> 报表 -> 服务注册变量 -> asapi.public.endpoint
	Protocol        string
	RegionID        string
	AccessKeyID     string // 运营控制台运营账号/组织的ak
	AccessKeySecret string // 运营控制台运营账号/组织的sk
	AasSecret       string // 天基 -> 报表 -> 服务注册变量 -> 搜索 "baseService-aas" -> 得到 "bid.accesssecret"

	ConnectionTimeout time.Duration
	ReadTimeout       time.Duration
	MaxIdleConns      int
	UserAgent         string
	EnableLog         bool
}
