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
type KoubeiMarketingDataDishdiagnosetypeBatchquery struct {
	AlipayApi
}

func (k *KoubeiMarketingDataDishdiagnosetypeBatchquery) SetAppId(app_id string) error {
	k.setApiMethod(k.apiMethod())
	k.setApiName(k.apiName())
	return k.AlipayApi.SetAppId(app_id)
}

func (k *KoubeiMarketingDataDishdiagnosetypeBatchquery) apiMethod() string {
	return "koubei.marketing.data.dishdiagnosetype.batchquery"
}

func (k *KoubeiMarketingDataDishdiagnosetypeBatchquery) apiName() string {
	return "菜品类型查询"
}

func init() {
	registerApi(new(KoubeiMarketingDataDishdiagnosetypeBatchquery))
}
