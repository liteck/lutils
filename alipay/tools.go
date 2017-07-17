/**
 工具类api

**/
package alipay

/**
换取授权访问令牌
alipay.system.oauth.token
换取授权访问令牌
*/
type alipay_system_oauth_token struct {
	AlipayApi
}

func (a *alipay_system_oauth_token) SetAppId(app_id string) error {
	a.setApiMethod(a.apiMethod())
	a.setApiName(a.apiName())
	return a.AlipayApi.SetAppId(app_id)
}

func (k *alipay_system_oauth_token) apiMethod() string {
	return "alipay.system.oauth.token"
}

func (k *alipay_system_oauth_token) apiName() string {
	return "换取授权访问令牌"
}

func (a *alipay_system_oauth_token) packageBizContent() string {
	return ""
}

func init() {
	registerApi(new(alipay_system_oauth_token))
}
