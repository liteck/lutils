package alipay

import "errors"

var (
	ErrMethodNotSupport  = errors.New("METHOD NOT SUPPORT")
	ErrMethodNameNil     = errors.New("METHOD NAME NIL")
	ErrBizContentNameNil = errors.New("BIZ CONTENT NIL")
	ErrAppIdNil          = errors.New("APPID NIL")
	ErrSecretNil         = errors.New("SECRET NIL")
	ErrSign              = errors.New("SIGN ERROR")
	ErrVerifySign        = errors.New("VERIFY SIGN ERROR")
)

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
