/*
** ===============================================
** USER NAME: garlic(QQ:3173413)
** FILE NAME: auth.go
** DATE TIME: 2017-07-17 23:58:01
** ===============================================
 */

package alipay

func PublicAppAuthorize(app_id, scope, redirect_uri, state) string {
	uri := ""

	if config.SandBoxEnable {
		uri = "https://openauth.alipaydev.com/"
	} else {
		uri = "https://openauth.alipay.com/"
	}
	uri += "oauth2/publicAppAuthorize.htm"
	uri += "?app_id=" + app_id
	uri += "&scope=" + scope
	uri += "&redirect_uri=" + redirect_uri
	if len(state) > 0 {
		uri += "&state=" + state
	}
}
