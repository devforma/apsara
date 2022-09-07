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
func NewClient(cfg *Config, logger Logger) (*Client, error) {
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
	}, nil
}

// buildRPCRequest 构造RPC请求类型
func (c *Client) buildRPCRequest(r Request) *http.Request {
	timestamp := getTimestamp()
	nonce := getNonce()
	method := r.GetMethod()

	// 完善请求参数
	queries := r.GetQuery()
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
		headers["x-ascm-pass-through-mode"] = "true"
		queries["IsFormat"] = "false"
		headers["Date"] = getDateGMTString()

		asoStringToSign := getASOStringToSign(method, headers["Date"], pathname)
		headers["Authorization"] = getASOSignature(asoStringToSign, c.cfg.AasSecret)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req
}

// buildROARequest 构造ROA请求类型
func (c *Client) buildROARequest(r Request) *http.Request {
	timestamp := getTimestamp()
	nonce := getNonce()
	method := r.GetMethod()
	pathname := r.GetPathname()
	body := r.GetBody()
	query := r.GetQuery()

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
	var req *http.Request
	var resp *http.Response
	var err error
	var content []byte

	for i := request.GetRetryTimes(); i >= 0; i-- {
		if request.GetStyle() == RequestStyleRPC {
			req = c.buildRPCRequest(request)
		} else if request.GetStyle() == RequestStyleROA {
			req = c.buildROARequest(request)
		} else {
			return errors.New("request style is not supported")
		}

		resp, err = c.httpClient.Do(req)
		if err != nil {
			c.logger.Error("request failed: %v", err)
			continue
		}
		defer resp.Body.Close()

		content, err = io.ReadAll(resp.Body)
		if err != nil {
			c.logger.Error("read response failed: %v", err)
			continue
		}

		response.SetHeaders(nil)
		response.SetBody(content)
		response.SetStatusCode(resp.StatusCode)

		break
	}

	return err
}
