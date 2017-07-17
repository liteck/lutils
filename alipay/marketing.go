/**
营销api

**/
package alipay

/**
菜品类型查询
koubei.marketing.data.dishdiagnosetype.batchquery
获取菜品类型，可以查询类型与对应的类型说明。
与API：koubei.marketing.data.dishdiagnose.batchquery配合使用，先查询出支持的类型，然后根据类型去查询对应的数据
*/
type koubei_marketing_data_dishdiagnosetype_batchquery struct {
	AlipayApi
}

func (k *koubei_marketing_data_dishdiagnosetype_batchquery) SetAppId(app_id string) {
	k.params.AppId = app_id
	k.Method = k.getApiMethod()
	k.MethodName = k.getApiMethodName()
	k.BizContent = k.packageBizContent()
	k.init(app_id)
}

func (k *koubei_marketing_data_dishdiagnosetype_batchquery) packageBizContent() string {
	return ""
}

func (k *koubei_marketing_data_dishdiagnosetype_batchquery) getApiMethod() string {
	return "koubei.marketing.data.dishdiagnosetype.batchquery"
}

func (k *koubei_marketing_data_dishdiagnosetype_batchquery) getApiMethodName() string {
	return "菜品类型查询"
}

func init() {
	registerApi(new(koubei_marketing_data_dishdiagnosetype_batchquery))
}
