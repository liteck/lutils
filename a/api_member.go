/**
 会员api

**/
package a

import (
	"errors"
)

/**
支付宝会员授权信息查询接口
alipay.user.info.share
配合支付宝会员授权接口，根据授权token，查询授权信息。
*/
type alipay_user_info_share struct {
	AlipayApi
}

func (a *alipay_user_info_share) apiMethod() string {
	return "alipay.user.info.share"
}

func (a *alipay_user_info_share) apiName() string {
	return "支付宝会员授权信息查询接口"
}

type Biz_alipay_user_info_share struct {
	AuthToken string `json:"auth_token,omitempty"`
}

func (b Biz_alipay_user_info_share) valid() error {
	if len(b.AuthToken) == 0 {
		return errors.New("auth_token" + CAN_NOT_NIL)
	}

	return nil
}

type Resp_alipay_user_info_share struct {
	Response
	UserId             string `json:"user_id,omitempty"`
	Avater             string `json:"avatar,omitempty"`
	UserType           string `json:"user_type,omitempty"`
	UserStatus         string `json:"user_status,omitempty"`
	IsCertified        string `json:"is_certified,omitempty"`
	Province           string `json:"province,omitempty"`
	City               string `json:"city,omitempty"`
	NickName           string `json:"nick_name,omitempty"`
	IsStudentCertified string `json:"is_student_certified,omitempty"`
	Gender             string `json:"gender,omitempty"`
}

func init() {
	registerApi(new(alipay_user_info_share))
}
