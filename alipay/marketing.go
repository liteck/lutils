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

func (k *koubei_marketing_data_dishdiagnosetype_batchquery) SetParams(m map[string]string) error {
	k.AlipayApi.setApiMethod("koubei.marketing.data.dishdiagnosetype.batchquery")
	k.AlipayApi.setApiName("菜品类型查询")
	k.BizContent = k.packageBizContent()

	return k.AlipayApi.SetParams(m)
}

func (k *koubei_marketing_data_dishdiagnosetype_batchquery) getApiMethod() string {
	return "koubei.marketing.data.dishdiagnosetype.batchquery"
}

func (k *koubei_marketing_data_dishdiagnosetype_batchquery) getApiName() string {
	return "菜品类型查询"
}

func (k *koubei_marketing_data_dishdiagnosetype_batchquery) packageBizContent() string {
	return ""
}

func init() {
	registerApi(new(koubei_marketing_data_dishdiagnosetype_batchquery))
}
