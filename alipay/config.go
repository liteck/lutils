package alipay

import "errors"

/**
全局的支付宝参数配置。
全局请求参数。如果有参数则使用原有的。
如果没有。。则使用全局配置的。
**/

type config struct {
	SandBoxEnable bool
}

var conf = newConfig()

func newConfig() *config {
	return &config{SandBoxEnable: false}
}

func EnableSandBox(enable bool) {
	conf.SandBoxEnable = enable
}

type Secret struct {
	AppId     string
	Pid       string
	AliPubRSA []byte
	PrivRSA   []byte
}

func (s *Secret) valid() error {
	if len(s.AppId) == 0 {
		return errors.New("appid 不能为空")
	}

	if len(s.Pid) == 0 {
		return errors.New("pid 不能为空")
	}

	if len(s.AliPubRSA) == 0 {
		return errors.New("支付宝公钥 不能为空")
	}

	if len(s.PrivRSA) == 0 {
		return errors.New("客户私钥 不能为空")
	}

	return nil
}

var secretLst map[string]Secret

func RegisterAlipaySecret(s ...Secret) error {
	if len(s) == 0 {
		return errors.New("配置参数不能为空!")
	}

	for _, v := range s {
		if err := v.valid(); err != nil {
			return err
		}
		secretLst[v.AppId] = v
	}

	return nil
}

func getSecret(appid string) Secret {
	return secretLst[appid]
}

func init() {
	secretLst = map[string]Secret{}
}
