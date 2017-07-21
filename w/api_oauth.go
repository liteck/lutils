/*
** ===============================================
** USER NAME: garlic(QQ:3173413)
** FILE NAME: api_auth.go
** DATE TIME: 2017-07-21 09:09:23
** ===============================================
 */

package w

import (
	"errors"
	"lutils/logs"
)

/**
授权参考文档
https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=4_4
微信这里授权方式比较多样.针对各种场景下.微信做了独立的区分...
1. 网页授权 开放平台用户授权  https://open.weixin.qq.com/connect
2.
**/

//开放平台的网页授权模式.获取授权链接
func OpenWebAuth(app_id, scope, redirect_uri, state string) string {
	uri := "https://open.weixin.qq.com/connect/oauth2/authorize"
	uri += "?appid=" + app_id
	uri += "&scope=" + scope
	uri += "&redirect_uri=" + redirect_uri
	if len(state) > 0 {
		uri += "&state=" + state
	}
	uri += "&response_type=code#wechat_redirect"
	logs.DEBUG(uri)
	return uri
}

//通过授权回调之后的 code 换取 access_token
type api_wechat_sns_oauth2_access_token struct {
	WechatApi
}

func (o *api_wechat_sns_oauth2_access_token) apiUrl() string {
	return "https://api.weixin.qq.com/sns/oauth2/access_token"
}

func (o *api_wechat_sns_oauth2_access_token) apiName() string {
	return "通过code获取access_token的接口"
}

func (o *api_wechat_sns_oauth2_access_token) apiMethod() string {
	return "GET"
}

type Req_api_wechat_sns_oauth2_access_token struct {
	Code      string `json:"code"`
	GrantType string `json:"grant_type"`
}

func (p Req_api_wechat_sns_oauth2_access_token) valid() error {
	if len(p.GrantType) == 0 {
		return errors.New("grant_type" + CAN_NOT_NIL)
	}
	if len(p.Code) == 0 {
		return errors.New("code" + CAN_NOT_NIL)
	}

	return nil
}

type Resp_api_wechat_sns_oauth2_access_token struct {
	Response
	AccessToken  string  `json:"access_token,omitempty"`
	ExpiresIn    float64 `json:"expires_in,omitempty"`
	RefreshToken string  `json:"refresh_token,omitempty"`
	OpenId       string  `json:"openid,omitempty"`
	Scope        string  `json:"scope,omitempty"`
}

func init() {
	registerApi(new(api_wechat_sns_oauth2_access_token))
}
