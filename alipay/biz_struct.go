/*
** ===============================================
** USER NAME: garlic(QQ:3173413)
** FILE NAME: biz_struct.go
** DATE TIME: 2017-07-18 15:55:27
** 支付宝所有 bizcontent 的数据结构,以 Biz_method 命名
** ===============================================
 */

package alipay

import "errors"

const (
	CAN_NOT_NIL  = "不能为空"
	FORAMT_ERROR = "格式错误"
)

type BizInterface interface {
	valid() error
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