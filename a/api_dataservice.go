/*
** ===============================================
** USER NAME: garlic(QQ:3173413)
** FILE NAME: api_dataservice.go
** DATE TIME: 2017-08-22 13:40:49
** 开放平台财务 API
** ===============================================
 */

package a

import (
	"errors"
	"time"
)

/**
查询对账单下载地址
alipay.data.dataservice.bill.downloadurl.query
为方便商户快速查账，支持商户通过本接口获取商户离线账单下载地址
*/
type alipay_data_dataservice_bill_downloadurl_query struct {
	AlipayApi
}

func (a *alipay_data_dataservice_bill_downloadurl_query) apiMethod() string {
	return "alipay.data.dataservice.bill.downloadurl.query"
}

func (a *alipay_data_dataservice_bill_downloadurl_query) apiName() string {
	return "查询对账单下载地址"
}

type Biz_alipay_data_dataservice_bill_downloadurl_query struct {
	//账单类型，商户通过接口或商户经开放平台授权后其所属服务商通过接口可以获取以下账单类型：
	//trade、signcustomer；
	//trade 指商户基于支付宝交易收单的业务账单；
	//signcustomer 是指基于商户支付宝余额收入及支出等资金变动的帐务账单
	BillType string `json:"bill_type,omitempty"`
	//账单时间：日账单格式为yyyy-MM-dd，月账单格式为yyyy-MM。
	BillDate string `json:"bill_date,omitempty"`
}

func (b Biz_alipay_data_dataservice_bill_downloadurl_query) valid() error {
	if len(b.BillType) == 0 {
		return errors.New("bill_type" + CAN_NOT_NIL)
	}
	if b.BillType != "trade" && b.BillType != "signcustomer" {
		return errors.New("bill_type" + FORAMT_ERROR)
	}

	if len(b.BillDate) == 0 {
		return errors.New("bill_date" + CAN_NOT_NIL)
	} else if len(b.BillDate) == 10 {
		//日账单
		if _, err := time.Parse("2006-01-02", b.BillDate); err != nil {
			return errors.New("bill_date" + FORAMT_ERROR)
		}
	} else if len(b.BillDate) == 7 {
		//月账单
		if _, err := time.Parse("2006-01", b.BillDate); err != nil {
			return errors.New("bill_date" + FORAMT_ERROR)
		}
	} else {
		return errors.New("bill_date" + FORAMT_ERROR)
	}

	return nil
}

type Resp_alipay_data_dataservice_bill_downloadurl_query struct {
	Response
	BillDownloadUrl string `json:"bill_download_url,omitempty"` //账单下载地址链接，获取连接后30秒后未下载，链接地址失效。
}

func init() {
	registerApi(new(alipay_data_dataservice_bill_downloadurl_query))
}
