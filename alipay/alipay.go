package alipay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"lutils/httplib"
	"lutils/logs"
	"lutils/mahonia"
	"reflect"
	"sort"
	"strings"
)

//支付宝api接口的抽象方法
type alipayApiInterface interface {
	Run() error
	packageBizContent() string
	getApiMethod() string
	getApiName() string
	SetParams(m map[string]string) error
}

type apis map[string]alipayApiInterface

func (a apis) put(k string, v alipayApiInterface) {
	a[k] = v
}

func (a apis) get(k string) (alipayApiInterface, bool) {
	v, ok := a[k]
	return v, ok
}

var apiLst apis

func registerApi(v alipayApiInterface) {
	apiLst.put(v.getApiMethod(), v)
}

func GetSupportApis() []string {
	ret := []string{}
	for _, v := range apiLst {
		r := "\n" + v.getApiName() + ":" + v.getApiMethod()
		ret = append(ret, r)
	}
	return ret
}

func GetApi(k string) (alipayApiInterface, bool) {
	return apiLst.get(k)
}

type AlipayApi struct {
	params     requestParams
	MethodName string
	BizContent string
}

/**
 * 支付宝的数据返回应该是GBK的,,在golang中显示乱码.需要转换
 * @param  {[type]} utf string)       (gbk string [description]
 * @return {[type]}     [description]
 */
func (a *AlipayApi) convertUTF2GBK(utf string) (gbk string) {
	gbk_enc := mahonia.NewEncoder("GBK")
	return gbk_enc.ConvertString(utf)
}

func (a *AlipayApi) convertGBK2UTF(gbk string) (utf string) {
	enc := mahonia.NewDecoder("GBK")
	return enc.ConvertString(gbk)
}

func (a *AlipayApi) buildRequestParams() map[string]interface{} {
	biz_content := a.packageBizContent()
	a.params.BizContent = biz_content
	a.params.Method = a.getApiMethod()

	t := reflect.TypeOf(a.params)
	v := reflect.ValueOf(a.params)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Name
		value := v.Field(i).Interface()
		tag := t.Field(i).Tag.Get("ali")
		if tag != "" {
			if strings.Contains(tag, ",") {
				ps := strings.Split(tag, ",")
				key = ps[0]
			} else {
				key = tag
			}
		}
		data[key] = value
	}
	return data
}

func (a *AlipayApi) mTos(m map[string]interface{}) string {
	//对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range m {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	//对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", m[k])
		if value != "" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}
	if len(signStrings) == 0 {
		return ""
	} else {
		signStrings = signStrings[:len(signStrings)-1]
	}
	return signStrings
}

func (a *AlipayApi) sign(c string) (sign string, err error) {
	//签名
	s := getSecret(a.params.AppId)
	block, _ := pem.Decode(s.PrivRSA)
	if block == nil {
		return "", errors.New("privateKey error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	_t := crypto.SHA1.New()
	_t.Write([]byte(c))
	digest := _t.Sum(nil)
	rsa_data, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA1, digest)
	if err != nil {
		return "", err
	}
	result := string(rsa_data)
	result = base64.StdEncoding.EncodeToString(rsa_data)
	return result, nil
}

func (a *AlipayApi) verifySign(in string, origin_sign string) bool {
	return true
}

func (a *AlipayApi) request(m map[string]interface{}) (string, error) {
	url_link := "https://openapi.alipay.com/gateway.do"
	if conf.SandBoxEnable {
		url_link = "https://openapi.alipaydev.com/gateway.do"
	}
	http_request := httplib.Post(url_link)
	tmp_string := ""
	for k, _ := range m {
		value := fmt.Sprintf("%v", m[k])
		if value != "" {
			http_request.Param(k, value)
			tmp_string = tmp_string + k + "=" + value + "\t"
		}
	}

	logs.DEBUG(fmt.Sprintf("==[请求参数]==[%s]", tmp_string))
	var string_result string
	if v, err := http_request.String(); err != nil {
		return "", err
	} else {
		string_result = v

	}
	return string_result, nil
}

func (a *AlipayApi) getApiMethod() string {
	if len(a.params.Method) > 0 {
		return a.params.Method
	} else {
		return ErrMethodNotSupport.Error()
	}
}

func (a *AlipayApi) setApiMethod(method string) {
	a.params.Method = method
}

func (a *AlipayApi) packageBizContent() string {
	return ""
}

func (a *AlipayApi) getApiName() string {
	if len(a.MethodName) > 0 {
		return a.MethodName
	} else {
		return ErrMethodNameNil.Error()
	}
}

func (a *AlipayApi) setApiName(name string) {
	a.MethodName = name
}

func (a *AlipayApi) SetParams(m map[string]string) error {
	if v, ok := m["app_id"]; ok {
		a.params.AppId = v
	}
	if len(a.params.AppId) == 0 {
		return ErrAppIdNil
	}

	if _, ok := secretLst[a.params.AppId]; !ok {
		return ErrSecretNil
	}
	//有些接口需要
	if v, ok := m["auth_token"]; ok {
		a.params.AuthToken = v
	}

	if v, ok := m["code"]; ok {
		a.params.Code = v
	}

	if v, ok := m["grant_type"]; ok {
		a.params.GrantType = v
	}

	return nil
}

func (a *AlipayApi) Run() error {
	logs.DEBUG("=====================ALIPAY REQUEST START=====================")
	logs.DEBUG(fmt.Sprintf("==[沙盒模式]==[%v]", conf.SandBoxEnable))
	logs.DEBUG(fmt.Sprintf("==[调用方法]==[%s]:[%s]", a.MethodName, a.getApiMethod()))

	if err := a.params.valid(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	m := a.buildRequestParams()
	//做请求参数的签名
	__sign := ""
	tobe_sign := a.mTos(m)
	logs.DEBUG(fmt.Sprintf("==[准备签名]==[%s]", tobe_sign))
	if v, err := a.sign(tobe_sign); err != nil {
		return err
	} else if len(v) == 0 {
		return ErrSign
	} else {
		__sign = v
	}
	logs.DEBUG(fmt.Sprintf("==[签名结果]==[%s]", __sign))
	m["sign"] = __sign
	//准备请求
	result_string := ""
	if v, err := a.request(m); err != nil {
		return err
	} else {
		result_string = v
		logs.DEBUG(fmt.Sprintf("==[响应结果]==[GBK 编码]:[%s]", result_string))
	}
	//解析结果
	resp_map := map[string]interface{}{}
	if err := json.Unmarshal([]byte(result_string), &resp_map); err != nil {
		return err
	}
	//看看有没有sign
	if v, ok := resp_map["sign"].(string); ok {
		//有则校验签名
		if pass := a.verifySign(result_string, v); !pass {
			return ErrVerifySign
		}
	}
	//转码
	result_string = a.convertGBK2UTF(result_string)
	logs.DEBUG(fmt.Sprintf("==[响应结果]==[UTF 编码]:[%s]", result_string))
	//验证签名
	logs.DEBUG("=====================ALIPAY REQUEST END=====================")
	return nil
}

func init() {
	apiLst = apis{}
}
