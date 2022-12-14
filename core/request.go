package core

import (
	"github.com/devforma/apsara/util"
	"github.com/tidwall/gjson"
)

type RequestStyle string

const (
	RequestStyleROA RequestStyle = "ROA"
	RequestStyleRPC RequestStyle = "RPC"
	RequestStyleBCC RequestStyle = "BCC"
)

type Request interface {
	GetHeaders() map[string]string
	GetQueries() map[string]string
	GetBody() []byte
	GetMethod() string
	GetPathname() string
	GetStyle() RequestStyle
	GetCachedParams() map[string]string
	SetAccessKey(string, string)
	GetAccessKey() (string, string)
}

type Response interface {
	SetStatusCode(code int)
	SetHeaders(headers map[string]string)
	SetBody(body []byte)
	SetCachedRequestParams(map[string]string)
}

type EmbededRequest struct {
	Product         string
	Version         string
	Action          string
	RegionID        string
	OrganizationID  string
	ResourceGroupID string
	InstanceID      string
	Style           RequestStyle
	Pathname        string
	BizQueries      map[string]string
	BizHeaders      map[string]string
	Body            string
	CachedParams    map[string]string
	AccessKeyID     string
	AccessKeySecret string
}

func (r *EmbededRequest) SetAccessKey(accessKeyID, accessKeySecret string) {
	r.AccessKeyID = accessKeyID
	r.AccessKeySecret = accessKeySecret
}

func (r *EmbededRequest) GetAccessKey() (string, string) {
	return r.AccessKeyID, r.AccessKeySecret
}

func (r *EmbededRequest) GetHeaders() map[string]string {
	headers := map[string]string{
		"x-acs-regionid":         r.RegionID,
		"x-acs-organizationid":   r.OrganizationID,
		"x-acs-resourcegroupid":  r.ResourceGroupID,
		"x-acs-instanceid":       r.InstanceID,
		"x-acs-version":          r.Version,
		"x-acs-action":           r.Action,
		"x-ascm-product-name":    r.Product,
		"x-ascm-product-version": r.Version,
	}

	for k, v := range r.BizHeaders {
		headers[k] = v
	}

	return headers
}

func (r *EmbededRequest) GetQueries() map[string]string {
	query := make(map[string]string)
	if r.Style == RequestStyleRPC {
		query["Product"] = r.Product
		query["Version"] = r.Version
		query["Action"] = r.Action
		query["RegionId"] = r.RegionID
	}

	for k, v := range r.BizQueries {
		query[k] = v
	}

	return query
}

func (r *EmbededRequest) GetBody() []byte {
	return util.StringToBytes(r.Body)
}

func (r *EmbededRequest) GetMethod() string {
	return "GET"
}

func (r *EmbededRequest) GetPathname() string {
	return r.Pathname
}

func (r *EmbededRequest) GetStyle() RequestStyle {
	return r.Style
}

func (r *EmbededRequest) GetCachedParams() map[string]string {
	return r.CachedParams
}

type EmbededResponse struct {
	StatusCode          int
	Headers             map[string]string
	Body                gjson.Result
	CachedRequestParams map[string]string
}

func (r *EmbededResponse) SetStatusCode(code int) {
	r.StatusCode = code
}

func (r *EmbededResponse) SetHeaders(headers map[string]string) {
	r.Headers = headers
}

func (r *EmbededResponse) SetBody(body []byte) {
	r.Body = gjson.ParseBytes(body)
}

func (r *EmbededResponse) SetCachedRequestParams(params map[string]string) {
	r.CachedRequestParams = params
}

func (r *EmbededResponse) IsSuccess() bool {
	return r.Body.Get("asapiSuccess").Bool()
}
