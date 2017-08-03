package a

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"lutils/httplib"
	"lutils/logs"
	"lutils/mahonia"
	"reflect"
	"sort"
	"strings"
	"time"
)

//公共请求参数
type requestParams struct {
	AppId        string `ali:"app_id"`
	Method       string `ali:"method"`
	Format       string `ali:"format"`
	Charset      string `ali:"charset"`
	SignType     string `ali:"sign_type"`
	Sign         string `ali:"sign"`
	TimeStamp    string `ali:"timestamp"`
	Version      string `ali:"version"`
	AppAuthToken string `ali:"app_auth_token"`
	BizContent   string `ali:"biz_content"`
	AuthToken    string `ali:"auth_token"`
	Code         string `ali:"code"`
	GrantType    string `ali:"grant_type"`
}

func (params *requestParams) valid() error {
	if len(params.AppId) == 0 {
		return ErrAppIdNil
	}

	if len(params.Method) == 0 {
		return errors.New("method 不能为空")
	}

	if len(params.Format) == 0 {
		params.Format = "JSON"
	}

	if len(params.Format) > 0 && params.Format != "JSON" {
		return errors.New("format 仅支持JSON")
	}

	if len(params.Charset) == 0 {
		params.Charset = "GBK"
	}

	if len(params.Charset) > 0 && strings.ToUpper(params.Charset) != "GBK" {
		return errors.New("charset 目前仅支持GBK")
	}

	if len(params.SignType) == 0 {
		params.SignType = "RSA"
	}

	if len(params.SignType) > 0 && params.SignType != "RSA" {
		return errors.New("sign_type 仅支持RSA")
	}

	if len(params.TimeStamp) == 0 {
		params.TimeStamp = time.Now().Format("2006-01-02 15:04:05")
	}

	if len(params.TimeStamp) > 0 {
		if _, err := time.Parse("2006-01-02 15:04:05", params.TimeStamp); err != nil {
			return errors.New("timestamp 格式\"yyyy-MM-dd HH:mm:ss\"")
		}
	}

	if len(params.Version) == 0 {
		params.Version = "1.0"
	} else if params.Version != "1.0" {
		return errors.New("version 固定为：1.0")
	}

	return nil
}

type bizInterface interface {
	valid() error
}

type responseInterface interface{}

