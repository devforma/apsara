package core

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

type Client struct {
	logger     Logger
	cfg        *Config
	httpClient *http.Client
}

// NewClient 创建ASAPI客户端
func NewClient(cfg *Config, logger Logger) *Client {
	dialer := net.Dialer{
		Timeout: cfg.ConnectionTimeout,
	}
	return &Client{
		logger: logger,
		cfg:    cfg,
		httpClient: &http.Client{
			Transport: &http.Transport{
				DialContext:  dialer.DialContext,
				MaxIdleConns: cfg.MaxIdleConns,
			},
		},
	}
}

// buildRPCRequest 构造RPC请求类型
func (c *Client) buildRPCRequest(r Request) *http.Request {
	timestamp := getTimestamp()
	nonce := getNonce()
	method := r.GetMethod()

	// 完善请求参数
	queries := r.GetQueries()
	queries["Timestamp"] = timestamp
	queries["SignatureNonce"] = nonce
	queries["Format"] = "JSON"
	queries["SignatureMethod"] = "HMAC-SHA1"
	queries["SignatureVersion"] = "1.0"
	queries["AccessKeyId"] = c.cfg.AccessKeyID
	if queries["RegionId"] == "" {
		queries["RegionId"] = c.cfg.RegionID
	}

	// 构造query string
	queries = sortMap(queries)
	queryString := getQueryString(queries)

	// 签名
	stringToSign := getRpcStringToSign(method, queryString)
	queries["Signature"] = getRPCSignature(stringToSign, c.cfg.AccessKeySecret)

	// 构造完整url
	url := fmt.Sprintf("%s?%s", c.cfg.AsapiEndpoint, queryString)
	req, _ := http.NewRequest(method, url, nil)

	// 设置请求头
	headers := r.GetHeaders()
	headers["user-agent"] = c.cfg.UserAgent
	headers["accept"] = "application/json"
	headers["x-ascm-product-ak"] = c.cfg.AccessKeyID
	headers["x-acs-date"] = timestamp
	headers["x-acs-caller-sdk-source"] = c.cfg.UserAgent
	headers["x-acs-signature-nonce"] = nonce
	headers["x-acs-signature-method"] = "HMAC-SHA1"
	headers["x-acs-signature-version"] = "1.0"
	if headers["x-acs-regionid"] == "" {
		headers["x-acs-regionid"] = c.cfg.RegionID
	}

	// ASO请求，需要额外参数
	pathname := r.GetPathname()
	if pathname != "" {
		headers["Date"] = getDateGMTString()

		asoStringToSign := getASOStringToSign(method, headers["Date"], pathname)
		headers["Authorization"] = getASOSignature(asoStringToSign, c.cfg.AasSecret)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req
}

// GET\n\napplication/json\nWed, 21 Sep 2022 10:49:57 GMT\n/aso/v3/physicalInfo/getMachineList

// buildROARequest 构造ROA请求类型
func (c *Client) buildROARequest(r Request) *http.Request {
	timestamp := getTimestamp()
	nonce := getNonce()
	method := r.GetMethod()
	pathname := r.GetPathname()
	body := r.GetBody()
	query := r.GetQueries()

	// 完善请求头
	headers := r.GetHeaders()
	headers["user-agent"] = c.cfg.UserAgent
	headers["date"] = getDateGMTString()
	headers["accept"] = "application/json"
	headers["x-ascm-product-ak"] = c.cfg.AccessKeyID
	headers["x-acs-date"] = timestamp
	headers["x-acs-caller-sdk-source"] = c.cfg.UserAgent
	headers["x-acs-signature-nonce"] = nonce
	headers["x-acs-signature-method"] = "HMAC-SHA1"
	headers["x-acs-signature-version"] = "1.0"
	if headers["x-acs-regionid"] == "" {
		headers["x-acs-regionid"] = c.cfg.RegionID
	}
	if body != nil {
		headers["content-type"] = "application/json; charset=utf-8"
	}

	// 构造签名
	stringToSign := getROAStringToSign(method, headers, query, pathname)
	headers["authorization"] = fmt.Sprintf("acs %s:%s", c.cfg.AccessKeyID, getROASignature(stringToSign, c.cfg.AccessKeySecret))

	url := fmt.Sprintf("%s%s", c.cfg.AsapiEndpoint, pathname)

	// 构造query string
	queryString := getQueryString(query)
	if len(queryString) > 0 {
		if strings.Contains(url, "?") {
			url = fmt.Sprintf("%s&%s", url, queryString)
		} else {
			url = fmt.Sprintf("%s?%s", url, queryString)
		}
	}

	// 构造请求
	req, _ := http.NewRequest(method, url, bytes.NewReader(body))

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return nil
}

// DoRequest 发起请求
func (c *Client) DoRequest(request Request, response Response) error {
	var (
		req        *http.Request
		statusCode int
		err        error
		content    []byte
	)

	if request.GetStyle() == RequestStyleRPC {
		req = c.buildRPCRequest(request)
	} else if request.GetStyle() == RequestStyleROA {
		req = c.buildROARequest(request)
	} else {
		return errors.New("request style is not supported")
	}

	content, statusCode, err = c.request(req)
	if err != nil {
		if c.cfg.EnableLog {
			c.logger.Error("DoRequest error: %v", err)
		}
		return err
	}

	response.SetHeaders(nil)
	response.SetBody(content)
	response.SetStatusCode(statusCode)

	// 请求中的参数传递到响应中
	if cached := request.GetCachedParams(); cached != nil {
		response.SetCachedRequestParams(cached)
	}

	return err
}

func (c *Client) request(req *http.Request) ([]byte, int, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return content, resp.StatusCode, nil
}
