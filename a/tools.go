/**
 工具类api

**/
package a

import (
	"encoding/json"
	"errors"
)

/**
换取授权访问令牌
alipay.system.oauth.token
换取授权访问令牌
*/
type alipay_system_oauth_token struct {
	AlipayApi
}

// func (k *alipay_system_oauth_token) apiMethod() string {
// 	return "alipay.system.oauth.token"
// }

// func (k *alipay_system_oauth_token) apiName() string {
// 	return "换取授权访问令牌"
// }

// func (a *alipay_system_oauth_token) SetAuthCode(code string) {
// 	a.params.Code = code
// }

// func (a *alipay_system_oauth_token) SetGrantType(grant_type string) {
// 	a.params.GrantType = grant_type
// }

/**
换取授权访问令牌
alipay.open.auth.token.app
换取授权访问令牌
*/
type alipay_open_auth_token_app struct {
	AlipayApi
}

func (a *alipay_open_auth_token_app) apiMethod() string {
	return "alipay.open.auth.token.app"
}

func (a *alipay_open_auth_token_app) apiName() string {
	return "换取授权访问令牌"
}

func (a *alipay_open_auth_token_app) unmarshal(s string) interface{} {
	resp := RespAlipayOpenAuthTokenApp{}
	return resp
}

type BizAlipayOpenAuthTokenApp struct {
	GrantType    string `json:"grant_type,omitempty"`
	Code         string `json:"code,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (b BizAlipayOpenAuthTokenApp) valid() error {
	if len(b.GrantType) == 0 {
		return errors.New("grant_type" + CAN_NOT_NIL)
	}

	if b.GrantType != "authorization_code" && b.GrantType != "refresh_token" {
		return errors.New("grant_type" + FORAMT_ERROR)
	}

	if b.GrantType == "authorization_code" && len(b.Code) == 0 {
		return errors.New("code" + CAN_NOT_NIL)
	}

	if b.GrantType == "refresh_token" && len(b.RefreshToken) == 0 {
		return errors.New("refresh_token" + CAN_NOT_NIL)
	}
	return nil
}

func (b BizAlipayOpenAuthTokenApp) toString() (string, error) {
	if err := b.valid(); err != nil {
		return "", err
	}
	content := ""
	if v, err := json.Marshal(&b); err != nil {
		return "", err
	} else {
		content = string(v)
	}

	temp_map := map[string]interface{}{
		"biz": content,
	}

	if v, err := json.Marshal(&temp_map); err != nil {
		return "", err
	} else {
		content = string(v)
	}
	return content[8 : len(content)-2], nil
}

type RespAlipayOpenAuthTokenApp struct {
	Response
	UserId          string `json:"user_id,omitempty"`
	AuthAppId       string `json:"auth_app_id,omitempty"`
	AppAuthToken    string `json:"app_auth_token,omitempty"`
	AppRefreshToken string `json:"app_refresh_token,omitempty"`
	ExpiresIn       string `json:"expires_in,omitempty"`
	ReExpiresIn     string `json:"re_expires_in,omitempty"`
}

func (r *RespAlipayOpenAuthTokenApp) unmarshal(s string) error {
	if err := json.Unmarshal([]byte(s), r); err != nil {
		return err
	}
	return nil
}

func init() {
	// registerApi(new(alipay_system_oauth_token))
	registerApi(new(alipay_open_auth_token_app))
}