type Response struct {
	Code    string `json:"code,omitempty"`
	Msg     string `json:"msg,omitempty"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
}

type ApiHander interface {
	apiMethod() string
	apiName() string
}

var (
	apiRegistry map[string]AlipayApi
)

func init() {
	apiRegistry = map[string]AlipayApi{}
}

func registerApi(handler ApiHander) {
	apiRegistry[handler.apiMethod()] = AlipayApi{
		apiname:   handler.apiName,
		apimethod: handler.apiMethod,
	}
}

func GetApi(method string) AlipayApi {
	return apiRegistry[method]
}

func GetSupportApis() string {
	lst := "\n=====================SUPPORTED ALIPAY API LIST=====================\n"
	for _, v := range apiRegistry {
		lst += "====[" + v.apiname() + ":" + v.apimethod() + "]====" + "\n"
	}
	return lst
}

type AlipayApi struct {
	params    requestParams
	apiname   func() string
	apimethod func() string
}

func (a *AlipayApi) SetAppId(app_id string) error {
	a.params.AppId = app_id
	if len(a.params.AppId) == 0 {
		return ErrAppIdNil
	}

	if _, ok := secretLst[a.params.AppId]; !ok {
		return ErrSecretNil
	}
	return nil
}

func (a *AlipayApi) SetBizContent(biz bizInterface) error {
	//有几个接口要独立处理
	if reflect.TypeOf(biz).Name() == reflect.TypeOf(Biz_alipay_system_oauth_token{}).Name() {
		b := biz.(Biz_alipay_system_oauth_token)
		a.params.Code = b.Code
		a.params.GrantType = b.GrantType
		a.params.BizContent = ""
		return nil
	} else if reflect.TypeOf(biz).Name() == reflect.TypeOf(Biz_alipay_user_info_share{}).Name() {
		b := biz.(Biz_alipay_user_info_share)
		a.params.AuthToken = b.AuthToken
		a.params.BizContent = ""
		return nil
	} else if reflect.TypeOf(biz).Name() == reflect.TypeOf(Biz_alipay_user_userinfo_share{}).Name() {
		b := biz.(Biz_alipay_user_userinfo_share)
		a.params.AuthToken = b.AuthToken
		a.params.BizContent = ""
		return nil
	}

	if v, err := a.biz_to_string(biz); err != nil {
		return err
	} else {
		a.params.BizContent = v
		return nil
	}
}

func (a *AlipayApi) struct_to_map() map[string]interface{} {
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

func (a *AlipayApi) utf_to_gbk(utf string) (gbk string) {
	gbk_enc := mahonia.NewEncoder("GBK")
	return gbk_enc.ConvertString(utf)
}

func (a *AlipayApi) gbk_to_utf(gbk string) (utf string) {
	enc := mahonia.NewDecoder("GBK")
	return enc.ConvertString(gbk)
}

func (a *AlipayApi) map_to_string(m map[string]interface{}) string {
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

func (a *AlipayApi) biz_to_string(b bizInterface) (string, error) {
	if err := b.valid(); err != nil {
		return "", err
	}
	content := ""
	if v, err := json.Marshal(&b); err != nil {
		return "", err
	} else {
		content = string(v)
	}

	// temp_map := map[string]interface{}{
	// 	"biz": content,
	// }

	// if v, err := json.Marshal(&temp_map); err != nil {
	// 	return "", err
	// } else {
	// 	content = string(v)
	// }
	// return content[8 : len(content)-2], nil
	return content, nil
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

func (a *AlipayApi) verifySign(s, origin_sign, method_key string) bool {
	if a.params.Method == "alipay.user.userinfo.share" {
		//这里是要转义后校验的.NND
		s = strings.Replace(s, "\\", "", -1)
	}
	sign_start_index := strings.Index(s, ",\"sign\"")
	if sign_start_index == -1 {
		return false
	}
	tobe_signed := s[4+len(method_key) : sign_start_index]
	block, _ := pem.Decode(getSecret(a.params.AppId).AliPubRSA)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Printf("Failed to parse RSA public key: %s\n", err)
		return false
	}
	rsaPub, _ := pub.(*rsa.PublicKey)
	t := sha1.New()
	io.WriteString(t, tobe_signed)
	digest := t.Sum(nil)
	data, err := base64.StdEncoding.DecodeString(origin_sign)
	if err != nil {
		fmt.Println("DecodeString sig error, reason: ", err)
		return false
	}
	err = rsa.VerifyPKCS1v15(rsaPub, crypto.SHA1, digest, data)
	if err != nil {
		fmt.Println("Verify sig error, reason: ", err)
		return false
	}

	return true
}

func (a *AlipayApi) request(m map[string]interface{}) (string, error) {
	url_link := "https://openapi.alipay.com/gateway.do"
	if conf.SandBoxEnable {
		url_link = "https://openapi.alipaydev.com/gateway.do"
	}
	logs.DEBUG(fmt.Sprintf("==[请求参数]==[%s]", url_link))
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

func (a *AlipayApi) Run(resp responseInterface) error {
	defer logs.DEBUG("=====================ALIPAY REQUEST END=====================")
	logs.DEBUG("=====================ALIPAY REQUEST START=====================")
	logs.DEBUG(fmt.Sprintf("==[沙盒模式]==[%v]", conf.SandBoxEnable))
	logs.DEBUG(fmt.Sprintf("==[调用方法]==[%s]:[%s]", a.apiname(), a.apimethod()))
	a.params.Method = a.apimethod()
	if err := a.params.valid(); err != nil {
		return err
	}

	rv := reflect.ValueOf(resp)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		if rv.Type() == nil {
			return errors.New("responseInterface is nil")
		}

		if rv.Type().Kind() != reflect.Ptr {
			return errors.New("responseInterface is non-pointer" + rv.Type().String())
		}
		return errors.New("responseInterface is nil" + rv.Type().String())
	}

	m := a.struct_to_map()
	//做请求参数的签名
	__sign := ""
	tobe_sign := a.map_to_string(m)
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

	//把 method 转换为 key
	method_key := strings.Replace(a.params.Method, ".", "_", -1)
	method_key += "_response"

	result_string = strings.Replace(result_string, "error_response", method_key, -1)

	//解析结果
	resp_map := map[string]interface{}{}
	if err := json.Unmarshal([]byte(result_string), &resp_map); err != nil {
		return err
	}
	//原始签名
	if v, ok := resp_map["sign"].(string); ok && len(v) > 0 {
		if pass := a.verifySign(result_string, v, method_key); !pass {
			return ErrVerifySign
		}
	}

	//转码
	result_string = a.gbk_to_utf(result_string)
	logs.DEBUG(fmt.Sprintf("==[响应结果]==[UTF 编码]:[%s]", result_string))
	result_string = strings.Replace(result_string, "\\", "", -1)
	logs.DEBUG(fmt.Sprintf("==[响应结果]==[去掉转义]:[%s]", result_string))
	if err := json.Unmarshal([]byte(result_string), &resp_map); err != nil {
		return err
	}
	//把需要的内容再次换成string
	if v, err := json.Marshal(resp_map[method_key]); err != nil {
		return err
	} else {
		result_string = string(v)
	}

	if err := json.Unmarshal([]byte(result_string), resp); err != nil {
		return err
	}
	return nil
}
