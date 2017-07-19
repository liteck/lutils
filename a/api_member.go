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

/**
支付宝钱包用户信息共享
alipay.user.userinfo.share
外部应用上架到支付宝钱包，当支付宝用户从钱包访问外部应用时，会跳转到外部应用并带上用户的授权码。
外部应用用授权码调用授权令牌交换API（alipay.system.oauth.token）可得到授权令牌。
用授权令牌调用此接口得到支付宝会员相关信息。
特别说明：此接口的不需要授权是指不需外部应用主动引导用户授权，支付宝钱包会在引导用户授权后， 带上授权码再跳转到外部应用。
*/
type alipay_user_userinfo_share struct {
	AlipayApi
}

func (a *alipay_user_userinfo_share) apiMethod() string {
	return "alipay.user.userinfo.share"
}

func (a *alipay_user_userinfo_share) apiName() string {
	return "支付宝钱包用户信息共享"
}

type Biz_alipay_user_userinfo_share struct {
	AuthToken string `json:"auth_token,omitempty"`
}

func (b Biz_alipay_user_userinfo_share) valid() error {
	if len(b.AuthToken) == 0 {
		return errors.New("auth_token" + CAN_NOT_NIL)
	}

	return nil
}

type Resp_alipay_user_userinfo_share struct {
	Response
	UserId                string `json:"user_id,omitempty"`
	Avater                string `json:"avatar,omitempty"`
	UserTypeValue         string `json:"user_type_value,omitempty"`
	UserStatus            string `json:"user_status,omitempty"`
	FirmName              string `json:"firm_name,omitempty"`
	RealName              string `json:"real_name,omitempty"`
	Email                 string `json:"email,omitempty"`
	CertNo                string `json:"cert_no,omitempty"`
	Phone                 string `json:"phone,omitempty"`
	Mobile                string `json:"mobile,omitempty"`
	IsCertified           string `json:"is_certified,omitempty"`
	IsBankAuth            string `json:"is_bank_auth,omitempty"`
	IsIdAuth              string `json:"is_id_auth,omitempty"`
	IsMobileAuth          string `json:"is_mobile_auth,omitempty"`
	IsLiceneAuth          string `json:"is_licence_auth,omitempty"`
	CertTypeValue         string `json:"cert_type_value,omitempty"`
	DiliverPhone          string `json:"deliver_phone,omitempty"`
	DiliverMobile         string `json:"deliver_mobile,omitempty"`
	DiliverFullname       string `json:"deliver_fullname,omitempty"`
	Province              string `json:"province,omitempty"`
	City                  string `json:"city,omitempty"`
	Area                  string `json:"area,omitempty"`
	Address               string `json:"address,omitempty"`
	Zip                   string `json:"zip,omitempty"`
	DeliverProvince       string `json:"deliver_province,omitempty"`
	DeliverCity           string `json:"deliver_city,omitempty"`
	DeliverArea           string `json:"deliver_area,omitempty"`
	DefaultDeliverAddress string `json:"default_deliver_address,omitempty"`
	AddressCode           string `json:"address_code,omitempty"`
	NickName              string `json:"nick_name,omitempty"`
	IsStudentCertified    string `json:"is_student_certified,omitempty"`
	IsCertifyGradeA       string `json:"is_certify_grade_a,omitempty"`
	AlipayUserId          string `json:"alipay_user_id,omitempty"`
	Birthday              string `json:"birthday,omitempty"`
	FamilyName            string `json:"family_name,omitempty"`
	Gender                string `json:"gender,omitempty"`
	ReducedBirthday       string `json:"reduced_birthday,omitempty"`
	IsBalanceFrozen       string `json:"is_balance_frozen,omitempty"`
	BalanceFreezeType     string `json:"balance_freeze_type,omitempty"`
}

func init() {
	registerApi(new(alipay_user_userinfo_share))
	registerApi(new(alipay_user_info_share))
}
