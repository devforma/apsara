package core

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/devforma/apsara/util"
	"github.com/google/uuid"
)

var gmtTimeLoc = time.FixedZone("GMT", 0)

// GMT格式时间字符串
const DateTimeFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

// getTimestamp 请求时间戳
func getTimestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05Z")
}

// getDateGMTString 请求头date取值
func getDateGMTString() string {
	return time.Now().In(gmtTimeLoc).Format(DateTimeFormat)
}

// getNonce 签名唯一随机数，防止网络重放攻击
func getNonce() string {
	uuid, _ := uuid.New().MarshalBinary()
	return hex.EncodeToString(uuid)
}

// getROASignature 生成ROA请求的签名
func getROASignature(stringToSign string, secret string) string {
	return base64.StdEncoding.EncodeToString(shaHmac1(stringToSign, secret))
}

// getROAStringToSign 构造ROA请求待签名字符串
func getROAStringToSign(method string, headers map[string]string, queries map[string]string, pathname string) string {
	var resourceStringBuilder strings.Builder
	resourceStringBuilder.WriteString(pathname)

	// query处理
	if len(queries) > 0 {
		resourceStringBuilder.WriteString("?")

		queryKeys := sortMap(queries)
		for _, key := range queryKeys {
			resourceStringBuilder.WriteString(key)
			if queries[key] != "" {
				resourceStringBuilder.WriteString("=")
				resourceStringBuilder.WriteString(queries[key])
			}
			resourceStringBuilder.WriteString("&")
		}
	}
	resourceString := resourceStringBuilder.String()
	resourceString = strings.TrimRight(resourceString, "&")

	// header处理
	var headerStringBuilder strings.Builder
	headerKeys := sortMap(headers)
	for _, key := range headerKeys {
		// 过滤掉非x-acs-开头的header
		if !strings.HasPrefix(key, "x-acs-") {
			continue
		}

		headerStringBuilder.WriteString(key)
		headerStringBuilder.WriteString(":")
		headerStringBuilder.WriteString(headers[key])
		headerStringBuilder.WriteString("\n")
	}
	headerString := headerStringBuilder.String()

	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s%s", method, headers["accept"], headers["content-md5"], headers["content-type"], headers["x-acs-date"], headerString, resourceString)
}

// getASOSignature 生成RPC请求中的ASO签名
func getASOSignature(stringToSign string, secret string) string {
	whole := base64.StdEncoding.EncodeToString(shaHmac1(stringToSign, secret))
	return whole[:len(whole)-2]
}

// getASOStringToSign 构造RPC请求中的ASO待签名字符串
func getASOStringToSign(method string, date string, pathname string) string {
	//method + "\n" + md5Body + "\n" + contentType + "\n" + date + "\n" + resource
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s", method, "", "", date, pathname)
}

// sortMap对map内元素按照字符序递增排序
func sortMap(orig map[string]string) []string {
	var keys []string

	for key := range orig {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// getRPCSignature 生成RPC请求的签名
func getRPCSignature(stringToSign string, secret string) string {
	return base64.StdEncoding.EncodeToString(shaHmac1(stringToSign, secret+"&"))
}

// getRpcStringToSign 构造RPC请求待签名字符串
func getRpcStringToSign(method string, queryString string) string {
	queryString = strings.Replace(queryString, "+", "%20", -1)
	queryString = strings.Replace(queryString, "*", "%2A", -1)
	queryString = strings.Replace(queryString, "%7E", "~", -1)
	queryString = url.QueryEscape(queryString)
	return method + "&%2F&" + queryString
}

// getQueryString 构造query string
func getQueryString(queries map[string]string) string {
	urlEncoder := url.Values{}

	queryKeys := sortMap(queries)
	for _, key := range queryKeys {
		urlEncoder.Add(key, queries[key])
	}

	return urlEncoder.Encode()
}

// shaHmac1 shahmac1签名
func shaHmac1(source, secret string) []byte {
	hmac := hmac.New(sha1.New, []byte(secret))
	hmac.Write(util.StringToBytes(source))
	return hmac.Sum(nil)
}
