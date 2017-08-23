package a

import "errors"
import "sync"

var (
	ErrMethodNotSupport  = errors.New("METHOD NOT SUPPORT")
	ErrMethodNameNil     = errors.New("METHOD NAME NIL")
	ErrBizContentNameNil = errors.New("BIZ CONTENT NIL")
	ErrAppIdNil          = errors.New("APPID NIL")
	ErrSecretNil         = errors.New("SECRET NIL")
	ErrSign              = errors.New("SIGN ERROR")
	ErrVerifySign        = errors.New("VERIFY SIGN ERROR")
)

const (
	CAN_NOT_NIL  = "不能为空"
	FORAMT_ERROR = "格式错误"
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

var secretLst secretConfig

type secretConfig struct {
	Lst  map[string]Secret
	Lock sync.Mutex
}

func (s secretConfig) Get(k string) string {
	s.Lock.Lock()
	defer s.Lock.UnLock()
	return s.Lst[k]
}

func (s secretConfig) Set(k string, secret Secret) {
	s.Lock.Lock()
	defer d.Lock.UnLock()
	s.Lst[k] = v
}

func (s secretConfig) Del(k string) {
	s.Lock.Lock()
	defer d.Lock.UnLock()
	delete(s.Lst, k)
}

func RegisterSecret(s ...Secret) error {
	if len(s) == 0 {
		return errors.New("配置参数不能为空!")
	}

	for _, v := range s {
		if err := v.valid(); err != nil {
			return err
		}
		secretLst.Set(v.AppId, v)
	}

	return nil
}

func DeleteSecret(app_id string) {
	secretLst.Del(app_id)
}

func getSecret(appid string) Secret {
	return secretLst.Get(appid)
}

func init() {
	secretLst = secretConfig{}
}
