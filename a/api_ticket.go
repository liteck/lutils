/*
** ===============================================
** USER NAME: garlic(QQ:3173413)
** FILE NAME: api_ticket.go
** DATE TIME: 2017-08-04 10:08:00
** 支付宝线上商品核销类 api
** 内部接口.开放平台上暂时未开放
** ===============================================
 */
package a

import (
	"errors"
)

/**
口碑凭证码查询
koubei.trade.ticket.ticketcode.query
根据凭证码查询口碑凭证核销状态、核销明细、价格、有效期等信息; 仅允许查询本商户下的凭证
*/
type koubei_trade_ticket_ticketcode_query struct {
	AlipayApi
}

func (a *koubei_trade_ticket_ticketcode_query) apiMethod() string {
	return "koubei.trade.ticket.ticketcode.query"
}

func (a *koubei_trade_ticket_ticketcode_query) apiName() string {
	return "口碑凭证码查询"
}

type Biz_koubei_trade_ticket_ticketcode_query struct {
	TicketCode string `json:"ticket_code,omitempty"` //12位的券码，券码为纯数字 ，且唯一不重复
	ShopId     string `json:"shop_id,omitempty"`     //口碑门店id
}

func (b Biz_koubei_trade_ticket_ticketcode_query) valid() error {
	if len(b.TicketCode) == 0 {
		return errors.New("ticket_code" + CAN_NOT_NIL)
	}

	if len(b.ShopId) == 0 {
		return errors.New("shop_id" + CAN_NOT_NIL)
	}
	return nil
}

type Resp_koubei_trade_ticket_ticketcode_query struct {
	Response
	TicketCode       string `json:"ticket_code,omitempty"`   //12位的券码，券码为纯数字 ，且唯一不重复
	TicketStatus     string `json:"ticket_status,omitempty"` // 券状态
	TicketStatusDesc string `json:"ticket_status_desc,omitempty"`
	ItemName         string `json:"item_name,omitempty"`
	ItemId           string `json:"item_id,omitempty"`
	OriginalPrice    string `json:"original_price,omitempty"`
	CurrentPrice     string `json:"current_price,omitempty"`
	EffectDate       string `json:"effect_date,omitempty"`
	ExpireDate       string `json:"expire_date,omitempty"`
	SkuId            string `json:"sku_id,omitempty"`
	ItemMemo         string `json:"item_memo,omitempty"`
	TicketUseDetail  struct {
		UseDate             string `json:"use_date,omitempty"`
		UseShopId           string `json:"use_shop_id,omitempty"`
		UseShopName         string `json:"use_shop_name,omitempty"`
		BuyerPayAmount      string `json:"buyer_pay_amount,omitempty"`
		ReceiptAmount       string `json:"receipt_amount,omitempty"`
		DiscountAmount      string `json:"discount_amount,omitempty"`
		KoubeiSubsidyAmount string `json:"koubei_subsidy_amount,omitempty"`
		InvoiceAmount       string `json:"invoice_amount,omitempty"`
	} `json:"ticket_use_detail,omitempty"`
}

/**
口碑凭证码核销
koubei.trade.ticket.ticketcode.use
根据凭证码和门店 id 核销凭证
*/
type koubei_trade_ticket_ticketcode_use struct {
	AlipayApi
}

func (a *koubei_trade_ticket_ticketcode_use) apiMethod() string {
	return "koubei.trade.ticket.ticketcode.use"
}

func (a *koubei_trade_ticket_ticketcode_use) apiName() string {
	return "口碑凭证码核销"
}

type Biz_koubei_trade_ticket_ticketcode_use struct {
	RequestId  string `json:"request_id,omitempty"`  //开发者提供.外部请求号.支持英文和字母.保证唯一性
	TicketCode string `json:"ticket_code,omitempty"` //12位的券码，券码为纯数字 ，且唯一不重复
	ShopId     string `json:"shop_id,omitempty"`     //口碑门店id
}

func (b Biz_koubei_trade_ticket_ticketcode_use) valid() error {
	if len(b.TicketCode) == 0 {
		return errors.New("ticket_code" + CAN_NOT_NIL)
	}

	if len(b.ShopId) == 0 {
		return errors.New("shop_id" + CAN_NOT_NIL)
	}

	if len(b.RequestId) == 0 {
		return errors.New("request_id" + CAN_NOT_NIL)
	}
	return nil
}

type Resp_koubei_trade_ticket_ticketcode_use struct {
	Response
	RequestId           string `json:"request_id,omitempty"`
	TicketCode          string `json:"ticket_code,omitempty"` //12位的券码，券码为纯数字 ，且唯一不重复
	ItemName            string `json:"item_name,omitempty"`
	ItemId              string `json:"item_id,omitempty"`
	OriginalPrice       string `json:"original_price,omitempty"`
	CurrentPrice        string `json:"current_price,omitempty"`
	SkuId               string `json:"sku_id,omitempty"`
	UseDate             string `json:"use_date,omitempty"`
	UseShopId           string `json:"use_shop_id,omitempty"`
	UseShopName         string `json:"use_shop_name,omitempty"`
	ItemMemo            string `json:"item_memo,omitempty"`
	BuyerPayAmount      string `json:"buyer_pay_amount,omitempty"`
	ReceiptAmount       string `json:"receipt_amount,omitempty"`
	DiscountAmount      string `json:"discount_amount,omitempty"`
	KoubeiSubsidyAmount string `json:"koubei_subsidy_amount,omitempty"`
	InvoiceAmount       string `json:"invoice_amount,omitempty"`
}

func init() {
	registerApi(new(koubei_trade_ticket_ticketcode_query))
	registerApi(new(koubei_trade_ticket_ticketcode_use))
}
