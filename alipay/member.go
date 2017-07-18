/**
 会员api

**/
package alipay

/**
支付宝会员授权信息查询接口
alipay.user.info.share
配合支付宝会员授权接口，根据授权token，查询授权信息。
*/
type AlipayUserInfoShare struct {
	AlipayApi
}

func (a *AlipayUserInfoShare) SetAppId(app_id string) error {
	a.setApiMethod(a.apiMethod())
	a.setApiName(a.apiName())
	return a.AlipayApi.SetAppId(app_id)
}

func (k *AlipayUserInfoShare) apiMethod() string {
	return "alipay.user.info.share"
}

func (k *AlipayUserInfoShare) apiName() string {
	return "支付宝会员授权信息查询接口"
}

func (a *AlipayUserInfoShare) SetAuthToken(auth_token string) {
	a.params.AuthToken = auth_token
}

func init() {
	registerApi(new(AlipayUserInfoShare))
}
