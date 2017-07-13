/**
公共配置等等
**/

package alipay

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrMethodNotSupport = errors.New("METHOD NOT SUPPORT")
	ErrMethodNameNil    = errors.New("METHOD NAME NIL")
	ErrAppIdNil         = errors.New("APPID NIL")
	ErrSecretNil        = errors.New("SECRET NIL")
	ErrSign             = errors.New("SIGN ERROR")
)

//公共请求参数
type requestParams struct {
	AppId        string `ali:"app_id"`
	Method       string `ali:"method"`
	Format       string `ali:"format"`
	Charset      string `ali:"charset"`
	SignType     string `ali:"sign_type"`
	Sign         string `ali:"sign"`
	TimeStamp    string `ali:"timestamp"`
	Version      string `ali:"version"`
	AppAuthToken string `ali:"app_auth_token"`
	BizContent   string `ali:"biz_content"`
}

func (params *requestParams) valid() error {
	if len(params.AppId) == 0 {
		return ErrAppIdNil
	}

	// if len(params.Method) == 0 {
	// 	return errors.New("method 不能为空")
	// }

	if len(params.Format) == 0 {
		params.Format = "JSON"
	}

	if len(params.Format) > 0 && params.Format != "JSON" {
		return errors.New("format 仅支持JSON")
	}

	if len(params.Charset) == 0 {
		params.Charset = "GBK"
	}

	if len(params.Charset) > 0 && strings.ToUpper(params.Charset) != "GBK" {
		return errors.New("charset 目前仅支持GBK")
	}

	if len(params.SignType) == 0 {
		params.SignType = "RSA"
	}

	if len(params.SignType) > 0 && params.SignType != "RSA" {
		return errors.New("sign_type 仅支持RSA")
	}

	if len(params.TimeStamp) == 0 {
		params.TimeStamp = time.Now().Format("2006-01-02 15:04:05")
	}

	if len(params.TimeStamp) > 0 {
		if _, err := time.Parse("2006-01-02 15:04:05", params.TimeStamp); err != nil {
			return errors.New("timestamp 格式\"yyyy-MM-dd HH:mm:ss\"")
		}
	}

	if len(params.Version) == 0 {
		params.Version = "1.0"
	} else if params.Version != "1.0" {
		return errors.New("version 固定为：1.0")
	}

	return nil
}
