package w

import (
	"encoding/json"
	"fmt"
	"lutils/httplib"
	"lutils/logs"
	"reflect"
	"strings"
)

//公共请求参数
type requestParams struct {
	AppId  string `json:"appid"`
	Secret string `json:"secret"`
}

func (r *requestParams) valid() error {
	if len(r.AppId) == 0 {
		return ErrAppIdNil
	}
	if len(r.Secret) == 0 {
		return ErrSecretNil
	}
	return nil
}

type reqInterface interface {
	valid() error
}

type responseInterface interface{}

type Response struct {
	ErrCode float64 `json:"errcode,omitempty"`
	ErrMsg  string  `json:"errmsg,omitempty"`
}

type ApiHander interface {
	apiUrl() string
	apiName() string
	apiMethod() string
}

var (
	apiRegistry map[string]WechatApi
)

func init() {
	apiRegistry = map[string]WechatApi{}
}

func registerApi(handler ApiHander) {
	apiRegistry[handler.apiUrl()] = WechatApi{
		apiurl:    handler.apiUrl,
		apiname:   handler.apiName,
		apimethod: handler.apiMethod,
	}
}

func GetApi(method string) WechatApi {
	return apiRegistry[method]
}

func GetSupportApis() string {
	lst := "\n=====================SUPPORTED WECHAT API LIST=====================\n"
	for _, v := range apiRegistry {
		lst += "====[" + v.apiname() + ":" + v.apiurl() + "]====" + "\n"
	}
	return lst
}

type WechatApi struct {
	params    requestParams
	req       reqInterface
	apiurl    func() string
	apiname   func() string
	apimethod func() string
}

func (w *WechatApi) SetAppId(app_id string) error {
	w.params.AppId = app_id
	if len(w.params.AppId) == 0 {
		return ErrAppIdNil
	}

	if v, ok := secretLst[w.params.AppId]; !ok {
		return ErrSecretNil
	} else {
		w.params.Secret = v.AppSecret
	}
	return nil
}

func (w *WechatApi) SetReqContent(v reqInterface) error {
	if err := v.valid(); err != nil {
		return err
	}
	w.req = v
	return nil
}

func (w *WechatApi) apiMethod() string {
	return "POST"
}

func (w *WechatApi) getSecret() (*Secret, error) {
	if len(w.params.AppId) == 0 {
		return nil, ErrAppIdNil
	}

	if v, ok := secretLst[w.params.AppId]; !ok {
		return nil, ErrSecretNil
	} else {
		return &v, nil
	}
}

func (w *WechatApi) struct_to_map() map[string]interface{} {
	var data = make(map[string]interface{})
	{
		t := reflect.TypeOf(w.params)
		v := reflect.ValueOf(w.params)

		for i := 0; i < t.NumField(); i++ {
			key := t.Field(i).Name
			value := v.Field(i).Interface()
			tag := t.Field(i).Tag.Get("json")
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
	}

	{
		t := reflect.TypeOf(w.req)
		v := reflect.ValueOf(w.req)

		for i := 0; i < t.NumField(); i++ {
			key := t.Field(i).Name
			value := v.Field(i).Interface()
			tag := t.Field(i).Tag.Get("json")
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
	}

	return data
}

func (w *WechatApi) request(m map[string]interface{}) (string, error) {
	var http_request *httplib.BeegoHTTPRequest
	if w.apimethod() == "POST" {
		http_request = httplib.Post(w.apiurl())
	} else if w.apimethod() == "GET" {
		http_request = httplib.Get(w.apiurl())
	}
	logs.DEBUG(fmt.Sprintf("==[请求参数]==[%s]", w.apiurl()))
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

func (w *WechatApi) Run(resp responseInterface) error {
	defer logs.DEBUG("=====================WECHAT REQUEST END=====================")
	logs.DEBUG("=====================WECHAT REQUEST START=====================")
	logs.DEBUG(fmt.Sprintf("==[调用方法]==[%s]:[%s]", w.apiname(), w.apiurl()))

	if err := w.params.valid(); err != nil {
		return err
	}

	m := w.struct_to_map()
	//准备请求
	result_string := ""
	if v, err := w.request(m); err != nil {
		return err
	} else {
		result_string = v
		logs.DEBUG(fmt.Sprintf("==[响应结果]==[%s]", result_string))
	}

	if err := json.Unmarshal([]byte(result_string), resp); err != nil {
		return err
	}
	return nil
}
