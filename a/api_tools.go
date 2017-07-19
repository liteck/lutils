/**
 工具类api

**/
package a

import (
	"errors"
)

//=========================
/**
换取授权访问令牌
alipay.system.oauth.token
换取授权访问令牌
*/
type alipay_system_oauth_token struct {
	AlipayApi
}

func (a *alipay_system_oauth_token) apiMethod() string {
	return "alipay.system.oauth.token"
}

func (a *alipay_system_oauth_token) apiName() string {
	return "换取授权访问令牌"
}

type Biz_alipay_system_oauth_token struct {
	GrantType    string `json:"grant_type,omitempty"`
	Code         string `json:"code,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (b Biz_alipay_system_oauth_token) valid() error {
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

type Resp_alipay_system_oauth_token struct {
	Response
	AlipayUserId string  `json:"alipay_user_id,omitempty"`
	UserId       string  `json:"user_id,omitempty"`
	AccessToken  string  `json:"access_token,omitempty"`
	RefreshToken string  `json:"refresh_token,omitempty"`
	ExpiresIn    float64 `json:"expires_in,omitempty"`
	ReExpiresIn  float64 `json:"re_expires_in,omitempty"`
}

//=========================
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

type Biz_alipay_open_auth_token_app struct {
	GrantType    string `json:"grant_type,omitempty"`
	Code         string `json:"code,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (b Biz_alipay_open_auth_token_app) valid() error {
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

type Resp_alipay_open_auth_token_app struct {
	Response
	UserId          string  `json:"user_id,omitempty"`
	AuthAppId       string  `json:"auth_app_id,omitempty"`
	AppAuthToken    string  `json:"app_auth_token,omitempty"`
	AppRefreshToken string  `json:"app_refresh_token,omitempty"`
	ExpiresIn       float64 `json:"expires_in,omitempty"`
	ReExpiresIn     float64 `json:"re_expires_in,omitempty"`
}

func init() {
	registerApi(new(alipay_system_oauth_token))
	registerApi(new(alipay_open_auth_token_app))
}
