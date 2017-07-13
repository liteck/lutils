/**
营销api

**/
package alipay

import "fmt"

/**
菜品类型查询
koubei.marketing.data.dishdiagnosetype.batchquery
获取菜品类型，可以查询类型与对应的类型说明。
与API：koubei.marketing.data.dishdiagnose.batchquery配合使用，先查询出支持的类型，然后根据类型去查询对应的数据
*/
type koubei_marketing_data_dishdiagnosetype_batchquery struct {
	AlipayApi
}

func (k *koubei_marketing_data_dishdiagnosetype_batchquery) Init(app_id string) {
	k.params.AppId = app_id
	k.params.Method = "koubei.marketing.data.dishdiagnosetype.batchquery"
	k.params.MethodName = "菜品类型查询"
}

func (k *koubei_marketing_data_dishdiagnosetype_batchquery) packageBizContent() string {
	fmt.Println("koubei_marketing_data_dishdiagnosetype_batchquery PackageBizContent")
	return ""
}

func init() {
	fmt.Println("register koubei_marketing_data_dishdiagnosetype_batchquery")
	registerApi(new(koubei_marketing_data_dishdiagnosetype_batchquery))
}
