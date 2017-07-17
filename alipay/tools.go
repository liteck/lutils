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

func (a *alipay_system_oauth_token) SetParams(m map[string]string) error {
	a.AlipayApi.setApiMethod("alipay.system.oauth.token")
	a.AlipayApi.setApiName("换取授权访问令牌")
	a.BizContent = a.packageBizContent()

	return a.AlipayApi.SetParams(m)
}

func (a *alipay_system_oauth_token) packageBizContent() string {
	return ""
}

func init() {
	registerApi(new(alipay_system_oauth_token))
}
