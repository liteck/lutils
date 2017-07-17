/**
 工具类api

**/
package alipay

/**
换取授权访问令牌
alipay.system.oauth.token
换取授权访问令牌
*/
type AlipaySystemOauthToken struct {
	AlipayApi
}

func (a *AlipaySystemOauthToken) SetAppId(app_id string) error {
	a.setApiMethod(a.apiMethod())
	a.setApiName(a.apiName())
	return a.AlipayApi.SetAppId(app_id)
}

func (k *AlipaySystemOauthToken) apiMethod() string {
	return "alipay.system.oauth.token"
}

func (k *AlipaySystemOauthToken) apiName() string {
	return "换取授权访问令牌"
}

func (a *AlipaySystemOauthToken) packageBizContent() string {
	return ""
}

func init() {
	registerApi(new(AlipaySystemOauthToken))
}
